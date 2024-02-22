package config

import (
	"io/ioutil"
	"os"
	"strings"
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

func xTestInitForTest(t *testing.T) {
	input := `application:
  name: go-gin-server1
server:
  http:
    address: localhost1
    port: 80901
log:
  level: DEBUG1
  format: text1
`
	err := InitForTest(strings.NewReader(input))
	assert.NoError(t, err)
	assert.Equal(t, "go-gin-server1", options.GetString("application.name"))
	assert.Equal(t, "localhost1", options.GetString("server.http.address"))
	assert.Equal(t, 80901, options.GetInt("server.http.port"))
	assert.Equal(t, "DEBUG1", options.GetString("log.level"))
	assert.Equal(t, "text1", options.GetString("log.format"))
}

func TestLoadConfigFromFile(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	dir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatal(err)
	}
	//defer os.RemoveAll(dir)
	err = os.Chdir(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(pwd)
	configPath = dir
	t.Logf("Config dir is: %s", dir)
	f, err := ioutil.TempFile(dir, "config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	s, err := f.Stat()
	if err != nil {
		t.Fatal(err)
	}
	configFileName = s.Name()
	t.Logf("Config file name is: %s", configFileName)
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
	if n, err := f.Write([]byte(input)); err != nil || n != len(input) {
		t.Fatal(err)
	}
	err = Init()
	assert.NoError(t, err)
	assert.Equal(t, "go-gin-server2", options.GetString("application.name"))
	assert.Equal(t, "localhost2", options.GetString("server.http.address"))
	assert.Equal(t, 80902, options.GetInt("server.http.port"))
	assert.Equal(t, "DEBUG2", options.GetString("log.level"))
	assert.Equal(t, "text2", options.GetString("log.format"))
}
