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
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/robfig/cron"
	"golang.org/x/crypto/bcrypt"
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
	UserFile       string
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
	// Command-line options.
	//
	db := flag.String("database", "alerts.db", "The SQLite3 database to use")
	host := flag.String("host", "localhost", "The host to listen upon")

	hash := flag.String("hash", "", "Hash the specified plaintext password")
	port := flag.Int("port", 8080, "The port to listen upon")

	users := flag.String("auth-file", "users", "The username/password file to authentication against")
	notify := flag.String("notify-binary", "purple-notify", "The binary to execute to issue notifications")
	renotify := flag.String("renotify-binary", "purple-renotify", "The binary to execute for notification reminders")
	flag.Parse()

	if *hash != "" {
		pBytes, _ := bcrypt.GenerateFromPassword([]byte(*hash), 14)
		pCrypt := string(pBytes)
		fmt.Printf("%s\n", pCrypt)
		os.Exit(0)
	}

	CONFIG.UserFile = *users
	CONFIG.NotifyBinary = *notify
	CONFIG.ReNotifyBinary = *renotify

	//
	// Setup our database
	//
	SetupDB(*db)

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
	// Configure our routes.
	//
	var router = mux.NewRouter()
	router.HandleFunc("/", indexPageHandler)
	router.HandleFunc("/purple.js", jsPage).Methods("GET")

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
	// Launch our HTTP-server.
	//
	bind := fmt.Sprintf("%s:%d", *host, *port)

	fmt.Printf("Listening on http://%s/\n", bind)

	//
	// Wire up logging.
	//
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	//
	// Launch the server.
	//
	err := http.ListenAndServe(bind, loggedRouter)
	if err != nil {
		fmt.Printf("\nError: %s\n", err.Error())
	}
}
