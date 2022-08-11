# go-logger

[![GoDoc](https://godoc.org/github.com/masonschafercodes/go-logger?status.svg)](https://godoc.org/github.com/masonschafercodes/go-logger)
[![Go Report Card](https://goreportcard.com/badge/github.com/masonschafercodes/go-logger)](https://goreportcard.com/report/github.com/masonschafercodes/go-logger)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/masonschafercodes/go-logger/blob/master/LICENSE)

A light-weight [Go](http://golang.org) package for printing nicely formatted errors

## Installation

`go get -u github.com/masonschafercodes/go-logger`

## Usage

```go
package main

import (
	"os"

	log "github.com/masonschafercodes/go-logger"
)

func main() {

	dat, err := os.ReadFile("./fake_data.txt")

	if err != nil {
		log.Error(err)
	}

	log.Success(string(dat))

}
```
