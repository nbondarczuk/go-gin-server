package config

const (
	// Hardcoded initial file name, mauy be changed in test
	DefaultConfigFileName = "config.yaml"
	DefaultConfigPathName = "config"

	// Default option values
	DefaultApplicationName   = "go-gin-server"
	DefaultServerHTTPAddress = "localhost"
	DefaultServerHTTPPort    = "8000"
	DefaultLogLevel          = "INFO"
	DefaultLogFormat         = "json"

	// Log levels
	LogLevelInfo  = "INFO"
	LogLevelDebug = "DEBUG"

	// backend
	DefaultBackendDBName = "mongo"
	DefaultBackendURL    = "mongodb://localhost:27017/?connect=direct"
)
