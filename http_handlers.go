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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func RemoteIP(request *http.Request) string {

	//
	// Get the X-Forwarded-For header, if present.
	//
	xForwardedFor := request.Header.Get("X-Forwarded-For")

	//
	// No forwarded IP?  Then use the remote address directly.
	//
	if xForwardedFor == "" {
		ip, _, _ := net.SplitHostPort(request.RemoteAddr)
		return ip
	}

	entries := strings.Split(xForwardedFor, ",")
	address := strings.TrimSpace(entries[0])
	return (address)
}

//
// Parse the incoming POST request.
//
func parseGhPost(res http.ResponseWriter, request *http.Request) {

	//
	// We'll read JSON from STDIN.
	//
	// We do this manually because we want to see if we're
	// getting a single event:
	//
	//  {id:"blah",...}
	//
	// Or an array of events:
	//
	//  [{id:"blah", ..},{id:"more-blah"..}
	//
	content, err := ioutil.ReadAll(request.Body)

	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}

	//
	// The incoming JSON might contain a single entry, or
	// an array of entries.
	//
	var single Alert
	var multi []Alert

	//
	// Decode - into the array, or single entry, as appropriate.
	//
	if strings.HasPrefix(string(content), "[") {
		err = json.Unmarshal(content, &multi)
	} else {
		err = json.Unmarshal(content, &single)
	}

	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}

	//
	// Get the source of the submission.
	//
	ip := RemoteIP(request)

	//
	// Did we get multiple entries?
	//
	if len(multi) > 0 {

		//
		// For each one - add it
		//
		for _, ent := range multi {

			//
			// Ensure the alert has the correct source.
			//
			ent.Source = ip

			//
			// Add it.
			//
			err = addEvent(ent)

			if err != nil {
				http.Error(res, err.Error(), 400)
				return
			}
		}
	} else {

		//
		// Ensure the alert has the correct source.
		//
		single.Source = ip

		//
		// Add it.
		//
		err = addEvent(single)

		if err != nil {
			http.Error(res, err.Error(), 400)
			return
		}
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
	http.Redirect(res, req, "/#acknowledged", 302)

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
	http.Redirect(res, req, "/#pending", 302)

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
	http.Redirect(res, req, "/#raised", 302)

}

//
// Serve a static-resource
//
func serveResource(response http.ResponseWriter, request *http.Request, resource string, mime string) {
	tmpl, err := Asset(resource)
	if err != nil {
		fmt.Fprintf(response, err.Error())
		return
	}
	response.Header().Set("Content-Type", mime)
	fmt.Fprintf(response, string(tmpl))
}

//
// Serve the login-form
//
func loginForm(response http.ResponseWriter, request *http.Request) {
	serveResource(response, request, "data/login.html", "text/html")
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
	valid, err := validateUser(name, pass)
	if err != nil {
		http.Error(response, err.Error(), 400)
		return
	}

	//
	// If this succeeded then let the login succeed.
	//
	if valid {
		setSession(name, response)
		http.Redirect(response, request, "/", 302)
		return
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
// This is used by /purppura.js to dynamically update the display.
//
func eventsHandler(response http.ResponseWriter, request *http.Request) {

	//
	// Ensure the user is logged-in.
	//
	if !loggedIn(request) {
		http.Redirect(response, request, "/login", 302)
		return
	}

	//
	// Get all the alerts, and their states.
	//
	results, err := Alerts()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	//
	// Ensure that we send a suitable content-type.
	//
	response.Header().Set("Content-Type", "application/json")

	//
	// Output the alerts.
	//
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

	serveResource(response, request, "data/index.html", "text/html")
}

//
// serve our JS
//
func jsPage(response http.ResponseWriter, request *http.Request) {
	serveResource(response, request, "data/purppura.js", "application/javascript")
}
