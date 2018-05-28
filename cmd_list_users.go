//
// List our users
//

package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"github.com/skx/purppura/alerts"
)

type listUsersCmd struct {
}

//
// Glue
//
func (*listUsersCmd) Name() string     { return "list-users" }
func (*listUsersCmd) Synopsis() string { return "List our existing users." }
func (*listUsersCmd) Usage() string {
	return `list-users :
  Show all the existing users.
`
}

//
// Flag setup
//
func (p *listUsersCmd) SetFlags(f *flag.FlagSet) {
}

//
// Entry-point.
//
func (p *listUsersCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	storage, err := alerts.New()
	if err != nil {
		fmt.Printf("Creating database-handle failed: %s\n", err.Error())
	}

	var usernames []string
	usernames, err = storage.GetUsers()
	if err != nil {
		fmt.Printf("Getting the user-list failed: %s\n", err.Error())
	}

	for _, ent := range usernames {
		fmt.Printf("%s\n", ent)
	}

	return subcommands.ExitSuccess
}
