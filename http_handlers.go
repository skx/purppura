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
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
)

//
// The secure-cookie object we use.
//
var cookieHandler *securecookie.SecureCookie

//
// If there is an cookie-file then read it
//
func LoadCookie() {

	//
	// Read the hash
	//
	hash, err := ioutil.ReadFile(".cookie.hsh")
	if err == nil {

		//
		// If there was no error read the block
		//
		block, err := ioutil.ReadFile(".cookie.blk")
		if err == nil {

			//
			// And create the cookie-helper.
			//
			cookieHandler = securecookie.New(hash, block)
			return
		}
	}

	//
	// So we either failed to find, or failed to read, the existing
	// values.  (Perhaps this is the first run.)
	//
	// Generate random values.
	//
	h := securecookie.GenerateRandomKey(64)
	b := securecookie.GenerateRandomKey(32)

	//
	// Now write them out.
	//
	// If writing fails then we'll use the values, and this means
	// when the server restarts authentication will need to to be
	// repeated by the users.
	//
	// (i.e. They'll be logged out.)
	//
	err = ioutil.WriteFile(".cookie.hsh", h, 0644)
	if err != nil {
		fmt.Printf("WARNING: failed to write .cookie.hsh for persistent secure cookie")
		cookieHandler = securecookie.New(h, b)
		return
	}
	err = ioutil.WriteFile(".cookie.blk", b, 0644)
	if err != nil {
		fmt.Printf("WARNING: failed to write .cookie.blk for persistent secure cookie")
		cookieHandler = securecookie.New(h, b)
		return
	}

	//
	// Create the cookie, if we got here we've saved the data
	// for the next restart.
	//
	cookieHandler = securecookie.New(h, b)
}

//
// Add context to our HTTP-handlers.
//
func AddContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//
		// If we have a session-cookie
		//
		if cookie, err := r.Cookie("cookie"); err == nil {

			// Make a map
			cookieValue := make(map[string]string)

			// Decode it.
			if err = cookieHandler.Decode("cookie", cookie.Value, &cookieValue); err == nil {
				//
				// Add the context to the handler, with the
				// username.
				//
				userName := cookieValue["name"]
				ctx := context.WithValue(r.Context(), "Username", userName)
				//
				// And fire it up.
				//
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			} else {
				//
				// We failed to decode the cookie.
				//
				// Probably it was created with the random-key
				// of a previous run of the server.  So we
				// just fall-back to assuming we're not logged
				// in, and have no context.
				//
				next.ServeHTTP(w, r)
				return
			}
		} else {
			next.ServeHTTP(w, r)
			return
		}
	})
}

//
// Get the remote IP of the request-submitter.
//
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
			// Ensure we have all the fields we expect
			//
			if ent.Subject == "" {
				http.Error(res, "Missing 'subject' field", 500)
				return
			}
			if ent.ID == "" {
				http.Error(res, "Missing 'ID' field", 500)
				return
			}
			if ent.Raise == "" {
				http.Error(res, "Missing 'raise' field", 500)
				return
			}
			if ent.Detail == "" {
				http.Error(res, "Missing 'detail' field", 500)
				return
			}

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
		// Ensure we have all the fields we expect
		//
		if single.Subject == "" {
			http.Error(res, "Missing 'subject' field", 500)
			return
		}
		if single.ID == "" {
			http.Error(res, "Missing 'ID' field", 500)
			return
		}
		if single.Raise == "" {
			http.Error(res, "Missing 'raise' field", 500)
			return
		}
		if single.Detail == "" {
			http.Error(res, "Missing 'detail' field", 500)
			return
		}

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
	username := req.Context().Value("Username")
	if username == nil {
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
	username := req.Context().Value("Username")
	if username == nil {
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
	username := req.Context().Value("Username")
	if username == nil {
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
	tmpl, err := getResource(resource)
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
	serveResource(response, request, "data/login.html", "text/html; charset=utf-8")
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

		value := map[string]string{
			"name": name,
		}
		if encoded, err := cookieHandler.Encode("cookie", value); err == nil {
			cookie := &http.Cookie{
				Name:  "cookie",
				Value: encoded,
				Path:  "/",
			}
			http.SetCookie(response, cookie)
		}

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
	cookie := &http.Cookie{
		Name:   "cookie",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
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
	username := request.Context().Value("Username")
	if username == nil {
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
	username := request.Context().Value("Username")
	if username == nil {
		http.Redirect(response, request, "/login", 302)
		return
	}

	serveResource(response, request, "data/index.html", "text/html; charset=utf-8")
}

//
// serve our JS
//
func jsPage(response http.ResponseWriter, request *http.Request) {
	serveResource(response, request, "data/purppura.js", "application/javascript")
}
