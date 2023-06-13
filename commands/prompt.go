package commands

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

type PromptCommand struct{}

func (c *PromptCommand) Name() string             { return "prompt" }
func (c *PromptCommand) Synopsis() string         { return "Create Prompt" }
func (c *PromptCommand) Usage() string            { return "prompt" }
func (c *PromptCommand) SetFlags(f *flag.FlagSet) {}

func (c *PromptCommand) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {

	return subcommands.ExitSuccess
}
