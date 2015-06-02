package main

import (
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type Battery struct {
	acpiOutput      string
	acpiOutputError error
}

func (b *Battery) retrieveAcpiOutput() {
	output, err := exec.Command("acpi", "--battery").Output()
	if err != nil {
		b.acpiOutputError = err
		log.Println(err)
	}
	b.acpiOutput = string(output)
}

func (b *Battery) Exists() bool {
	if b.acpiOutput == "" && b.acpiOutputError == nil {
		b.retrieveAcpiOutput()
	}
	return b.acpiOutputError == nil
}

func (b *Battery) Charging() bool {
	if b.acpiOutput == "" && b.acpiOutputError == nil {
		b.retrieveAcpiOutput()
	}
	return strings.Contains(b.acpiOutput, "Charging")
}

func (b *Battery) Percentage() int {
	if b.acpiOutput == "" && b.acpiOutputError == nil {
		b.retrieveAcpiOutput()
	}
	percentageRegex := regexp.MustCompile("(\\d+)%")
	matches := percentageRegex.FindStringSubmatch(b.acpiOutput)
	percentage, err := strconv.ParseInt(matches[1], 0, 64)
	if err != nil {
		log.Println(err)
	}
	return int(percentage)
}
