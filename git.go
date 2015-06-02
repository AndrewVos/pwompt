package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func gitBranch() string {
	output, err := exec.Command("git", "symbolic-ref", "--short", "HEAD").Output()
	if err != nil {
		return "(" + err.Error() + ")"
	}
	branch := string(output)
	branch = strings.TrimSpace(branch)
	return branch
}

func isGitRepository() bool {
	_, err := os.Stat(".git/config")
	return err == nil
}

func isGitRepositoryDirty() bool {
	cmd := exec.Command("git", "status", "--porcelain")
	output, err := cmd.Output()
	if err != nil {
		log.Println(err)
		return false
	}
	lines := strings.Split(string(output), "\n")
	if len(lines) > 1 {
		return true
	}

	return false
}
