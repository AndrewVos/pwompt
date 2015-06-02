package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/AndrewVos/colour"
)

func main() {
	fmt.Print(workingDirectory())
	fmt.Print(tip())
}

func tip() string {
	tip := " $ "
	return tip
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

	wd = colour.Yellow(wd)
	return wd
}
