# Makefile go build

## Purpose

The make files are to be loaded with include GNU make command. They provide
some handy make targets useful in building go programs without any specific IDE.

## Used external tools

- GNU make
- docker
- golint
- golangci-lint
- goimports
- godef
- gosec

### Installation instructions for Go external tools (really good stuff)

- golint

go install golang.org/x/lint/golint@latest

- golangci-lint

Vide: https://golangci-lint.run/usage/install/

Mac OS X

brew install golangci-lint

- goimports

go install golang.org/x/tools/cmd/goimports@latest

- godef

go get -u github.com/rogpeppe/godef

- gosec

go install github.com/securego/gosec/v2/cmd/gosec@latest
