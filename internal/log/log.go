package log

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

var Logger *slog.Logger

// Init sets up new logger with a screen output.
func Init(version, level, format string) error {
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
		Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	case LogFormatText:
		Logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	default:
		return ErrInvalidFormat
	}

	Logger = Logger.With(
		slog.Group("process",
			slog.Int("pid", os.Getpid()),
			slog.String("version", version),
		),
	)

	slog.SetDefault(Logger)
	slog.SetLogLoggerLevel(l)

	return nil
}
