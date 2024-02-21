package main

import (
	"go-gin-server/internal/server"
	"go-gin-server/internal/config"
	"go-gin-server/internal/logger"
)

var version string

func main() {
	// Load config (with no logger created yet).
	if err := config.Init(); err != nil {
		panic(err)
	}

	// Start logger now as it may require to change log format.
	if err := logger.Init(config.LogLevel(), config.LogFormat()); err != nil {
		panic(err)
	}

	config.Show()

	// Start web service API.
	server, err := server.New()
	if err != nil {
		panic(err)
	}
	server.Run()
}
