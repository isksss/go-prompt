package commands

import (
	"context"
	"flag"

	"github.com/fatih/color"
	"github.com/google/subcommands"
	"github.com/isksss/go-prompt/config"
)

type InitCommand struct{}

func (c *InitCommand) Name() string             { return "init" }
func (c *InitCommand) Synopsis() string         { return "Init config" }
func (c *InitCommand) Usage() string            { return "init [args]" }
func (c *InitCommand) SetFlags(f *flag.FlagSet) {}

func (c *InitCommand) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	cCyan := color.New(color.FgCyan)
	cWhite := color.New(color.FgWhite)

	cCyan.Printf("Init config\n")

	filepath, err := config.MakeConfigFile()
	if err != nil {
		return subcommands.ExitFailure
	}
	cWhite.Printf("Config file path: %s\n", filepath)
	config.CreateDefaultConfig()

	cCyan.Printf("Done\n")
	return subcommands.ExitSuccess
}
