package config

import (
	"io/ioutil"
	"os"
	"testing"
)

func makeTestConfigFile(t *testing.T, input string) {
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
	ConfigPath = dir
	t.Logf("Config dir is: %s", dir)
	f, err := ioutil.TempFile(dir, "config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	s, err := f.Stat()
	if err != nil {
		t.Fatal(err)
	}
	ConfigFileName = s.Name()
	t.Logf("Config file name is: %s", ConfigFileName)

	n, err := f.Write([]byte(input))
	if err != nil || n != len(input) {
		t.Fatal(err)
	}

	t.Logf("Produced test file: %s/%s, size: %d", ConfigPath, ConfigFileName, n)
}

func cleanupTestConfigFile() {
	os.RemoveAll(ConfigPath)
}
