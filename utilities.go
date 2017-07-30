package main

import "strings"

func prependGitHubDomain(path string) string {
	githubDomain := "https://github.com/"
	return githubDomain + path
}

func trimGitSuffix(path string) string {
	return strings.Split(path, ".git")[0]
}

func getRepositoryURL(remote string) string {
	if strings.HasPrefix(remote, "https") {
		return trimGitSuffix(remote)
	}
	if strings.HasPrefix(remote, "git") {
		return prependGitHubDomain(trimGitSuffix(strings.Split(remote, ".com:")[1]))
	}
	return ""
}
