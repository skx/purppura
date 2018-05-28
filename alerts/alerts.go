//
// The Alerts package contains code relating to the persistance and
// manipulation of alerts.
//
//

package alerts

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/microcosm-cc/bluemonday"
	"github.com/skx/purppura/alert"
	"github.com/skx/purppura/util"
)

// Alerts is our object for interfacing with alerts.
type Alerts struct {
	// Storage for our DB-handle, which is a private implementation-detail
	db *sql.DB
}

// AlertRaise is the function signature of something that can raise an
// alert
type AlertRaise func(x alert.Alert, config string) error

// New is the constructor for our object.
func New() (*Alerts, error) {
	var err error

	m := new(Alerts)

	//
	// Create a default configuration structure for MySQL.
	//
	config := mysql.NewConfig()

	//
	// Populate the username & password fields.
	//
	// This all needs to be configurable in the future.
	//
	config.User = "purple"
	config.Passwd = "purple"
	config.DBName = "purple"
	config.Net = "tcp"
	config.Addr = "www.vpn:3306"
	config.Timeout = 5 * time.Second

	//
	// Now convert the connection-string to a DSN, which
	// is used to connect to the database.
	//
	dsn := config.FormatDSN()

	//
	// Show the DSN, if appropriate.
	//
	//	fmt.Printf("\tMySQL DSN is %s\n", dsn)

	//
	// Connect to the database
	//
	m.db, err = sql.Open("mysql", dsn)
	if err != nil {
		return m, err
	}

	return m, err
}

// Close ensures that our database-connection is closed.
func (s *Alerts) Close() {
	s.db.Close()
	s.db = nil
}

//
// AddEvent adds a new event to our persistant storage.
//
func (s *Alerts) AddEvent(data alert.Alert) error {

	if s.db == nil {
		panic("Working with a closed database - bug")
	}

	//
	// See if we've seen this alert before
	//
	id := -1

	//
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

	raise, err := util.Str2Unix(data.Raise)
	if err != nil {
		return err
	}

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

//
// Return all alerts, and their details.
//
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
	// Select the status - for nodes seen in the past 24 hours.
	//
	rows, err := s.db.Query("SELECT i,source,status,subject,detail,raise_at, notified_at from events")
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

		err := rows.Scan(&tmp.ID, &tmp.Source, &tmp.Status, &tmp.Subject, &tmp.Detail, &tmp.RaiseAt, &tmp.NotifiedAt)
		if err != nil {
			return nil, err
		}

		//
		// The detail should be sanitized.
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
	return results, nil
}

//
// GetAlert returns a single alert, via the identifier.
//
func (s *Alerts) GetAlert(id int) (alert.Alert, error) {

	if s.db == nil {
		panic("Working with a closed database - bug")
	}

	//
	// Our return-result.
	//
	var result alert.Alert

	//
	// Select the status - for nodes seen in the past 24 hours.
	//
	rows, err := s.db.Query("SELECT i,source,status,subject,detail,raise_at, notified_at from events WHERe i=?", id)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	//
	// For each row in the result-set parse into an instance of our
	// structure and add to the list.
	//
	for rows.Next() {
		err := rows.Scan(&result.ID, &result.Source, &result.Status, &result.Subject, &result.Detail, &result.RaiseAt, &result.NotifiedAt)
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

//
// setEvent changes the given alert to have the specified state.
//
func (s *Alerts) setEvent(id string, state string) error {
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
	return (s.setEvent(id, "acknowledged"))
}

// ClearEvent changes the state of the given alert to be "cleared".
func (s *Alerts) ClearEvent(id string) error {
	return (s.setEvent(id, "cleared"))
}

// RaiseEvent changes the state of the given alert to be "raised".
func (s *Alerts) RaiseEvent(id string) error {
	return (s.setEvent(id, "raised"))
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

//
// Notify each outstanding alert
//
func (s *Alerts) Notify(callback AlertRaise, config string) error {
	var err error
	if s.db == nil {
		panic("Working with a closed database - bug")
	}

	var toRaise []int

	//
	// Our query.
	//
	//  For every alert in the pending state
	//  if there is a `raise_at` time, and that time is in the past
	// then we should have damn well raised already!
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
	// Parse into a structure and add to the list.
	//
	for rows.Next() {
		i := 0

		err = rows.Scan(&i)
		if err != nil {
			return err
		}

		//
		// Get the event we're going to notify.
		//
		toRaise = append(toRaise, i)

		//
		// Mark this as raised
		//
		var raised *sql.Stmt
		raised, err = s.db.Prepare("UPDATE events SET notified_at=?,status=? WHERE i=?")
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
	// For each of the alerts we've got to raise we should
	// invoke the callback function.
	//
	for _, id := range toRaise {
		data, err := s.GetAlert(id)
		if err != nil {
			fmt.Printf("error fetching id: %s\n", err.Error())
			return err
		}

		if callback != nil {
			callback(data, config)
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
	// The login we're going to fetch
	//
	login := ""

	//
	//
	rows, err := s.db.Query("SELECT username FROM users WHERE username=? AND password=?", user, pass)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&login)
		if err != nil {
			return false, err
		}
	}
	err = rows.Err()
	if err != nil {
		return false, err
	}

	if (login == user) && (login != "") {
		return true, nil
	}
	return false, nil
}

// AddUser adds a new user to the database
func (s *Alerts) AddUser(user string, pass string) error {
	if s.db == nil {
		panic("Working with a closed database - bug")
	}

	stmt, err := s.db.Prepare("INSERT INTO users( username, password ) VALUES( ?, ? )")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user, pass)
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
