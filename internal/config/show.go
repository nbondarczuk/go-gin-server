package config

import (
	"log/slog"
)

// Show prints all loaded options in the log.
func Show() {
	slog.Info("Config options:")
	slog.Info("   application.name", ": ", ApplicationName())
	slog.Info("server.http.address", ": ", ServerHTTPAddress())
	slog.Info("   server.http.port", ": ", ServerHTTPPort())
	slog.Info("          log.level", ": ", LogLevel())
	slog.Info("         log.format", ": ", LogFormat())
}
