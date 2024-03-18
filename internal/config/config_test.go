package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDefaults(t *testing.T) {
	err := InitForTest(nil)
	assert.NoError(t, err)
	assert.Equal(t, DefaultApplicationName, options.GetString("application.name"))
	assert.Equal(t, DefaultServerHTTPAddress, options.GetString("server.http.address"))
	assert.Equal(t, DefaultServerHTTPPort, options.GetInt("server.http.port"))
	assert.Equal(t, DefaultLogLevel, options.GetString("log.level"))
	assert.Equal(t, DefaultLogFormat, options.GetString("log.format"))
}

func TestLoadConfigFromFile(t *testing.T) {
	input := `application:
  name: go-gin-server2
server:
  http:
    address: localhost2
    port: 80902
log:
  level: DEBUG2
  format: text2
`
	makeTestConfigFile(t, input)
	defer cleanupTestConfigFile()

	err := Init()
	assert.NoError(t, err)
	assert.Equal(t, "go-gin-server2", options.GetString("application.name"))
	assert.Equal(t, "localhost2", options.GetString("server.http.address"))
	assert.Equal(t, 80902, options.GetInt("server.http.port"))
	assert.Equal(t, "DEBUG2", options.GetString("log.level"))
	assert.Equal(t, "text2", options.GetString("log.format"))
}
