package config

import (
	"os"
	"os/exec"
	"strings"
)

func isGitInstalled() bool {
	// Gitコマンドのパスを取得
	_, err := exec.LookPath("git")
	if err != nil {
		return false
	}

	return true
}

func GetCurrentBranch() (string, error) {

	// Gitコマンドを実行してカレントブランチの名前を取得
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	branchName := strings.TrimSpace(string(output))
	return branchName, nil
}

func GetGitUser() (string, error) {

	// Gitコマンドを実行してユーザー名を取得
	cmd := exec.Command("git", "config", "user.name")
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	userName := strings.TrimSpace(string(output))
	return userName, nil
}

func IsGitRepository() bool {
	// Gitコマンドがインストールされているか確認
	if !isGitInstalled() {
		return false
	}

	// Gitコマンドを実行してリポジトリの存在を確認
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return false
	}

	return true
}
