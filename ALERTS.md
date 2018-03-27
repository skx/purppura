# Alerts

To generate an alert you first need to pick an ID.

IDs are human readable labels for the alerts, and they don't need to be globally unique.  Instead alerts are keyed upon the ID __and__ the source IP address from which they were received.

Example alert IDs might be `disk-full`, `heartbeat`, `unread-mail-foo.com` and `unread-mail-bar.com`.



## Submitting Alerts

Submissions are expected to be JSON-encoded POST payloads, sent
to the http://1.2.3.4:port/events end-point.

The required fields are:

|Field Name | Purpose                                                   |
|-----------|-----------------------------------------------------------|
|id         | Name of the alert                                         |
|subject    | Human-readable description of the alert-event.            |
|detail     | Human-readable (expanded) description of the alert-event. |
|raise      | When this alert should be raised.                         |

As an example the following is a heartbeat alert.  Five minutes after the last update sent by this we'll receive an alert-notification:


     {
       "id"      : "heartbeat"
       "subject" : "The heartbeat wasn't sent for deagol.lan",
       "detail"  : "<p>This indicates that <tt>deagol.lan</tt> might be down!</p>",
       "raise"   : "+5m",
     }

Before the `+5m` timeout has been reached the alert will be in the `pending` state and will be visible in the web user-interface.  Five minutes after the last submission the alert will be moved into the `raised` state, and a notification will be generated.

As you might expect the `raise` field is pretty significant.  Permissable values include:

|`raise`| Purpose                                                 |
|-------|---------------------------------------------------------|
|`12345`| Raise at the given number of seconds past the epoch.    |
| `+5m` | Raise in 5 minutes.                                     |
| `+5h` | Raise in 5 hours.                                       |
| `now` | Raise immediately.                                      |
|`clear`| Clear the alert immediately.                            |

> **NOTE**: Submitting an update which misses any of the expected fields is an error.



## Explicitly Raising Alerts

To raise an alert send a JSON message with the `raise` field set to `now`:

     {
       "id"      : "morning"
       "subject" : "Time to get up!",
       "detail"  : "<p>The time is 5am, you should be awake now.</p>",
       "raise"   : "now",
     }

This alert will be immediately raised, and the notifications will repeat until the alert is cleared, or acknowledged via the the web user-interface (or an update is received from the submitter).


## Explicitly Clearing Alerts

To clear an existing alert send a JSON message with the `raise` field set to `clear`:

     {
       "id"      : "morning"
       "subject" : "Time to get up",
       "detail"  : "<p>The time is 5am, you should be awake now.</p>",
       "raise"   : "clear",
     }

**NOTE**: You are forced to submit `detail` and `subject` fields, even though you're clearing the existing alert.


## Self-Clearing Alerts

If you're writing an alert to tell you that a website is down you can bundle up the previous sections as you would expect:

    v = Hash.new()
    v['subject'] = "http://example.com/ is down"
    v['detail']  = "<p>The fetch failed.</p>"
    v['id']      = "web-example.com"

    if site_alive
       v['raise'] = 'clear'
    else
       v['raise'] = 'now'
    end

With code like this you can send an alert which will either have the raise field set to either `now` or `clear`.  Each update will change the state appropriately.


## Heartbeat Alerts

Heartbeat alerts have already been documented, but in brief instead of sending simple "raise" or "clear" events you instead set a relative time in your `raise` field.

For example you could send, every minute, a submission like this:


     {
       "id"      : "heartbeat"
       "subject" : "example.my.flat is down.",
       "detail"  : "<p>The heartbeat wasn't received for example.my.flat.</p>",
       "raise"   : "+5m",
       }

Assuming that this update is sent every 60 seconds the alert will raise five minutes after the last update.  That would require the host was down for five minutes, or that five updates were lost en route.

This works because the `raise` time is updated every time the incoming alert is received.  So when the first update is seen the `+5m` field might expand to the absolute time `Tue Jun 14 09:20:21 EEST 2016`, then a minute later the field will change to `Tue Jun 14 09:21:21 EEST 2016`.  The time at which the alert will raise will be pushed back a minute on each update, unless these updates cease.


# Notifications

By default notifications are repeated for each alert in the raised-state.  These notifications repeat every 60 seconds.

The alerting is divided into two types:

* An initial notification.
  * This is issued once.
* A repeated notification.
  * These repeat notifications are issued once per minute, indefinitely.

This behaviour is useful if you're using an external service to deliver your alert-messages.  For example I use the [pushover](http://pushover.net/) service, and there is a facility there to repeat the notifications until they are read with the mobile phone application.  If I raise the alert once there, the phone will beep every minute - so there is no need to repeatedly send the message.

In my case what I do is configure the server to use a script to raise the event the first time, and the script which _should_ repeat the notifications is a NOP.


## Sample Clients

You can find clients for submitting events beneath [examples/](examples/).
