//
// This file contains cookie-related handlers, for setting/testing
// the HTTP-sessions of the remote user.
//

package main

import (
	"net/http"

	"github.com/gorilla/securecookie"
)

//
// The secure-cookie object we use.
//
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

//
// Get the username of the logged-in user, if any.
//
func getUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

//
// Set the name of the user in the session, such that it persists.
//
func setSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

//
// Clear the existing session, if any.
//
func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

//
// Is there a user logged in?  There is if there is a non-empty
// username set.
//
func loggedIn(request *http.Request) bool {
	userName := getUserName(request)
	if userName != "" {
		return true
	} else {
		return false
	}
}
