package config

const (
	// Hardcoded initial file name, mauy be changed in test
	DefaultConfigFileName = "config.yaml"

	// Default option values
	DefaultApplicationName   = "go-gin-server"
	DefaultServerHTTPAddress = "localhost"
	DefaultServerHTTPPort    = 8080
	DefaultLogLevel          = "DEBUG"
	DefaultLogFormat         = "text"

	// Log levels
	LogLevelNormal = "NORMAL"
	LogLevelDebug  = "DEBUG"
)
