# spechar

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gowizzard/spechar.svg)](https://golang.org/) [![Go](https://github.com/gowizzard/spechar/actions/workflows/go.yml/badge.svg)](https://github.com/gowizzard/spechar/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/gowizzard/spechar)](https://goreportcard.com/report/github.com/gowizzard/spechar) [![Go Reference](https://pkg.go.dev/badge/github.com/gowizzard/spechar.svg)](https://pkg.go.dev/github.com/gowizzard/spechar)

Is a small library for removing special characters from strings.

## Install

First you have to install the package:

```console
go get github.com/gowizzard/spechar
```

## How to use

Here is a small example how to convert the string:

```go
convert := spechar.Convert("{Mein Test&!}")
fmt.Println(convert)
```
