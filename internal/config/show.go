package config

import (
	"log/slog"
)

// Show prints all loaded options in the log.
func Show() {
	slog.Info("Config", slog.String("application.name", ApplicationName()))
	slog.Info("Config", slog.String("server.http.address", ServerHTTPAddress()))
	slog.Info("Config", slog.String("server.http.port", ServerHTTPPort()))
	slog.Info("Config", slog.String("log.level", LogLevel()))
	slog.Info("Config", slog.String("log.format", LogFormat()))
}
