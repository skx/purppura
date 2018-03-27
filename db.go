//
// This file contains all our alert-specific routines for interfacing
// with our SQLite database
//

package main

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

//
// The global DB handle.
//
var db *sql.DB

//
// Setup the database, creating it if it is missing.
//
func SetupDB(path string) error {

	var err error

	//
	// Return if the database already exists.
	//
	db, err = sql.Open("sqlite3", path)
	if err != nil {
		return err
	}

	//
	// Create the table.
	//
	sqlStmt := `

        PRAGMA automatic_index = ON;
        PRAGMA cache_size = 32768;
        PRAGMA journal_size_limit = 67110000;
        PRAGMA locking_mode = NORMAL;
        PRAGMA synchronous = NORMAL;
        PRAGMA temp_store = MEMORY;
        PRAGMA journal_mode = WAL;
        PRAGMA wal_autocheckpoint = 16384;

       CREATE TABLE events (
           i INTEGER PRIMARY KEY,
          id    text not null,
         source text not null,
         status char(10) DEFAULT 'pending',
        raise_at int default '0',
        notified_at int default '0',
        subject text not null,
        detail  text not null
       )
	`

	//
	// Create the table, if missing.
	//
	// Errors here are pretty unlikely.
	//
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	return nil
}

//
// Add a new event into the database
//
func addEvent(data Alert) error {

	//
	// Ensure we have a DB-handle
	//
	if db == nil {
		return errors.New("SetupDB not called")
	}

	//
	// See if we've seen this alert before
	//
	id := -1

	row := db.QueryRow("SELECT i FROM Events WHERE id=? AND source=?", data.ID, data.Source)
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			id = -1
		} else {
			return err
		}
	}

	raise, err := Str2Unix(data.Raise)
	if err != nil {
		return err
	}

	if id == -1 {
		tx, err := db.Begin()
		if err != nil {
			return err
		}
		stmt, err := tx.Prepare("INSERT INTO Events( id, source, subject, detail, raise_at ) VALUES( ?, ?, ?, ?, ? )")
		if err != nil {
			return err
		}
		defer stmt.Close()

		stmt.Exec(data.ID,
			data.Source,
			data.Subject,
			data.Detail,
			raise)
		tx.Commit()

	} else {

		// This is updating an existing alert
		tx, err := db.Begin()
		if err != nil {
			return err
		}
		up, err := tx.Prepare("UPDATE Events SET raise_at=?, subject=?, detail=?  WHERE i=?")
		if err != nil {
			return err
		}

		up.Exec(raise, data.Subject, data.Detail, id)
		tx.Commit()

	}

	return nil
}

//
// Return all alerts, and their details.
//
func Alerts() ([]Alert, error) {

	//
	// Our return-result.
	//
	var results []Alert

	//
	// Ensure we have a DB-handle
	//
	if db == nil {
		return nil, errors.New("SetupDB not called")
	}

	//
	// Select the status - for nodes seen in the past 24 hours.
	//
	rows, err := db.Query("SELECT i,source,status,subject,detail,raise_at, notified_at from events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//
	// For each row in the result-set parse into an instance of our
	// structure and add to the list.
	//
	for rows.Next() {
		var tmp Alert

		err := rows.Scan(&tmp.ID, &tmp.Source, &tmp.Status, &tmp.Subject, &tmp.Detail, &tmp.RaiseAt, &tmp.NotifiedAt)
		if err != nil {
			return nil, err
		}

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
// Return a single alert
//
func GetAlert(id int) (Alert, error) {

	//
	// Our return-result.
	//
	var result Alert

	//
	// Ensure we have a DB-handle
	//
	if db == nil {
		return result, errors.New("SetupDB not called")
	}

	//
	// Select the status - for nodes seen in the past 24 hours.
	//
	rows, err := db.Query("SELECT i,source,status,subject,detail,raise_at, notified_at from events WHERe i=?", id)
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
// Change the given event ID to have the specified state.
//
func SetEvent(id string, state string) error {
	//
	// Ensure we have a DB-handle
	//
	if db == nil {
		return errors.New("SetupDB not called")
	}

	stmt, err := db.Prepare("UPDATE events SET status=? WHERE i=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(state, id)
	if err != nil {
		return err
	}
	return nil
}

//
// Change the state of the given alert to be "acknowledged".
//
func AckEvent(id string) error {
	return (SetEvent(id, "acknowledged"))
}

//
// Change the state of the given alert to be "cleared".
//
func ClearEvent(id string) error {
	return (SetEvent(id, "cleared"))
}

//
// Change the state of the given alert to be "raised".
//
func RaiseEvent(id string) error {
	return (SetEvent(id, "raised"))
}

//
// Remove old/obsolete events from the database.
//
func Reap() error {

	//
	// Ensure we have a DB-handle
	//
	if db == nil {
		return errors.New("SetupDB not called")
	}

	//
	// The query we'll execute.
	//
	clean, err := db.Prepare("DELETE FROM events WHERE status='cleared' OR raise_at < 1")
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
func Notify() error {

	//
	// Ensure we have a DB-handle
	//
	if db == nil {
		return errors.New("SetupDB not called")
	}

	//
	// Our query.
	//
	rows, err := db.Query("SELECT i FROM events WHERE status='pending' AND ( raise_at < strftime('%s','now') ) AND raise_at > 0")
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

		err := rows.Scan(&i)
		if err != nil {
			return err
		}

		//
		// Get the event we're going to notify.
		//
		data, err := GetAlert(i)
		if err != nil {
			return err
		}

		//
		// Execute our alerting-binary, giving it the alert
		// as JSON on STDIN.
		//
		NotifyAlert(data)

		//
		// Mark this as raised
		//
		raised, err := db.Prepare("UPDATE events SET notified_at=?,status=? WHERE i=?")
		if err != nil {
			return err
		}
		_, err = raised.Exec(time.Now().Unix(), "raised", i)
		if err != nil {
			return err
		}
	}

	return nil
}

//
// ReNotify alerts which continue to be outstanding.
//
func ReNotify() error {

	//
	// Ensure we have a DB-handle
	//
	if db == nil {
		return errors.New("SetupDB not called")
	}

	//
	// Our query - find events which have been outstanding for more
	// than 60 seconds.
	//
	rows, err := db.Query("SELECT i FROM events WHERE status='raised' AND ( abs( notified_at - strftime('%s','now') ) >= 60 )")
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

		err := rows.Scan(&i)
		if err != nil {
			return err
		}

		//
		// Get the event we're going to notify.
		//
		data, err := GetAlert(i)
		if err != nil {
			return err
		}

		//
		// Execute our alerting-binary, giving it the alert
		// as JSON on STDIN.
		//
		ReNotifyAlert(data)

		//
		// Mark this as raised
		//
		raised, err := db.Prepare("UPDATE events SET notified_at=? WHERE i=?")
		if err != nil {
			return err
		}
		_, err = raised.Exec(time.Now().Unix(), i)
		if err != nil {
			return err
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
func Warp() error {

	//
	// Ensure we have a DB-handle
	//
	if db == nil {
		return errors.New("SetupDB not called")
	}

	//
	// The query we'll execute.
	//
	clean, err := db.Prepare("UPDATE events SET status='pending' WHERE ( raise_at > strftime('%s','now') ) AND raise_at > 0")
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
