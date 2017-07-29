package main

import (
	"fmt"
	"os/exec"

	"github.com/fatih/color"
)

var execCommand = exec.Command

// GitLog retrieves the log and returns the formatted text
func GitLog(directory string, formatter string) (string, error) {
	color.Green("Pulling the history: %s", directory)

	commandName := "git"
	customFormatter := "--pretty=format:" + formatter

	args := []string{"log", customFormatter}

	command := execCommand(commandName, args...)
	command.Dir = directory

	output, err := command.CombinedOutput()

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	log := string(output[:])
	return log, nil
}
