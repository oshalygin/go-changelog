package main

import (
	"testing"

	. "github.com/franela/goblin"
)

func Test_Utilities(t *testing.T) {
	g := Goblin(t)
	g.Describe("Utilities", func() {
		g.It("should return the string `https://github.com/oshalygin/go-changelog` when passing `https://github.com/oshalygin/go-changelog.git` to trimGitSuffix", func() {

			expected := "https://github.com/oshalygin/go-changelog"
			remote := "https://github.com/oshalygin/go-changelog.git"

			actual := trimGitSuffix(remote)
			g.Assert(actual).Equal(expected)
		})

		g.It("should return the string `oshalygin/go-changelog` when passing `oshalygin/go-changelog.git` to trimGitSuffix", func() {

			expected := "oshalygin/go-changelog"
			remote := "oshalygin/go-changelog.git"

			actual := trimGitSuffix(remote)
			g.Assert(actual).Equal(expected)
		})

		g.It("should prepend `https://github.com/` to `oshalygin/go-changelog` and return `https://github.com/oshalygin/go-changelog` when calling prependGitHubDomain", func() {

			expected := "https://github.com/oshalygin/go-changelog"
			remote := "oshalygin/go-changelog"

			actual := prependGitHubDomain(remote)
			g.Assert(actual).Equal(expected)
		})

		g.It("should return the string `https://github.com/oshalygin/go-changelog` when passing `https://github.com/oshalygin/go-changelog.git` to getRepositoryURL", func() {

			expected := "https://github.com/oshalygin/go-changelog"
			remote := "https://github.com/oshalygin/go-changelog.git"

			actual := getRepositoryURL(remote)
			g.Assert(actual).Equal(expected)
		})

		g.It("should return the string `https://github.com/oshalygin/go-changelog` when passing `git@github.com:oshalygin/go-changelog.git` to getRepositoryURL", func() {

			expected := "https://github.com/oshalygin/go-changelog"
			remote := "git@github.com:oshalygin/go-changelog.git"

			actual := getRepositoryURL(remote)
			g.Assert(actual).Equal(expected)
		})

		g.It("should return an empty string when calling getRepositoryURL if the remote does not include http or git", func() {

			expected := ""
			remote := "foobar"

			actual := getRepositoryURL(remote)
			g.Assert(actual).Equal(expected)
		})
	})
}
