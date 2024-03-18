package main

import (
	"log/slog"
	"os"

	"go-gin-server/internal/config"
	"go-gin-server/internal/logger"
	"go-gin-server/internal/server"
)

var version string

func main() {
	if err := config.Init(); err != nil {
		slog.Error("Invalid config, exiting")
		os.Exit(1)
	}

	// Start logger now as it may require to change log format.
	if err := logger.Init(config.LogLevel(), config.LogFormat()); err != nil {
		slog.Error("Error initializing logger, exiting")
		os.Exit(2)
	}

	config.Show()

	// Start web service API.
	server, err := server.New(version)
	if err != nil {
		slog.Error("Error creating server, exiting")
		os.Exit(3)
	}

	server.Run()
}
