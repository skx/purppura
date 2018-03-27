//
// This is a golang port of the purple-alert server.
//
// Incoming submissions are received over HTTP POSTS to /events, and
// alerts are processed as expected.
//
// There is also a web-view, for processing events.
//
// Steve
// --
//
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

//
// Parse the incoming POST request.
//
func parseGhPost(res http.ResponseWriter, request *http.Request) {

	//
	// We'll read JSON from STDIN.
	//
	decoder := json.NewDecoder(request.Body)

	//
	// We parse into a new structure.
	//
	var t Alert
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	//
	// Get the source of the submission.
	//
	ip, _, err := net.SplitHostPort(request.RemoteAddr)
	if err != nil {
		panic(err)
	}

	//
	// Update the structure with the source
	//
	t.Source = ip

	//
	// Add the event
	//
	err = addEvent(t)
	if err != nil {
		panic(err)
	}

	//
	// Send a simple reply to the caller.
	//
	fmt.Fprintf(res, "OK")
}

//
// Acknowledge an event.
//
func ackEvent(res http.ResponseWriter, req *http.Request) {

	//
	// Ensure the user is logged-in.
	//
	if !loggedIn(req) {
		http.Redirect(res, req, "/login", 302)
		return
	}

	//
	// Get the ID we're going to acknowledge
	//
	vars := mux.Vars(req)
	id := vars["id"]

	AckEvent(id)
	http.Redirect(res, req, "/", 302)

}

//
// Clear an event.
//
func clearEvent(res http.ResponseWriter, req *http.Request) {
	//
	// Ensure the user is logged-in.
	//
	if !loggedIn(req) {
		http.Redirect(res, req, "/login", 302)
		return
	}

	//
	// Get the ID we're going to clear.
	//
	vars := mux.Vars(req)
	id := vars["id"]

	ClearEvent(id)
	http.Redirect(res, req, "/", 302)

}

//
// Raise an event.
//
func raiseEvent(res http.ResponseWriter, req *http.Request) {
	//
	// Ensure the user is logged-in.
	//
	if !loggedIn(req) {
		http.Redirect(res, req, "/login", 302)
		return
	}

	//
	// Get the ID we're going to raise.
	//
	vars := mux.Vars(req)
	id := vars["id"]

	RaiseEvent(id)
	http.Redirect(res, req, "/", 302)

}

//
// Serve a static-resource
//
func serveResource(response http.ResponseWriter, request *http.Request, resource string) {
	tmpl, err := Asset(resource)
	if err != nil {
		fmt.Fprintf(response, err.Error())
		return
	}
	fmt.Fprintf(response, string(tmpl))
}

//
// Serve the login-form
//
func loginForm(response http.ResponseWriter, request *http.Request) {
	serveResource(response, request, "data/login.html")
}

//
// Process a login-event.
//
func loginHandler(response http.ResponseWriter, request *http.Request) {
	//
	// Get the username/password from the incoming form
	// submission.
	//
	name := request.FormValue("name")
	pass := request.FormValue("password")

	//
	// Open our list of users/passwords
	//
	inFile, err := os.Open(CONFIG.UserFile)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	//
	// Process line by line
	//
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {

		//
		// The user/password are space-separated
		//
		fields := strings.Fields(scanner.Text())

		//
		// If we got two fields then we can test them
		//
		if len(fields) == 2 {

			//
			// The username + hash from the field-splitting
			//
			user := fields[0]
			hash := fields[1]

			//
			// Do we have a match - if so set the session
			// and redirect to the server root
			//
			if user == name {
				err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
				if err == nil {
					setSession(name, response)
					http.Redirect(response, request, "/", 302)
					return
				}
			}
		}
	}

	//
	// Failure to login, redirect to try again.
	//
	http.Redirect(response, request, "/login#failed", 302)
}

//
// logout handler
//
func logoutHandler(response http.ResponseWriter, request *http.Request) {
	clearSession(response)
	http.Redirect(response, request, "/", 302)
}

//
// Get all events as a JSON array.
//
// This is used by /purple.js to dynamically update the display.
//
func eventsHandler(response http.ResponseWriter, request *http.Request) {

	//
	// Ensure the user is logged-in.
	//
	if !loggedIn(request) {
		http.Redirect(response, request, "/login", 302)
		return
	}

	results, err := Alerts()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")

	if len(results) > 0 {
		out, _ := json.Marshal(results)
		fmt.Fprintf(response, "%s", out)
	} else {
		fmt.Fprintf(response, "[]")
	}
}

//
// index page
//
func indexPageHandler(response http.ResponseWriter, request *http.Request) {
	//
	// Ensure the user is logged-in.
	//
	if !loggedIn(request) {
		http.Redirect(response, request, "/login", 302)
		return
	}

	serveResource(response, request, "data/index.html")
}

//
// serve our JS
//
func jsPage(response http.ResponseWriter, request *http.Request) {
	serveResource(response, request, "data/purple.js")
}
