package main

import "os"

func User() string {
	return os.Getenv("USER")
}
