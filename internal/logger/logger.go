package logger

import (
	"log/slog"
	"os"
)

const (
	// LogLevelInfo is the default log level.
	LogLevelInfo = "INFO"
	// LogLevelDebug is used to provide detailed info on oprocessing level.
	LogLevelDebug = "DEBUG"
	// TraceLogLevel is used to provide detailed info on records level.

	// LogFormatJSON is a format where all fields are JSON encoded.
	LogFormatJSON = "json"
	// LogFormatText is a human readable format.
	LogFormatText = "text"
)

var logger *slog.Logger

// Init sets up new logger with a screen output.
func Init(level, format string) error {
	var l slog.Level

	// Log level code check and mapping to slog internal values
	switch level {
	case LogLevelDebug:
		l = slog.LevelDebug
	case LogLevelInfo:
		l = slog.LevelInfo
	default:
		return ErrInvalidLevel
	}

	// Log format code check and creation of specific handler.
	switch format {
	case LogFormatJSON:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	case LogFormatText:
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	default:
		return ErrInvalidFormat
	}

	slog.SetDefault(logger)
	slog.SetLogLoggerLevel(l)

	return nil
}
