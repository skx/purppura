//
// This file contains the code to notify users of outstanding alerts.
//

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/skx/purppura/alert"
	"github.com/skx/purppura/alerts"
)

var (
	NotifyBinary = "moi"
)

//
// ProcessAlertsScheduler will ensure that our alerts are processed
// regularly.
//
func ProcessAlertsScheduler(cmd string) {
	for range time.Tick(time.Second * 11) {
		err := ProcessAlerts(cmd)
		if err != nil {
			fmt.Printf("Error Processing Alerts: %s\n", err.Error())
		}
	}
}

//
// ProcessAlerts handles the state-transitions for alerts.
//
func ProcessAlerts(cmd string) error {

	fmt.Printf("Processing events at %s\n", time.Now())

	//
	// Connect to the database to get our alert(s)
	//
	helper, err := alerts.New()
	if err != nil {
		return err
	}

	//
	// Reap any events which have expired.
	//
	err = helper.Reap()
	if err != nil {
		return err
	}

	//
	// Lets do the time-warp, again!
	//
	// (This is required, although with some restructuring it could
	// be omitted.)
	//
	err = helper.Warp()
	if err != nil {
		return err
	}

	//
	// Notify outstanding alerts.
	//
	err = helper.Notify(NotifyAlert, cmd)
	if err != nil {
		return err
	}

	//
	// Renotify any outstanding alerts
	//
	err = helper.Renotify(NotifyAlert, cmd)
	if err != nil {
		return err
	}

	//
	// Close the database-connection now.
	//
	helper.Close()
	fmt.Printf("Processing events complete: %s\n", time.Now())
	return nil
}

//
// Execute the given command with the Alert-event piped to it
// on STDIN, as JSON.
//
func ExecWithAlert(command string, event alert.Alert) error {

	cmd := strings.Split(command, " ")
	login := exec.Command(cmd[0], cmd[1:]...)

	buffer := bytes.Buffer{}
	input, _ := json.Marshal(event)
	buffer.Write(input)

	login.Stdout = os.Stdout
	login.Stdin = &buffer
	login.Stderr = os.Stderr

	err := login.Run()
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	return nil

}

//
// Send a notification for an alert which has become raised.
//
func NotifyAlert(event alert.Alert, config string) error {
	return (ExecWithAlert(config, event))
}
