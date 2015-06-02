package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path"
	"strconv"
	"strings"
	"syscall"
)

var shortenPath = true

func init() {
	flag.BoolVar(&shortenPath, "shorten-path", shortenPath, "shorten the path from \"/some/path/name\" to \"/some/.../name\"")
}

func main() {
	flag.Parse()

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
		branch = Red(branch + gitMarkDirty)
	} else {
	}
	branch = Green(branch) + White(gitMark)

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
	if shortenPath {
		parts := strings.Split(wd, "/")
		if len(parts) > 2 {
			wd = path.Join(parts[0], "...", parts[len(parts)-1])
		}
	}

	wd = Yellow("[" + wd + "]")
	return wd
}

func Black(text string) string   { return colour(text, 0) }
func Red(text string) string     { return colour(text, 1) }
func Green(text string) string   { return colour(text, 2) }
func Yellow(text string) string  { return colour(text, 3) }
func Blue(text string) string    { return colour(text, 4) }
func Magenta(text string) string { return colour(text, 5) }
func Cyan(text string) string    { return colour(text, 6) }
func White(text string) string   { return colour(text, 7) }
func Default(text string) string { return colour(text, 9) }

func colour(text string, colour int) string {
	return "\001\x1b[3" + strconv.Itoa(colour) + ";1m\002" + text + "\001\x1b[0m\002"
}
