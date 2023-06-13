package main

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	"github.com/isksss/go-prompt/commands"
)

func init() {
	subcommands.Register(subcommands.CommandsCommand(), "help")
	subcommands.Register(subcommands.FlagsCommand(), "help")
	subcommands.Register(subcommands.HelpCommand(), "help")

	subcommands.Register(&commands.InitCommand{}, "")
	subcommands.Register(&commands.PromptCommand{}, "")

	flag.Parse()
}
func main() {
	ctx := context.Background()
	subcommands.Execute(ctx)
}
