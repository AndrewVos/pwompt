package main

import (
	"fmt"
	"os"
	"regexp"
)

const (
	DefaultPwomptConfig = `battery_charging?("white", "⏚")not_battery_charging?("white", "⌁")battery_percentage("red", "yellow", "green")battery?("white", " ")c("yellow", "[")cwd_short("blue")c("yellow", "] ")git_branch("red")git_dirty?("red", "* ")last_exit_code("magenta")last_exit_failed?("white", " ")git?("white", "±")not_git?("white", "$")c("white", " ")`
)

type If struct {
	Name         string
	Result       func() bool
	Precondition func() bool
}

func (i If) IfName() string {
	return i.Name + "?"
}

func (i If) IfNotName() string {
	return "not_" + i.Name + "?"
}

type Method struct {
	Name   string
	Result func(colour string) string
}

func main() {
	pwomptConfig := os.Getenv("PWOMPT_CONFIG")
	if pwomptConfig == "" {
		pwomptConfig = DefaultPwomptConfig
	}

	battery := &Battery{}

	methodFinder := regexp.MustCompile(`([\w?!]+)\(([^\(\)]*)\)`)
	argumentFinder := regexp.MustCompile(`"([^"]*)"`)

	ifs := []If{
		If{Name: "git", Result: isGitRepository},
		If{Name: "git_dirty", Precondition: isGitRepository, Result: isGitRepositoryDirty},
		If{Name: "last_exit_failed", Result: func() bool { return lastExitCode() > 0 }},
		If{Name: "battery", Result: battery.Exists},
		If{Name: "battery_charging", Precondition: battery.Exists, Result: battery.Charging},
	}
	methods := []Method{
		Method{Name: "cwd", Result: func(colour string) string {
			return Colourise(colour, workingDirectory())
		}},
		Method{Name: "cwd_short", Result: func(colour string) string {
			return Colourise(colour, shortWorkingDirectory())
		}},
		Method{Name: "git_branch", Result: func(colour string) string {
			if isGitRepository() {
				return Colourise(colour, gitBranch())
			}
			return ""
		}},
		Method{Name: "last_exit_code", Result: func(colour string) string {
			if code := lastExitCode(); code > 0 {
				return Colourise(colour, fmt.Sprintf("%v", code))
			}
			return ""
		}},
	}

	execute := func(methodName string, arguments []string) bool {
		for _, i := range ifs {
			if methodName == i.IfName() || methodName == i.IfNotName() {
				if i.Precondition == nil || i.Precondition() {
					result := i.Result()
					if methodName == i.IfName() {
						if result {
							fmt.Print(Colourise(arguments[0], arguments[1]))
						}
						return true
					} else {
						if !result {
							fmt.Print(Colourise(arguments[0], arguments[1]))
						}
						return true
					}
				}
			}
		}
		for _, m := range methods {
			if methodName == m.Name {
				fmt.Print(m.Result(arguments[0]))
				return true
			}
		}
		return false
	}

	groups := methodFinder.FindAllStringSubmatch(pwomptConfig, -1)
	for _, group := range groups {
		methodName := group[1]
		var arguments []string
		argumentMatches := argumentFinder.FindAllStringSubmatch(group[2], -1)
		for _, argumentMatch := range argumentMatches {
			arguments = append(arguments, argumentMatch[1])
		}

		if execute(methodName, arguments) {
			continue
		}

		if methodName == "c" {
			fmt.Print(Colourise(arguments[0], arguments[1]))
		} else if methodName == "battery_percentage" {
			if battery.Exists() {
				percentage := battery.Percentage()
				percentageDisplay := fmt.Sprintf("%v", percentage) + "%"

				if percentage < 10 {
					fmt.Print(Colourise(arguments[0], percentageDisplay))
				} else if percentage < 70 {
					fmt.Print(Colourise(arguments[1], percentageDisplay))
				} else {
					fmt.Print(Colourise(arguments[2], percentageDisplay))
				}
			}
		}
	}
}
