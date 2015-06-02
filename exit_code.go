package main

import (
	"os"
	"strconv"
)

func lastExitCode() int {
	lastExitCode := os.Getenv("PWOMPT_LAST_EXIT_CODE")
	code, _ := strconv.ParseInt(lastExitCode, 0, 64)
	return int(code)
}
