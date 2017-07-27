<p align="center">
  <img alt="Golang Logo" src="docs/golang_logo.png" height="140" style="margin-left: 2em;" />
  <h3 align="center">GO Changelog</h3>
  <p align="center">Generate beautiful changelogs based on your git history.</p>
  <p align="center">
    <a href="https://github.com/oshalygin/go-changelog/releases/latest"><img alt="Release" src="https://img.shields.io/github/release/oshalygin/go-changelog.svg?style=flat-square"></a>
    <a href="https://travis-ci.org/oshalygin/go-changelog"><img alt="Travis" src="https://travis-ci.org/oshalygin/go-changelog.svg?branch=master"></a>
    <a href="/LICENSE.md"><img alt="Software License" src="https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square"></a>
    <a href="https://coveralls.io/github/oshalygin/go-changelog?branch=master"><img alt="Coveralls" src="https://coveralls.io/repos/github/oshalygin/go-changelog/badge.svg?branch=master"></a>
    <a href="https://codeclimate.com/repos/597a730a060d05027e000fad/feed"><img alt="Code Climate Issue Count" src="https://codeclimate.com/repos/597a730a060d05027e000fad/badges/d8e88772201d137ea8b7/issue_count.svg"></a>
    <a href="https://goreportcard.com/report/github.com/oshalygin/go-changelog"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/oshalygin/go-changelog"></a>
    <a href="https://godoc.org/github.com/oshalygin/go-changelog"><img src="https://godoc.org/github.com/oshalygin/go-changelog?status.svg" alt="GoDoc"></a>
  </p>
</p>

# Introduction

This is a simple and straightforward CLI utility that will allow you to generate beautiful CHANGELOG's based on your git history.

# Motivation

Face it, you're already spending time writing detailed commit messages, so why not leverage that work in a beautiful way down to your changelog?  Why bother recreating the wheel and literally writing the same thing all over again?

# Requirements

You need to start writing detailed commit messages, if you're not, you're doing a huge disservice to everyone working on the project, including yourself.

This utility is all about convention.  Just run the binary at the root of your repository where your `CHANGELOG.md` lives.

```bash
# This is all you need to do
$ go-changelog
```

# Installation

```bash
go get -u github.com/oshalygin/go-changelog
```

# Usage

```
Usage of git-backup:
  --save boolean
        Generate the new changelog and persist the changes to your current commit.
```

### Example

```bash
# Generate the new CHANGELOG and override the previous one
go-changelog

# This will generate a new changelog based on your current git history
# and subsequently persist the changelog to that same commit.
go-changelog --save

```

### Command Line Arguments

**save**: If you'd like to save the new changelog at the current HEAD commit

# Limitations
The utility depends on writing _CLEAN_ commit messages.  It doesn't read your
GitHub repository pull request history or closed issues.

# License

[MIT](LICENSE)
