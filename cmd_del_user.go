//
// Delete an existing user
//

package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/google/subcommands"
	"github.com/skx/purppura/alerts"
)

type delUserCmd struct {
}

//
// Glue
//
func (*delUserCmd) Name() string     { return "del-user" }
func (*delUserCmd) Synopsis() string { return "Delete an existing user." }
func (*delUserCmd) Usage() string {
	return `delete-user :
  Remove a user from the system.
`
}

//
// Flag setup
//
func (p *delUserCmd) SetFlags(f *flag.FlagSet) {
}

//
// Entry-point.
//
func (p *delUserCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter Username: ")
	user, _ := reader.ReadString('\n')
	user = strings.TrimSpace(user)

	storage, err := alerts.New()
	if err != nil {
		fmt.Printf("Creating database-handle failed: %s\n", err.Error())
		os.Exit(1)
	}
	err = storage.DelUser(user)
	if err != nil {
		fmt.Printf("Deleting the user failed: %s\n", err.Error())
	}

	return subcommands.ExitSuccess
}
