package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	input := `application:
  name: go-gin-server2
server:
  http:
    address: localhost2
    port: 80902
log:
  level: DEBUG2
  format: text2
repsitory:
  dbname: mongo
  url: mongodb://localhost:27017
cache:
  redis:
    address: 127.0.0.1:6379
    password: xyz
    db: 0
`
	makeTestConfigFile(t, input)
	defer cleanupTestConfigFile(t)

	err := Init()
	assert.NoError(t, err)
	assert.Equal(t, "go-gin-server2", options.GetString("application.name"))
	assert.Equal(t, "localhost2", options.GetString("server.http.address"))
	assert.Equal(t, 80902, options.GetInt("server.http.port"))
	assert.Equal(t, "DEBUG2", options.GetString("log.level"))
	assert.Equal(t, "text2", options.GetString("log.format"))
	assert.Equal(t, "mongo", options.GetString("repository.dbname"))
	assert.Equal(t, "mongodb://localhost:27017", options.GetString("repository.url"))
	assert.Equal(t, "127.0.0.1:6379", options.GetString("cache.redis.address"))
	assert.Equal(t, "xyz", options.GetString("cache.redis.password"))
	assert.Equal(t, "0", options.GetString("cache.redis.db"))
}
