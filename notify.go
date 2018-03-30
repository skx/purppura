//
// This file contains the code to notify users of new/repeating
// events which are outstanding.
//

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//
// Execute the given command with the Alert-event piped to it
// on STDIN, as JSON.
//
func ExecWithAlert(command string, event Alert) error {

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
		return err
	}

	return nil

}

//
// Send a notification the first time an alert is raised.
//
func NotifyAlert(event Alert) error {
	fmt.Printf("Notifying for new event %s - via '%s'\n", event.ID, CONFIG.NotifyBinary)

	return (ExecWithAlert(CONFIG.NotifyBinary, event))
}

//
// Send a notification that an alert continues to be raised.
//
func ReNotifyAlert(event Alert) error {
	fmt.Printf("Repeating notification for outstanding event %s - via '%s'\n", event.ID, CONFIG.NotifyBinary)
	return (ExecWithAlert(CONFIG.ReNotifyBinary, event))
}
