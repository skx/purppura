// Package alert contains our alert-structure details.
package alert

//
// Alert is the structure which is used to describe a single event.
//
// Events are submitted to purppura, and when they are raised they
// become alerts.
//
type Alert struct {
	//
	// The detail of the event, which is formatted as HTML
	// in the user-interface.  Note that we protect against
	// XSS attacks via filtering.
	//
	Detail string

	//
	// The human-readable ID of the event, such as "heartbeat", "backup",
	// or "unread-mail-$sender".
	//
	// The ID doesn't need to be globally-unique, just unique per
	// source IP address.
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

	// Status of the alert
	Status string

	// When an alert will raise at
	RaiseAt string

	// Last notification time recorded here.
	NotifiedAt string

	// Number of times this alert has been notified.
	NotifyCount string
}
