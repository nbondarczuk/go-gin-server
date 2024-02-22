package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGettersWithStringRV(t *testing.T) {
	tests := []struct {
		label    string
		getter   func() string
		expected string
	}{
		{
			label:    "ApplicationName",
			getter:   ApplicationName,
			expected: DefaultApplicationName,
		},
		{
			label:    "ServerHTTPAddress",
			getter:   ServerHTTPAddress,
			expected: DefaultServerHTTPAddress,
		},
		{
			label:    "LogLevel",
			getter:   LogLevel,
			expected: DefaultLogLevel,
		},
		{
			label:    "LogFormat",
			getter:   LogFormat,
			expected: DefaultLogFormat,
		},
	}

	err := InitForTest(nil)
	assert.NoError(t, err)

	for _, td := range tests {
		t.Run(td.label, func(t *testing.T) {
			result := td.getter()
			assert.Equal(t, td.expected, result)
		})
	}
}

func TestGettersWithIntRV(t *testing.T) {
	tests := []struct {
		label    string
		getter   func() int
		expected int
	}{
		{
			label:    "ServerHTTPPort",
			getter:   ServerHTTPPort,
			expected: DefaultServerHTTPPort,
		},
	}

	err := InitForTest(nil)
	assert.NoError(t, err)

	for _, td := range tests {
		t.Run(td.label, func(t *testing.T) {
			result := td.getter()
			assert.Equal(t, td.expected, result)
		})
	}
}
