package commands

import (
	"context"
	"flag"
	"fmt"
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

const (
	allow         = ""
	allow2        = ""
	gitBranchMark = ""
)

func (c *PromptCommand) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	var PROMPT string
	var PROMPT2 string
	currentDir := getCurrentDir()
	PROMPT = makeLine(currentDir)

	if config.IsGitRepository() {
		// git user
		gitUser, _ := config.GetGitUser()
		PROMPT2 += makeLine(gitUser)
		// git branch
		currentBranch, _ := config.GetCurrentBranch()
		PROMPT += makeLine(gitBranchMark + " " + currentBranch)
	}
	co := color.New(color.FgBlack).Add(color.BgHiWhite)
	line := co.Sprintf(PROMPT)
	line2 := co.Sprintf(PROMPT2)
	lastAllow := color.New(color.FgHiWhite).Sprintf(allow2 + " \n")

	fmt.Printf(line + lastAllow + line2 + color.New(color.FgHiWhite).Sprintf(allow2))

	return subcommands.ExitSuccess
}

// MakePart
func makeLine(s string) string {
	line := " " + s + " " + allow
	return line
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
