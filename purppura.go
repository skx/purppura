//
// This is a golang port of the purple-alert server.
//
// Incoming submissions are received over HTTP POSTS to /events, and
// alerts are processed as expected.
//
// There is also a web-view, for viewing events and that can be used
// to raise/ack/clear them.
//
// Steve
// --
//
package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/robfig/cron"
)

var (
	version = "unreleased"
)

//
// This structure describes a single event.
//
type Alert struct {
	//
	// The detail of the event.
	//
	Detail string

	//
	// The ID of the event, e.g. "heartbeat", "backup"
	//
	// The ID should be unique per-source.
	//
	ID string

	//
	// The time the alert should raise, which will be submitted
	// as "+5m", "now", "clear", and be translated to an absolute
	// time - in seconds past the epoch.
	//
	Raise string

	//
	// The subject/summary of the event.
	//
	Subject string

	//
	// The source (IP address) of the event.
	//
	Source string

	//
	// Sigh - these are private fields, used when we output
	// our list of events to JSON.
	//
	Status     string
	RaiseAt    string
	NotifiedAt string
}

//
// Configuration-options, as set by the command-line flags.
//
type Config struct {
	NotifyBinary   string
	ReNotifyBinary string
}

//
// The configuration options we're actually using.
//
var CONFIG Config

//
// cronjob: Process outstanding alerts.
//
func ProcessAlerts() {
	fmt.Printf("Processing events at %s\n", time.Now())

	//
	// Reap expired events.
	//
	err := Reap()
	if err != nil {
		fmt.Printf("\tError reaping - %s\n", err.Error())
		return
	}

	//
	// Lets do the time-warp, again!
	//
	// (This is required, although with some restructuring it could
	// be omitted.)
	//
	err = Warp()
	if err != nil {
		fmt.Printf("\tError warping time - %s\n", err.Error())
		return
	}

	//
	// Notify outstanding alerts.
	//
	err = Notify()
	if err != nil {
		fmt.Printf("\tError notifying - %s\n", err.Error())
		return
	}

	//
	// ReNotify repeatedly outstanding alerts.
	//
	err = ReNotify()
	if err != nil {
		fmt.Printf("\tError repeating old notification - %s\n", err.Error())
		return
	}
}

//
// Entry point.
//
func main() {

	//
	// Command-line options - database
	//
	db := flag.String("database", "alerts.db", "The SQLite3 database to use")
	host := flag.String("host", "localhost", "The host to listen upon")
	port := flag.Int("port", 8080, "The port to listen upon")
	notify := flag.String("notify-binary", "purppura-notify", "The binary to execute to issue notifications")
	renotify := flag.String("renotify-binary", "purppura-renotify", "The binary to execute for notification reminders")
	ver := flag.Bool("version", false, "Show our release, and exit")

	//
	// User-options
	//
	userAdd := flag.String("user-add", "", "Add the specified user.")
	userDel := flag.String("user-delete", "", "Remove the specified user.")

	flag.Parse()

	//
	// Setup our configuration-options.
	//
	CONFIG.NotifyBinary = *notify
	CONFIG.ReNotifyBinary = *renotify

	//
	// Setup our database
	//
	SetupDB(*db)

	//
	// Are we adding a user?
	//
	if *userAdd != "" {
		//
		// Add the user.
		//

		reader := bufio.NewReader(os.Stdin)

		fmt.Printf("Enter Password for user %s: ", *userAdd)
		pass, _ := reader.ReadString('\n')
		pass = strings.TrimSpace(pass)
		err := addUser(*userAdd, pass)
		if err != nil {
			fmt.Printf("Error %s", err.Error())
		}
		os.Exit(0)
	}

	//
	// Deleting a user?
	//
	if *userDel != "" {
		//
		// Delete the user.
		//
		err := delUser(*userDel)
		if err != nil {
			fmt.Printf("Error %s", err.Error())
		}
		os.Exit(0)
	}

	//
	// Showing our version number?
	//
	if *ver {
		fmt.Printf("purppura %s\n", version)
		os.Exit(0)
	}

	//
	// Create a cron scheduler
	//
	c := cron.New()

	//
	// Add a cron-job to process raise/clear events.
	//
	c.AddFunc("@every 10s", func() { ProcessAlerts() })

	//
	// Launch the cron-scheduler.
	//
	c.Start()

	//
	// Configure our secure cookies
	//
	LoadCookie()

	//
	// Configure our routes.
	//
	var router = mux.NewRouter()
	router.HandleFunc("/", indexPageHandler)
	router.HandleFunc("/purppura.js", jsPage).Methods("GET")

	router.HandleFunc("/login", loginForm).Methods("GET")
	router.HandleFunc("/login", loginHandler).Methods("POST")

	router.HandleFunc("/logout", logoutHandler).Methods("GET")
	router.HandleFunc("/logout", logoutHandler).Methods("POST")

	router.HandleFunc("/events", eventsHandler).Methods("GET")
	router.HandleFunc("/events", parseGhPost).Methods("POST")

	router.HandleFunc("/acknowledge/{id}", ackEvent).Methods("GET")
	router.HandleFunc("/clear/{id}", clearEvent).Methods("GET")
	router.HandleFunc("/raise/{id}", raiseEvent).Methods("GET")

	http.Handle("/", router)

	//
	// Show what we're goign to bind upon.
	//
	bind := fmt.Sprintf("%s:%d", *host, *port)
	fmt.Printf("Listening on http://%s/\n", bind)

	//
	// Wire up logging.
	//
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	//
	// Wire up context (i.e. cookie-based session stuff.)
	//
	contextRouter := AddContext(loggedRouter)

	//
	// Launch the server.
	//
	err := http.ListenAndServe(bind, contextRouter)
	if err != nil {
		fmt.Printf("\nError: %s\n", err.Error())
	}
}
