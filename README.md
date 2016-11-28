# tic-tac-toe [![Build Status](https://travis-ci.org/Go-number-1-fan/tic-tac-toe.svg?branch=master)](https://travis-ci.org/Go-number-1-fan/tic-tac-toe)
The Game of Tic Tac Toe, written in Go.

# Setup

## Requirements

* Git (go uses git to manage external packages)

## Install Go

#### With HomeBrew

`brew install go`

#### Without HomeBrew
[Follow Instructions Here] (https://golang.org/doc/install)

## Set up $GOPATH
* GoPath is a directory that contains a `bin/` `src/` `pkg/` folder, all projects you get or dependencies a project has are installed here. Simply create a directory and set it to whatever you want.
* EXAMPLE: `export GOPATH=$HOME/GoOut`
* `export GOPATH=$HOME/{YOUR_GOPATH_HERE}`
* Add the above line to your `.bash_profile` or `.zhsrc`

## Get the Project
`go get github.com/go-number-1-fan/tic-tac-toe`

## Navigate to the project directory
`cd $GOPATH/src/github.com/go-number-1-fan/tic-tac-toe/`

## Build the project
` go build`

## Run the project
`./tic-tac-toe`

