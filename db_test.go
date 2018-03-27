//
//  Basic testing of our DB primitives
//

package main

import (
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

//
//  Temporary location for database
//
var path string

//
// Create a temporary database
//
func FakeDB() {
	p, err := ioutil.TempDir(os.TempDir(), "prefix")
	if err == nil {
		path = p
	}

	//
	// Setup the tables.
	//
	SetupDB(p + "/db.sql")

}

//
// Test that functions return errors if setup hasn't been called.
//
func TestMissingInit(t *testing.T) {

	//
	// Regexp to match the error we expect to receive.
	//
	reg, _ := regexp.Compile("SetupDB not called")

	var x Alert
	err := addEvent(x)
	if !reg.MatchString(err.Error()) {
		t.Errorf("Got wrong error: %v", err)
	}

	_, err = Alerts()
	if !reg.MatchString(err.Error()) {
		t.Errorf("Got wrong error: %v", err)
	}

	_, err = GetAlert(3)
	if !reg.MatchString(err.Error()) {
		t.Errorf("Got wrong error: %v", err)
	}

	err = SetEvent("2", "test")
	if !reg.MatchString(err.Error()) {
		t.Errorf("Got wrong error: %v", err)
	}

	err = Reap()
	if !reg.MatchString(err.Error()) {
		t.Errorf("Got wrong error: %v", err)
	}

	err = Notify()
	if !reg.MatchString(err.Error()) {
		t.Errorf("Got wrong error: %v", err)
	}

	err = ReNotify()
	if !reg.MatchString(err.Error()) {
		t.Errorf("Got wrong error: %v", err)
	}

	err = Warp()
	if !reg.MatchString(err.Error()) {
		t.Errorf("Got wrong error: %v", err)
	}

	_, err = validateUser("moi", "kissa")
	if !reg.MatchString(err.Error()) {
		t.Errorf("Got wrong error: %v", err)
	}

	err = addUser("moi", "kissa")
	if !reg.MatchString(err.Error()) {
		t.Errorf("Got wrong error: %v", err)
	}

	err = delUser("moi")
	if !reg.MatchString(err.Error()) {
		t.Errorf("Got wrong error: %v", err)
	}

}

//
// Test login works.
//
func TestUserAuth(t *testing.T) {
	// Create a fake database
	FakeDB()

	// Missing user won't validate
	res, err := validateUser("_moi", "_kissa")

	if res != false {
		t.Errorf("Missing user was validated!")
	}

	// Add the user.
	err = addUser("_moi", "_kissa")
	if err != nil {
		t.Errorf("Error adding user!")
	}

	// Now the user should validate
	res, err = validateUser("_moi", "_kissa")
	if res != true {
		t.Errorf("Failed to validate user, post-addition")
	}

	// Delete the user
	err = delUser("_moi")
	if err != nil {
		t.Errorf("Failed to delete a user")
	}

	// Deleted user won't validate
	res, err = validateUser("_moi", "_kissa")
	if res != false {
		t.Errorf("Deleted user was validated!")
	}
	if err != nil {
		t.Errorf("Error validating deleted user!")
	}

	//
	// Cleanup here because otherwise later tests will
	// see an active/valid DB-handle.
	//
	db.Close()
	db = nil
	os.RemoveAll(path)
}

//
// Test adding an alert works.
//
func TestAddRetrieveAlert(t *testing.T) {
	// Create a fake database
	FakeDB()

	var tmp Alert
	tmp.Source = "127.0.0.1"
	tmp.ID = "heartbeat"
	tmp.Raise = "+5m"
	tmp.Detail = "This is some detail"
	tmp.Subject = "Re: your mail"

	//
	// Get all alerts - should be zero.
	//
	alerts, err := Alerts()
	if err != nil {
		t.Errorf("Error fetching empty alerts")
	}

	if len(alerts) != 0 {
		t.Errorf("Expected zero alerts, but found more")
	}

	//
	// Now add our event
	//
	err = addEvent(tmp)
	if err != nil {
		t.Errorf("Error adding event")
	}

	//
	// Retrieve the update alerts and we should see it.
	//
	alerts, err = Alerts()
	if err != nil {
		t.Errorf("Error fetching updated alerts")
	}
	if len(alerts) != 1 {
		t.Errorf("Expected one alert, but found a different number")
	}

	//
	// We should be able to re-add the alert, and that won't
	// result in a change - because the source+ID are the same
	//
	tmp.Raise = "+40m"
	err = addEvent(tmp)
	if err != nil {
		t.Errorf("Error re-adding event")
	}

	//
	// Retrieve the update alerts and we should see there is still one.
	//
	alerts, err = Alerts()
	if err != nil {
		t.Errorf("Error fetching updated alerts")
	}
	if len(alerts) != 1 {
		t.Errorf("Expected one alert, but found a different number")
	}

	//
	// But if we change the ID then it will be "new".
	//
	tmp.ID = "steve-mail-check"
	err = addEvent(tmp)
	if err != nil {
		t.Errorf("Error re-adding event")
	}

	//
	// Retrieve the update alerts and we should see there is still one.
	//
	alerts, err = Alerts()
	if err != nil {
		t.Errorf("Error fetching updated alerts")
	}
	if len(alerts) != 2 {
		t.Errorf("Expected two alerts, but found a different number")
	}

	//
	// Cleanup here because otherwise later tests will
	// see an active/valid DB-handle.
	//
	db.Close()
	db = nil
	os.RemoveAll(path)
}

// Test adding an alert works.
//
func TestUpdatingAlert(t *testing.T) {
	// Create a fake database
	FakeDB()

	// Add an alert
	// Get all alerts - to find the id
	// change the state : ack
	// check that worked
	// change the state : raise
	// check that worked
	// change the state : clear
	// check that worked
	// reap
	// check it is gone

	// Cleanup here because otherwise later tests will
	// see an active/valid DB-handle.
	//
	db.Close()
	db = nil
	os.RemoveAll(path)
}
