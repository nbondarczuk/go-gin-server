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
			expected: "go-gin-server3",
		},
		{
			label:    "ServerHTTPAddress",
			getter:   ServerHTTPAddress,
			expected: "localhost3",
		},
		{
			label:    "LogLevel",
			getter:   LogLevel,
			expected: "DEBUG3",
		},
		{
			label:    "LogFormat",
			getter:   LogFormat,
			expected: "text3",
		},
		{
			label:    "ServerHTTPPort",
			getter:   ServerHTTPPort,
			expected: "80903",
		},
	}
	input := `application:
  name: go-gin-server3
server:
  http:
    address: localhost3
    port: 80903
log:
  level: DEBUG3
  format: text3
`
	makeTestConfigFile(t, input)
	defer cleanupTestConfigFile()

	err := Init()
	assert.NoError(t, err)

	for _, td := range tests {
		t.Run(td.label, func(t *testing.T) {
			result := td.getter()
			assert.Equal(t, td.expected, result)
		})
	}
}
