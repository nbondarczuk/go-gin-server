package logger

import (
	"errors"
)

var (
	// ErrInvalidFormat signifies invalid format text is ised in config.
	// Allowed values are: json, text.
	ErrInvalidFormat = errors.New("provided logging format is invalid")
)
