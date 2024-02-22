package config

import (
	"go-gin-server/internal/logger"
)

// Show prints all loaded options in the log.
func Show() {
	logger.Log.Info("Config options:")
	logger.Log.Info("   application.name", ": ", ApplicationName())
	logger.Log.Info("server.http.address", ": ", ServerHTTPAddress())
	logger.Log.Info("   server.http.port", ": ", ServerHTTPPort())
	logger.Log.Info("          log.level", ": ", LogLevel())
	logger.Log.Info("         log.format", ": ", LogFormat())
}
