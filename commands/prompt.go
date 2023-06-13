package commands

import (
	"context"
	"flag"
	"os"
	"os/user"
	"strings"

	"github.com/fatih/color"
	"github.com/google/subcommands"
	"github.com/isksss/go-prompt/config"
)

type PromptCommand struct{}

func (c *PromptCommand) Name() string             { return "prompt" }
func (c *PromptCommand) Synopsis() string         { return "Create Prompt" }
func (c *PromptCommand) Usage() string            { return "prompt" }
func (c *PromptCommand) SetFlags(f *flag.FlagSet) {}

func (c *PromptCommand) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	var PROMPT string
	PROMPT = getUserName() + " @ " + getCurrentDir() + " $"
	color.New(color.FgCyan).Add(color.BgHiWhite).Printf(PROMPT)
	return subcommands.ExitSuccess
}

// ユーザ名取得
func getUserName() string {
	userName, err := user.Current()
	if err != nil {
		return ""
	}
	return userName.Username
}

// カレントディレクトリ取得
func getCurrentDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		return ""
	}

	home := config.GetHomeDir()
	inHome := strings.Contains(pwd, home)

	if inHome {
		pwd = strings.Replace(pwd, home, "~", 1)
	}
	return pwd
}
