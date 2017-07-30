package main

import (
	"os"
	"os/exec"
	"testing"

	. "github.com/franela/goblin"
)

// Command spys
var passedInCommand string
var passedInArgs []string

// From Go's source code
// https://golang.org/src/os/exec/exec_test.go

func fakeExecCommand(command string, args ...string) *exec.Cmd {

	passedInCommand = command
	passedInArgs = args

	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}

	return cmd
}

func Test_GitService(t *testing.T) {
	g := Goblin(t)
	g.Describe("Git Service", func() {

		g.It("should fail when calling GitLog on a directory that does not exist", func() {

			execCommand = fakeExecCommand
			defer func() {
				execCommand = exec.Command
			}()

			expected := ""
			directory := "./foobar"
			formatter := "foobarFormatter"

			actual, _ := GitLog(directory, formatter)
			g.Assert(actual).Equal(expected)
		})

		g.It("should return an error when calling GitLog on a directory that does not exist", func() {

			execCommand = fakeExecCommand
			defer func() {
				execCommand = exec.Command
			}()

			expected := "chdir ./foobar: no such file or directory"
			directory := "./foobar"
			formatter := "foobarFormatter"

			_, err := GitLog(directory, formatter)
			actual := err.Error()

			g.Assert(actual).Equal(expected)
		})

		g.It("should call the git command when calling GitLog", func() {

			expected := "git"
			execCommand = fakeExecCommand
			defer func() {
				execCommand = exec.Command
			}()

			formatter := "commit:       %H%nAuthor:       %ae%nAuthor Date:  %cd%n%n"
			directory := "."
			GitLog(directory, formatter)

			actual := passedInCommand
			g.Assert(actual).Equal(expected)
		})

		g.It("should call the GitLog with the args 'log' and the customFormatter", func() {

			formatter := "commit:       %H%nAuthor:       %ae%nAuthor Date:  %cd%n%n"
			customFormatter := "--pretty=format:" + formatter

			expected := []string{"log", customFormatter}
			execCommand = fakeExecCommand
			defer func() {
				execCommand = exec.Command
			}()

			directory := "."
			GitLog(directory, formatter)

			actual := passedInArgs
			g.Assert(actual).Equal(expected)
		})

		g.It("should call the git command when calling GetRepository", func() {

			expected := "git"
			execCommand = fakeExecCommand
			defer func() {
				execCommand = exec.Command
			}()

			directory := "."
			GetRepository(directory)

			actual := passedInCommand
			g.Assert(actual).Equal(expected)
		})

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

	})
}
