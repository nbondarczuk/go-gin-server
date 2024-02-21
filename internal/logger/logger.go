package logger

import (
	"github.com/sirupsen/logrus"
)

const (
	// LogLevelNormal is the default loge level.
	LogLevelNormal = "NORMAL"
	// LogLevelDebug is used to provide detailed info on oprocessing level.
	LogLevelDebug  = "DEBUG"
	// TraceLogLevel is used to provide detailed info on records level.

	// LogFormatJSON is a format where all fields are JSON encoded.
	LogFormatJSON = "json"
	// LogFormatText is a human readable format.
	LogFormatText = "text"
)

// Log is a global object for accessing log stream with minimal overhead.
var Log = logrus.New()

// Init sets up new logger with normal screen output.
func Init(level, format string) error {
	switch format {
	case LogFormatJSON:
		Log.SetFormatter(&logrus.JSONFormatter{})
	case LogFormatText:
		Log.SetFormatter(&logrus.TextFormatter{})
	default:
		return ErrInvalidFormat
	}

	return nil
}
