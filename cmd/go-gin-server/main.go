package main

import (
	"log/slog"
	"os"

	"go-gin-server/internal/config"
	"go-gin-server/internal/log"
	"go-gin-server/internal/server"
)

var version string

func showStartupInfo() {
	log.Logger.Info("Config", slog.String("application.name", config.ApplicationName()))
	log.Logger.Info("Config", slog.String("server.http.address", config.ServerHTTPAddress()))
	log.Logger.Info("Config", slog.String("server.http.port", config.ServerHTTPPort()))
	log.Logger.Info("Config", slog.String("log.level", config.LogLevel()))
	log.Logger.Info("Config", slog.String("log.format", config.LogFormat()))
}

func main() {
	if err := config.Init(); err != nil {
		slog.Error("Invalid config, exiting",
			slog.String("err",
				err.Error()))
		os.Exit(1)
	}

	// Start logger now as it may require to change log format.
	if err := log.Init(version, config.LogLevel(), config.LogFormat()); err != nil {
		slog.Error("Error initializing logger, exiting",
			slog.String("err",
				err.Error()))
		os.Exit(2)
	}

	showStartupInfo()

	// Start web service API.
	server, err := server.New(version)
	if err != nil {
		slog.Error("Error creating server, exiting",
			slog.String("err",
				err.Error()))
		os.Exit(3)
	}

	server.Run()
}
