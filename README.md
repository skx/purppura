[![Go Report Card](https://goreportcard.com/badge/github.com/skx/purppura)](https://goreportcard.com/report/github.com/skx/purppura)
[![license](https://img.shields.io/github/license/skx/purppura.svg)](https://github.com/skx/purppura/blob/master/LICENSE)
[![Release](https://img.shields.io/github/release/skx/purppura.svg)](https://github.com/skx/purppura/releases/latest)

Table of Contents
=================

* [Purppura](#purppura)
* [Alerts](#alerts)
* [Installation](#installation)
  * [Source Installation go &lt;=  1.11](#source-installation-go---111)
  * [Source installation go  &gt;= 1.12](#source-installation-go---112)
  * [Post-Installation Setup](#post-installation-setup)
* [Notifications](#notifications)
* [Links](#links)
* [Docker](#docker)
* [Github Setup](#github-setup)


# Purppura

Purppura is an alert manager which allows the centralised collection and distribution of events or alerts.  (Things submitted are _events_ and they become alerts when they _alert_ a human!)

For example a trivial heartbeat-style alert might be implemented by having a host send a message every minute:

* "Raise an alert if you don't hear from me in 5 minutes".
   * i.e. This is a [dead man's switch](https://en.wikipedia.org/wiki/Dead_man%27s_switch) system.

If that host were to suffer a crash then five minutes after the last submission of the event an alert would be raised, and a human would be notified.


# Alerts

Events are submitted by making a HTTP POST-request to the server, with a JSON-payload containing a [number of fields](ALERTS.md).  When a new POST request is received it will be transformed into an event:

* If the event is new it will be saved into the database.
* If the event has been previously seen, then the fields of that existing entry will be updated.
   * This is possible because events are uniquely identified by a combination of the submitted `id` field and the source IP address from which it was received.

Events have several states:

* Pending.
   * An event might become raised at some point in the future.
* Raised.
   * A raised event will trigger a notification every **minute** to inform your sysadmin(s).
* Acknowledged
   * An alert in the acknowledged state will not re-notify.
   * An event can be acknowledged via the HTTP-server, and it is assumed a human will do that to indicate they're handling the issue.
* Cleared
   * Alerts in the cleared-state are reaped over time.

The required fields for a submission are documented in [ALERTS.md](ALERTS.md),
but in brief you need to submit:

|Field Name | Purpose                                                   |
|-----------|-----------------------------------------------------------|
|id         | Name of the alert                                         |
|subject    | Human-readable description of the alert-event.            |
|detail     | Human-readable (expanded) description of the alert-event. |
|raise      | When this alert should be raised. ("now", "+5m", etc)     |




## Installation

There are two ways to install this project from source, which depend on the version of the [go](https://golang.org/) version you're using.

If you just need the binaries you can find them upon the [project release page](https://github.com/skx/purppura/releases).


### Source Installation go <=  1.11

If you're using `go` before 1.11 then the following command should fetch/update the project, and install it upon your system:

     $ go get -u github.com/skx/purppura

### Source installation go  >= 1.12

If you're using a more recent version of `go` (which is _highly_ recommended), you need to clone to a directory which is not present upon your `GOPATH`:

    git clone https://github.com/skx/purppura
    cd purppura
    go install


### Post-Installation Setup

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



## Notifications

The web-based user-interface lists alerts which are pending, raised, or acknowledged.  While this is useful it isn't going to wake anybody up if something fails overnight, so we have to allow notification via SMS, WhatsApp, etc.

There is no built-in facility for routing notifications to people directly, instead the default alerting behaviour is to simply pipe any event which is in the raised state into a binary called `purppura-notify`.

**NOTE**: Remember that you need to add this script somewhere upon your `PATH`.

* You can find a sample `purppura-notify` beneath [notifiers/](notifiers/).

The notification binary is executed with a JSON-representation of the event
piped to it on STDIN, and will be executed in two situations:

* The first time an event becomes raised.
* Once every minute, as a reminder, as the event continues to be in the raised state.

In addition to the actual event-details the JSON object will have a `NotifyCount` attribute, which will incremented once each time the alert has been piped to the binary.  This allows you to differentiate between the two obvious states:

* The event has become raised for the first time.
* This is a reminder that the event continues to be outstanding, neither fixed nor acknowledged.

Using the count is useful if you're using an external service to deliver your alert-messages which has its own reminder-system.  For example I use the [pushover](http://pushover.net/) service, and there is a facility there to repeat the notifications until they are read with the mobile phone application.

Using the count-facility I configure my alerter to notify Pushover __once__,
and if the event continues to be outstanding I don't need to needlessly repeat the phone-notification.

A second use for the reminder-facility is to allow alerts to be raised through a hierarchy:

* You might notify an individual the first ten times an event is raised.
* But after that you'd switch to notifying their boss too.



## Links

A simpler implementation of a similar idea is nanny:

* https://github.com/lunemec/nanny/

As an example of something that extensively communicates with `purppura` please see my network monitoring tool:

* https://github.com/skx/overseer

`overseer` carries out network testing, and submits the results of each test to a central purppura instance - automatically raising/clearing alerts as systems and services come and go.  The notification system that `overseer` provides is [very flexible](https://github.com/skx/overseer/#notifications); but I use purppura exclusively.



## Docker

Building the Docker image is as simple as you would expect:

```
$ docker build -t purppura:latest .
```

However note that when it comes to _deployment_ there are a couple of complications:

* We need the MySQL database to be created and populated.
  * Due to this you'll need to copy the `purppura.sql` file to the host you're running on.
* We need to have an external binary to issue the notifications.
  * When an alert is raised an executable will be launched, with the alert details piped to STDIN.
  * So you'll need to bind-mount the notification script to `/srv/bin/purppura-notify`.

On the host you're running the application upon that means you'll need these three files:

* `docker-compose.yml`
  * The helper to launch the containers.
* `purppura-notify`
  * The script executed when an alert is raised.
* `purppura.sql`
  * The database file.

Assuming those files are present:

```
$ docker-compose up -d
```

Now you can add your user:

```
$ docker exec -t -i purppura_purppura_1 /app/purppura add-user
```

After the first run you won't need the `purppura.sql` file, as the MySQL state will be persisted locally, but leaving it in-place is safest.



## Github Setup

This repository is configured to run tests upon every commit, and when
pull-requests are created/updated.  The testing is carried out via
[.github/run-tests.sh](.github/run-tests.sh) which is used by the
[github-action-tester](https://github.com/skx/github-action-tester) action.

Releases are automated in a similar fashion via [.github/build](.github/build),
and the [github-action-publish-binaries](https://github.com/skx/github-action-publish-binaries) action.
