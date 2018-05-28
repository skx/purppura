[![Travis CI](https://img.shields.io/travis/skx/purppura/master.svg?style=flat-square)](https://travis-ci.org/skx/purppura)
[![Go Report Card](https://goreportcard.com/badge/github.com/skx/purppura)](https://goreportcard.com/report/github.com/skx/purppura)
[![license](https://img.shields.io/github/license/skx/purppura.svg)](https://github.com/skx/purppura/blob/master/LICENSE)
[![Release](https://img.shields.io/github/release/skx/purppura.svg)](https://github.com/skx/purppura/releases/latest)
[![gocover store](http://gocover.io/_badge/github.com/skx/purppura)](http://gocover.io/github.com/skx/purppura)

# Purppura

Purppura is an alert manager which allows the centralised collection and distribution of "alerts".

For example a trivial heartbeat-style alert might be implemented by having a host send a message every minute:

* "Raise an alert if you don't hear from me in 5 minutes".

If that host were to suffer a crash then five minutes after the last submission of the event an alert would be raised, and a human would be notified.


## Installation

To install the software run:

     ~ $ go get -u github.com/skx/purppura
     ~ $ go install github.com/skx/purppura

Once installed you'll be ready to launch the server, but first of all you
must create the (MySQL) database and save the connection-details in the
environment.  The definition of the appropriate tables can be found in
the [purppura.sql](purppura.sql) file.

Assuming you're using MySQL on the local-host you can export the details
like so:

      ~ $ export PURPLE_DSN="user:pass@tcp(localhost:3306)/purple?timeout=5s"

Once the environment has the correct details you can now launch the
server:

      ~ $ purppura serve
      Listening on http://localhost:8080/

You'll want to add at least one user who can login to the web-based user-interface.  Users are stored in the database, and can be added/listed/removed  while the server is running:

      ~ $ purppura add-user
      Enter Username: moi
      Enter Password: kissa
      ~ $

> **NOTE**: You must set the `$PURPLE_DSN` environmental-variable for adding, listing, or removing users.

Once the user has been added you should be able to login to the web interface with username `moi` and password `kissa`.

To see your users you can run:

      ~ $ purppura list-users

And to delete a user:

      ~ $ purppura del-user
      Enter Username: moi

You can use the `del-user` sub-command to remove the user in the future, or the `list-users` sub-command to see the users which are present.


# Alerts

Alerts are submitted by making a HTTP POST-request to the server, with a JSON-payload containing a [number of fields](ALERTS.md).

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

The required fields for a submission are:

|Field Name | Purpose                                                   |
|-----------|-----------------------------------------------------------|
|id         | Name of the alert                                         |
|subject    | Human-readable description of the alert-event.            |
|detail     | Human-readable (expanded) description of the alert-event. |
|raise      | When this alert should be raised.                         |

Further details are available in the [alert guide](ALERTS.md) - and you can see some example scripts which submit events, beneath [examples/](examples/).


## Notifications

The web-based user-interface lists alerts which are pending, raised, or acknowledges.  While this is useful it isn't going to wake anybody up if something fails overnight, so we have to allow notification via SMS, WhatsApp, etc.

There is no built-in facility for sending notifications, instead the default alerting behaviour is to simply pipe any alert which is in the raised state into an external binary:

* `purppura-notify`
  * Executed when an alert is raised, or re-raised.
  * Will receive all the details of the alert as a JSON-object on STDIN

By moving the notification into an external process you gain the flexibility
to route alerts to humans in whichever way seems best to you.  You can find a sample notification-script [notifiers/](notifiers/) which alerts a human via pushover.

**NOTE**: Remember that you need to add this script somewhere upon your `PATH`.
