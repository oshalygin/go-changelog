package main

import (
	"fmt"
	"os/exec"
	"strings"

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

func prependGitHubDomain(path string) string {
	githubDomain := "https://github.com/"
	return githubDomain + path
}

func trimGitSuffix(path string) string {
	return strings.TrimSuffix(path, ".git")
}

func getRepositoryURL(remote string) string {
	if strings.HasPrefix(remote, "https") {
		return trimGitSuffix(remote)
	}
	return prependGitHubDomain(trimGitSuffix(strings.Split(remote, ".com:")[1]))
}

// GetRepository returns the origin repository path, eg: https://github.com/oshalygin/go-changelog
func GetRepository(directory string) (string, error) {

	commandName := "git"
	// args := []string{"config", "--get", "remote.origin.url"}
	args := []string{"log"}

	command := execCommand(commandName, args...)
	command.Dir = directory

	output, err := command.CombinedOutput()

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	remote := string(output[:])
	return remote, nil

}
