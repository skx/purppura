package alert

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
	Status      string
	RaiseAt     string
	NotifiedAt  string
	NotifyCount string
}
