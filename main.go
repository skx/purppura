package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
)

//
// Setup our sub-commands and use them.
//
func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&addUserCmd{}, "")
	subcommands.Register(&delUserCmd{}, "")
	subcommands.Register(&listUsersCmd{}, "")
	subcommands.Register(&serveCmd{}, "")
	subcommands.Register(&versionCmd{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
