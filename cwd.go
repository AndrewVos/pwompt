package main

import (
	"os"
	"os/user"
	"path"
	"strings"
)

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

	return wd
}

func shortWorkingDirectory() string {
	wd := workingDirectory()
	parts := strings.Split(wd, "/")
	if len(parts) > 2 {
		wd = path.Join(parts[0], "...", parts[len(parts)-1])
	}
	return wd
}
