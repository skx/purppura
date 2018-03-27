# Purple.go

This is a port to [golang](https://golang.org/) of the [purple-alert](https://github.com/skx/purple).

All existing documentation should stay the same, except for the notification section.

## Authentication

The authentication requires lookup for a username/password in a flat-file.

By default the flat file is `./users`, but this can be specified via the command-line flag `-auth-file`.  To generate a suitable hash use `util/hash-password.go`:

     deagol ~/p/util $ go run hash-password.go steve
     steve  $2a$14$.lqLUTyghdGGgdEfOMDBzeD8a5pE2t2UsQ1B574ek7uvYgLsXxwLG

## Notifications

The web-based user-interface lists alerts which are pending, raised, or acknowledges.  While this is useful it isn't going to wake anybody up if something fails overnight, so we have to allow notification via SMS, WhatsApp, etc.

The way we do this is to execute external commands to pass on the messages.  Each event will trigger twice:

* `purple-notify`
   * Executed the _first_ time an event is raised.
* `purple-renotify`
   * Executed once per minute while an event continues to be raised.

Both of these commands will receive the event in question, as JSON, piped to STDIN.
