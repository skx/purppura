# Purppura

Purppura (Finnish for "purple") is a port to [golang](https://golang.org/) of the [purple-alert](https://github.com/skx/purple) server.

The purple-alert software is an alert manager which allows the centralised collection and distribution of "alerts".

For example a trivial heartbeat-style alert might be implemented by having a host send a message every minute with a body containing:

* "Raise an alert if you don't hear from me in 5 minutes".

If that host were to suffer a crash then five minutes after the last submission and alert would be raised, and you would be notified.


## Installation

To install the software run:

     ~ $ go get github.com/skx/purppura

Once installed you'll be able to launch the server like so:

      ~ $ purppura
      Listening on http://localhost:8080/

**NOTE** however that the server presents a web interface which requires a login, so you'll need to add at least one user to the system.  This can be done while the server is running, or before you launch it:

      ~ $ purppura -user-add steve
      Enter Password for user steve: secret
      ~ $

Now you should be able to login to the web interface with username `steve` and password `secret`.

The usernames and passwords which will grant access to the web interface are stored inside the same SQLite file that is used to hold event-details.



# Alerts

Alerts are submitted by making a HTTP POST-request to the purple-server, with a JSON-payload of a [number of fields](ALERTS.md).

When a new POST request is received it will be transformed into an alert:

* If the alert is new it will be saved into the database.
* If the alert has been previously seen, then the fields of that existing alert will be updated.
     * This is possible because alerts are uniquely identified by a combination of the submitted `id` field and the source IP address from which it was received.

Alerts have several states:

* Pending.
   * An alert will raise at some point in the future.
* Raised.
   * A raised alert will trigger a notification every **minute** to inform your sysadmin(s).
* Acknowledged
   * An alert in the acknowledged state will not re-notify.
* Cleared
   * Alerts which are cleared have previously been raised but have now cleared.
   * Alerts in the cleared-state are reaped over time.

Submissions are expected to be JSON-encoded POST payloads, sent
to the http://1.2.3.4:port/events end-point.  The required fields are:

|Field Name | Purpose                                                   |
|-----------|-----------------------------------------------------------|
|id         | Name of the alert                                         |
|subject    | Human-readable description of the alert-event.            |
|detail     | Human-readable (expanded) description of the alert-event. |
|raise      | When this alert should be raised.                         |

Further details are available in the [alert guide](ALERTS.md).

## Notifications

The web-based user-interface lists alerts which are pending, raised, or acknowledges.  While this is useful it isn't going to wake anybody up if something fails overnight, so we have to allow notification via SMS, WhatsApp, etc.

There is no built-in facility for sending text-messages, sending pushover notifications, or similar.  Instead the default alerting behaviour is to simply pipe any alert which is in the raised state into an external binary.

* `purple-notify`
   * Executed the _first_ time an alert is raised.
* `purple-renotify`
   * Executed once per minute while an alert continues to be raised.

By moving the notification into an external process you gain the flexibility
to route alerts to humans in whichever way seems best to you.
