package main

import (
	"go-gin-server/internal/config"
	"go-gin-server/internal/logger"
	"go-gin-server/internal/server"
)

var version string

func main() {
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
