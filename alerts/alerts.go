//
// The Alerts package contains code relating to the persistance and
// manipulation of alerts.
//
//

package alerts

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/microcosm-cc/bluemonday"
	"github.com/skx/purppura/alert"
	"github.com/skx/purppura/util"
	"golang.org/x/crypto/bcrypt"
)

// Alerts is our object for interfacing with alerts.
type Alerts struct {
	// Storage for our DB-handle.
	db *sql.DB
}

// AlertRaise is a function-signature.
//
// When an alert becomes raised we need to notify a human, we don't have
// support for doing that directly, instead we execute an external process
// to do that - here is the callback function which is invoked to do that.
type AlertRaise func(x alert.Alert, config string) error

// New is the constructor for our object.
func New() (*Alerts, error) {
	var err error

	//
	// Create the object.
	//
	m := new(Alerts)

	//
	// Get our DSN from the environment
	//
	dsn := os.Getenv("PURPLE_DSN")
	if dsn == "" {
		return m, errors.New("You must specify the environmental variable 'PURPLE_DSN' with your DB details")
	}

	//
	// Connect to the database, using the DSN specified.
	//
	m.db, err = sql.Open("mysql", dsn)
	if err != nil {
		return m, err
	}

	//
	// Ensure that the connection succeeded.
	//
	err = m.db.Ping()
	if err != nil {
		return m, err
	}

	//
	// All is OK
	//
	return m, err
}

// Close ensures that our database-connection is closed.
func (s *Alerts) Close() {
	s.db.Close()
	s.db = nil
}

// AddEvent adds a new event to our persistant storage.
func (s *Alerts) AddEvent(data alert.Alert) error {

	if s.db == nil {
		panic("Working with a closed database - bug")
	}

	//
	// Convert the raise-time submitted into an absolute time.
	//
	raise, err := util.Str2Unix(data.Raise)
	if err != nil {
		return err
	}

	//
	// See if we've seen this alert before
	//
	id := -1

	//
	// Find any previous event with this source/ID combo.
	//
	rows, err := s.db.Query("SELECT i FROM events WHERE id=? AND source=?", data.ID, data.Source)
	if err != nil {
		return err
	}
	defer rows.Close()

	//
	// For each row in the result-set parse into an instance of our
	// structure and add to the list.
	//
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return err
		}
	}
	err = rows.Err()
	if err != nil {
		return err
	}

	//
	// If we didn't find a previous ID then we must believe this
	// is a new event.  Insert it.
	//
	if id == -1 {
		stmt, err := s.db.Prepare("INSERT INTO events( id, source, subject, detail, raise_at ) VALUES( ?, ?, ?, ?, ? )")
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(data.ID,
			data.Source,
			data.Subject,
			data.Detail,
			raise)
		if err != nil {
			return err
		}

	} else {
		//
		// Otherwise update the existing event in-place.
		//
		up, err := s.db.Prepare("UPDATE events SET raise_at=?, subject=?, detail=?  WHERE i=?")
		defer up.Close()
		if err != nil {
			return err
		}

		_, err = up.Exec(raise, data.Subject, data.Detail, id)
		if err != nil {
			return err
		}
	}

	return nil
}

// Return all alerts, and their details.
func (s *Alerts) Alerts() ([]alert.Alert, error) {

	if s.db == nil {
		panic("Working with a closed database - bug")
	}

	//
	// We're going to sanitize body-details.
	//
	helper := bluemonday.UGCPolicy()

	//
	// Our return-result.
	//
	var results []alert.Alert

	//
	// Select the data.
	//
	rows, err := s.db.Query("SELECT i,source,status,subject,detail,raise_at, notified_at, notify_count from events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//
	// For each row in the result-set parse into an instance of our
	// structure and add to the list.
	//
	for rows.Next() {
		var tmp alert.Alert

		err := rows.Scan(&tmp.ID, &tmp.Source, &tmp.Status, &tmp.Subject, &tmp.Detail, &tmp.RaiseAt, &tmp.NotifiedAt, &tmp.NotifyCount)
		if err != nil {
			return nil, err
		}

		//
		// The detail should be sanitized to avoid XSS attacks
		// against our Web-UI.
		//
		tmp.Detail = helper.Sanitize(tmp.Detail)

		//
		// Add the new record.
		//
		results = append(results, tmp)

	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	//
	// Return the results.
	//
	return results, nil
}

// GetAlert returns a single alert, via the identifier.
func (s *Alerts) GetAlert(id int) (alert.Alert, error) {

	if s.db == nil {
		panic("Working with a closed database - bug")
	}

	//
	// Our return-result.
	//
	var result alert.Alert

	//
	// Select the data.
	//
	rows, err := s.db.Query("SELECT i,source,status,subject,detail,raise_at, notified_at, notify_count from events WHERE i=?", id)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	//
	// For each row in the result-set parse into an instance of our
	// structure and add to the list.
	//
	for rows.Next() {
		err := rows.Scan(&result.ID, &result.Source, &result.Status, &result.Subject, &result.Detail, &result.RaiseAt, &result.NotifiedAt, &result.NotifyCount)
		if err != nil {
			return result, err
		}

	}
	err = rows.Err()
	if err != nil {
		return result, err
	}
	return result, nil
}

// setEventState changes the given alert to have the specified state.
func (s *Alerts) setEventState(id string, state string) error {
	if s.db == nil {
		panic("Working with a closed database - bug")
	}

	stmt, err := s.db.Prepare("UPDATE events SET status=? WHERE i=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(state, id)
	if err != nil {
		return err
	}
	stmt.Close()
	return nil
}

// AckEvent changes the state of the given alert to be "acknowledged".
func (s *Alerts) AckEvent(id string) error {
	return (s.setEventState(id, "acknowledged"))
}

// ClearEvent changes the state of the given alert to be "cleared".
func (s *Alerts) ClearEvent(id string) error {
	return (s.setEventState(id, "cleared"))
}

// RaiseEvent changes the state of the given alert to be "raised".
func (s *Alerts) RaiseEvent(id string) error {
	return (s.setEventState(id, "raised"))
}

//
// Remove old/obsolete events from the database.
//
func (s *Alerts) Reap() error {
	if s.db == nil {
		panic("Working with a closed database - bug")
	}

	//
	// The query we'll execute.
	//
	clean, err := s.db.Prepare("DELETE FROM events WHERE status='cleared' OR raise_at < 1")
	if err != nil {
		return err
	}
	defer clean.Close()

	_, err = clean.Exec()
	if err != nil {
		return err
	}

	return nil
}

// Notify is called to trigger the _initial_ notification for any event
// which has become raised.
//
// It will not be involved in _re_notification.
func (s *Alerts) Notify(callback AlertRaise, config string) error {
	var err error
	if s.db == nil {
		panic("Working with a closed database - bug")
	}

	var toRaise []int

	//
	// Our query.
	//
	//  * For every alert in the pending state
	//  * If there is a `raise_at` time, and that time is in the past
	//    then we should have damn well raise!
	//
	var rows *sql.Rows
	rows, err = s.db.Query("SELECT i FROM events WHERE status='pending' AND ( raise_at < UNIX_TIMESTAMP(NOW()) ) AND raise_at > 0")
	if err != nil {
		return err
	}
	defer rows.Close()

	//
	// For each row in the result-set
	//
	// Record the alert-ID in a temporary array.
	//
	for rows.Next() {
		i := 0

		err = rows.Scan(&i)
		if err != nil {
			return err
		}

		//
		// Record the ID in our list.
		//
		toRaise = append(toRaise, i)

		//
		// Mark this event as raised.
		//
		var raised *sql.Stmt
		raised, err = s.db.Prepare("UPDATE events SET notified_at=?,status=?, notify_count=1 WHERE i=?")
		if err != nil {
			return err
		}
		_, err = raised.Exec(time.Now().Unix(), "raised", i)
		if err != nil {
			return err
		}
		raised.Close()
	}

	//
	// For each of the alerts we've got to raise we can
	// now invoke the callback function.
	//
	for _, id := range toRaise {

		//
		// Get the data.
		//
		data, err := s.GetAlert(id)
		if err != nil {
			fmt.Printf("error fetching id: %s\n", err.Error())
			return err
		}

		//
		// Invoke the callback
		//
		if callback != nil {
			go callback(data, config)
		}
	}
	return nil
}

// Notify anew any still-outstanding alerts.
//
// This function handles re-notification for outstanding alerts.
func (s *Alerts) Renotify(callback AlertRaise, config string) error {
	var err error
	if s.db == nil {
		panic("Working with a closed database - bug")
	}

	//
	// The list of IDs to notify.
	//
	var toRaise []int

	//
	// Our query:
	//
	//  * For every alert in the raised state
	//  * If it was last notified more than a minute ago, renotify
	//
	var rows *sql.Rows
	rows, err = s.db.Query("SELECT i FROM events WHERE status='raised' AND ( abs( notified_at - UNIX_TIMESTAMP(NOW()) ) >= 60 )")
	if err != nil {
		return err
	}
	defer rows.Close()

	//
	// For each row in the result-set
	//
	// Record the IDs
	//
	for rows.Next() {
		i := 0

		err = rows.Scan(&i)
		if err != nil {
			return err
		}

		//
		// Record the ID of the thing we need to renotify.
		//
		toRaise = append(toRaise, i)

		//
		// Update the notification time AND the count of how many
		// times this alert has been notified about.
		//
		var raised *sql.Stmt
		raised, err = s.db.Prepare("UPDATE events SET notified_at=?,notify_count=notify_count+1 WHERE i=?")
		if err != nil {
			return err
		}
		_, err = raised.Exec(time.Now().Unix(), i)
		if err != nil {
			return err
		}
		raised.Close()
	}

	//
	// For each of the alerts we need to notify we can now
	// invoke our callback function to do that.
	//
	for _, id := range toRaise {
		data, err := s.GetAlert(id)
		if err != nil {
			fmt.Printf("error fetching id: %s\n", err.Error())
			return err
		}

		if callback != nil {
			go callback(data, config)
		}
	}
	return nil
}

//
// If an alert is in a raised state, but the `raise_at` time is in the
// future then we can clear it.
//
// This allows heartbeat alerts to auto-clear when they return.
//
func (s *Alerts) Warp() error {

	if s.db == nil {
		panic("Working with a closed database - bug")
	}

	//
	// The query we'll execute.
	//
	//  If there is a time to raise.
	//  and that time is in teh future
	//    (i.e. not the past)
	//  then the alert should be "pending".
	//
	clean, err := s.db.Prepare("UPDATE events SET status='pending' WHERE ( raise_at > UNIX_TIMESTAMP(NOW()) ) AND raise_at > 0")
	if err != nil {
		return err
	}
	defer clean.Close()

	_, err = clean.Exec()
	if err != nil {
		return err
	}

	return nil
}

//
// User/Pass setup
//
func (s *Alerts) ValidateLogin(user string, pass string) (bool, error) {
	if s.db == nil {
		panic("Working with a closed database - bug")
	}

	//
	// The password we're going to fetch
	//
	password := ""

	//
	//
	rows, err := s.db.Query("SELECT password FROM users WHERE username=?", user)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&password)
		if err != nil {
			return false, err
		}
	}
	err = rows.Err()
	if err != nil {
		return false, err
	}

	//
	// Now we have a hashed password, so we need to compare it.
	//
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(pass))
	if err != nil {
		return false, nil
	}
	return true, nil
}

// AddUser adds a new user to the database
func (s *Alerts) AddUser(user string, pass string) error {
	if s.db == nil {
		panic("Working with a closed database - bug")
	}

	//
	// We're going to store a hashed password.
	//
	// So we need to create the hash.
	//
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		return err
	}

	stmt, err := s.db.Prepare("INSERT INTO users( username, password ) VALUES( ?, ? )")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user, hash)
	if err != nil {
		return err
	}

	return nil
}

// DelUser deletes a user from the system.
func (s *Alerts) DelUser(user string) error {
	if s.db == nil {
		panic("Working with a closed database - bug")
	}

	//
	// The query we'll execute.
	//
	clean, err := s.db.Prepare("DELETE FROM users WHERE username=?")
	if err != nil {
		return err
	}
	defer clean.Close()

	_, err = clean.Exec(user)
	if err != nil {
		return err
	}

	return nil
}

// GetUsers returns all known usernames.
func (s *Alerts) GetUsers() ([]string, error) {

	if s.db == nil {
		panic("Working with a closed database - bug")
	}

	//
	// Our return-result.
	//
	var result []string

	//
	// Get the usernames
	//
	rows, err := s.db.Query("SELECT username FROM users")
	if err != nil {
		return result, err
	}
	defer rows.Close()

	//
	// For each row in the result-set parse into an instance of our
	// structure and add to the list.
	//
	for rows.Next() {
		var tmp string
		err := rows.Scan(&tmp)
		if err != nil {
			return result, err
		}

		result = append(result, tmp)
	}
	err = rows.Err()
	if err != nil {
		return result, err
	}
	return result, nil
}
