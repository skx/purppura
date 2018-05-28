//
// Add a new user
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

type addUserCmd struct {
}

//
// Glue
//
func (*addUserCmd) Name() string     { return "add-user" }
func (*addUserCmd) Synopsis() string { return "Add a new user." }
func (*addUserCmd) Usage() string {
	return `add-user :
  Add a new user to the system.
`
}

//
// Flag setup
//
func (p *addUserCmd) SetFlags(f *flag.FlagSet) {
}

//
// Entry-point.
//
func (p *addUserCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter Username: ")
	user, _ := reader.ReadString('\n')
	user = strings.TrimSpace(user)

	fmt.Printf("Enter Password: ")
	pass, _ := reader.ReadString('\n')
	pass = strings.TrimSpace(pass)

	storage, err := alerts.New()
	if err != nil {
		fmt.Printf("Creating database-handle failed: %s\n", err.Error())
		os.Exit(1)
	}
	err = storage.AddUser(user, pass)
	if err != nil {
		fmt.Printf("Adding user failed: %s\n", err.Error())
	}

	return subcommands.ExitSuccess
}
