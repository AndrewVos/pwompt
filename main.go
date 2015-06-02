package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"syscall"

	"github.com/AndrewVos/colour"
)

func main() {
	fmt.Print(workingDirectory())
	fmt.Print(tip())
}

func tip() string {
	tip := " $ "
	if isGitRepository() {
		tip = gitTip()
	}
	return tip
}

func isGitRepository() bool {
	_, err := os.Stat(".git/config")
	return err == nil
}

var gitMark = " ± "
var gitMarkDirty = "*"

func gitTip() string {
	output, err := exec.Command("git", "symbolic-ref", "--short", "HEAD").Output()
	if err != nil {
		return "(" + err.Error() + ")"
	}
	branch := string(output)
	branch = strings.TrimSpace(branch)

	branch = " " + branch
	if isGitRepositoryDirty() {
		branch = colour.Red(branch + gitMarkDirty)
	} else {
	}
	branch = colour.Green(branch) + colour.White(gitMark)

	return branch
}

func isGitRepositoryDirty() bool {
	cmd := exec.Command("git", "diff-files", "--quiet")

	if err := cmd.Run(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if _, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				return true
			}
			return false
		} else {
			return true
		}
	}

	return false
}

func workingDirectory() string {
	wd, err := os.Getwd()
	if err != nil {
		wd = ""
	}

	usr, err := user.Current()
	if err == nil {
		if wd == usr.HomeDir {
			wd = "~"
		} else if strings.HasPrefix(wd, usr.HomeDir) {
			wd = strings.TrimPrefix(wd, usr.HomeDir)
			wd = "~" + wd
		}
	}

	wd = colour.Yellow("[" + wd + "]")
	return wd
}
