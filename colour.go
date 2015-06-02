package main

import "fmt"

func colour(colour string, s string) string {
	colours := map[string]int{
		"black":   0,
		"red":     1,
		"green":   2,
		"yellow":  3,
		"blue":    4,
		"magenta": 5,
		"cyan":    6,
		"white":   7,
		"default": 9,
	}

	return "\001\x1b[3" + fmt.Sprintf("%v", colours[colour]) + ";1m\002" + s + "\001\x1b[0m\002"
}
