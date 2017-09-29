# spwd
[![Build Status](https://travis-ci.org/pinzolo/spwd.png)](http://travis-ci.org/pinzolo/spwd)
[![Go Report Card](https://goreportcard.com/badge/github.com/pinzolo/spwd)](https://goreportcard.com/report/github.com/pinzolo/spwd)
[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/pinzolo/spwd)
[![license](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/pinzolo/spwd/master/LICENSE)

## Description

Secret file based password management tool.

Save your password interactively with `spwd new`.  
Input password is encrypted with AES-256 using your secret file.  
Decrypt and copy password to clipboard with `spwd copy <NAME>`.

## Screenshot

[Screenshot](https://pinzolo.github.io/assets/img/20170928_spwd-sample.gif)

## Configuration

If `~/.config/spwd/config.yml` exists, use it.

```yml
# using secret file path
key_file: /path/to/your/secret/file
# data file path
data_file: /path/to/your/data/file
```

If config file is not found, use below configuration as default.

```yml
key_file: ~/.ssh/id_rsa
data_file: ~/.local/share/spwd/data.dat
```

## Install

```bash
$ go get github.com/pinzolo/spwd
```

## Contribution

1. Fork ([https://github.com/pinzolo/spwd/fork](https://github.com/pinzolo/spwd/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[pinzolo](https://github.com/pinzolo)
