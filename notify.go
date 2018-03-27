//
// This file contains the code to notify users of new/repeating
// events whcih are outstanding.
//

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

//
// Execute the given command with the Alert-event piped to it
// on STDIN, as JSON.
//
func ExecWithAlert(command string, event Alert) error {

	login := exec.Command(command)

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

func NotifyAlert(event Alert) error {
	fmt.Printf("Notifying for new event %s\n", event.ID)

	return (ExecWithAlert(CONFIG.NotifyBinary, event))
}

func ReNotifyAlert(event Alert) error {
	fmt.Printf("Repeating notification for outstanding event %s\n", event.ID)
	return (ExecWithAlert(CONFIG.ReNotifyBinary, event))
}
