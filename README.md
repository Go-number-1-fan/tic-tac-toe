# tic-tac-toe [![Build Status](https://travis-ci.org/Go-number-1-fan/tic-tac-toe.svg?branch=master)](https://travis-ci.org/Go-number-1-fan/tic-tac-toe)
The Game of Tic Tac Toe, written in Go.

# Project Description

## Features

* Custom Names with Custom Messages!
* X and O? More Like X and No! Express yourself by picking your own Marker!
* Three Types of Players to play as or against!
* Two Rulesets for maximum variety in gameplay!
* Fully open source code with active and friendly developer -- Feel free to open a pull request!

# Setup

## Notes
* You do not need to `git clone` this repository, `go get {repository}` effectively clones it to your $GOPATH
* some commands use `./...` that is 3 dots, and is a go exclusive way of saying this directory and all subdirectories
* Any issues in project setup can be addressed to Tom via slack | email, he's more than happy to help!

## Requirements

* Git (go uses git to manage external packages)

## Install Go

#### With HomeBrew

`brew install go`

#### Without HomeBrew
[Follow Instructions Here] (https://golang.org/doc/install)

## Set up $GOPATH
* `export GOPATH=$HOME/{YOUR_GOPATH_HERE}`
* EXAMPLE: `export GOPATH=$HOME/GoOut`
* *Note*: Add setting your $GOPATH to your `.bash_profile` or `.zshrc` otherwise it will only work in the current session

## Get the Project
`go get github.com/go-number-1-fan/tic-tac-toe`

## Navigate to the project directory
`cd $GOPATH/src/github.com/go-number-1-fan/tic-tac-toe/`

## Install dependencies
`go get -t -v ./...`

## Run the tests
`go test ./...`

if you get errors with dependencies try `go get github.com/stretchr/testify`

## Build the project
`go build`

## Run the project
`./tic-tac-toe`

`go run main.go`

