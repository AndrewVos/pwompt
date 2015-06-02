package main

import (
	"fmt"
	"os"
	"regexp"
)

const (
	DefaultPwomptConfig = `battery_charging?("white", "⏚")battery_discharging?("white", "⌁")battery_percentage("red", "yellow", "green")battery?("white", " ")c("yellow", "[")cwd_short("blue")c("yellow", "] ")git_branch("red")git_dirty?("red", "* ")last_exit_code("magenta")last_exit_failed?("white", " ")git?("white", "±")not_git?("white", "$")c("white", " ")`
)

func main() {
	pwomptConfig := os.Getenv("PWOMPT_CONFIG")
	if pwomptConfig == "" {
		pwomptConfig = DefaultPwomptConfig
	}

	battery := &Battery{}

	methodFinder := regexp.MustCompile(`([\w?!]+)\(([^\(\)]*)\)`)
	argumentFinder := regexp.MustCompile(`"([^"]*)"`)

	groups := methodFinder.FindAllStringSubmatch(pwomptConfig, -1)
	for _, group := range groups {
		method := group[1]
		var arguments []string
		argumentMatches := argumentFinder.FindAllStringSubmatch(group[2], -1)
		for _, argumentMatch := range argumentMatches {
			arguments = append(arguments, argumentMatch[1])
		}
		if method == "c" {
			fmt.Print(colour(arguments[0], arguments[1]))
		} else if method == "cwd" {
			fmt.Print(colour(arguments[0], workingDirectory()))
		} else if method == "cwd_short" {
			fmt.Print(colour(arguments[0], shortWorkingDirectory()))
		} else if method == "battery_charging?" {
			if battery.isBatteryCharging() {
				fmt.Print(colour(arguments[0], arguments[1]))
			}
		} else if method == "battery_discharging?" {
			if !battery.isBatteryCharging() {
				fmt.Print(colour(arguments[0], arguments[1]))
			}
		} else if method == "battery_percentage" {
			percentage := battery.percentage()
			percentageDisplay := fmt.Sprintf("%v", percentage) + "%"

			if percentage < 10 {
				fmt.Print(colour(arguments[0], percentageDisplay))
			} else if percentage < 70 {
				fmt.Print(colour(arguments[1], percentageDisplay))
			} else {
				fmt.Print(colour(arguments[2], percentageDisplay))
			}

		} else if method == "battery?" {
			if battery.batteryExists() {
				fmt.Print(colour(arguments[0], arguments[1]))
			}
		} else if method == "git_branch" {
			if isGitRepository() {
				fmt.Print(colour(arguments[0], gitBranch()))
			}
		} else if method == "git_dirty?" {
			if isGitRepository() && isGitRepositoryDirty() {
				fmt.Print(colour(arguments[0], arguments[1]))
			}
		} else if method == "git?" {
			if isGitRepository() {
				fmt.Print(colour(arguments[0], arguments[1]))
			}
		} else if method == "not_git?" {
			if !isGitRepository() {
				fmt.Print(colour(arguments[0], arguments[1]))
			}
		} else if method == "last_exit_code" {
			if code := lastExitCode(); code > 0 {
				fmt.Print(colour(arguments[0], fmt.Sprintf("%v", code)))
			}
		} else if method == "last_exit_failed?" {
			if lastExitCode() > 0 {
				fmt.Print(colour(arguments[0], arguments[1]))
			}
		}

	}
}
