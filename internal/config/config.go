package config

import (
	"os"

	"github.com/spf13/viper"
)

var (
	// Options stores all loaded optinos.
	Options *ConfigOptions
)

type (
	// ConfigOptions is a viper embedding.
	ConfigOptions struct {
		*viper.Viper
	}
)

// Init config options from file, flags and env vars.
func Init() (err error) {
	Options, err = NewConfigOptions()
	if err != nil {
		return err
	}

	return nil
}

// SetDefaults sets all default option values.
func (o *ConfigOptions) SetDefaults() {
	// application settings
	o.Viper.SetDefault("application.name", DefaultApplicationName)

	// HTTP service settings
	o.Viper.SetDefault("server.http.address", DefaultServerHTTPAddress)
	o.Viper.SetDefault("server.http.port", DefaultServerHTTPPort)

	// logging settings
	o.Viper.SetDefault("log.level", DefaultLogLevel)
	o.Viper.SetDefault("log.format", DefaultLogFormat)
}

// NewConfigOptions creates all config env.
func NewConfigOptions() (*ConfigOptions, error) {
	// Create Config
	o := &ConfigOptions{
		Viper: viper.New(),
	}

	// Load config from env variables.
	o.Viper.AutomaticEnv()

	// Load hardcoded defaults.
	o.SetDefaults()

	// Load config.yaml file.
	err := o.LoadFromFile()
	if err != nil {
		return nil, err
	}

	return o, nil
}

// LoadFromFile uses config file like $CWD/config/config.yaml.
func (o *ConfigOptions) LoadFromFile() error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	path += "/config"

	v := o.Viper
	v.AddConfigPath(path)
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	err = v.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}

// Show prints all loaded options in the log.
func Show() {

}
