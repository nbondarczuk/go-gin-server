package config

import (
	"github.com/spf13/viper"
)

var (
	options        *Options
	ConfigPath     = DefaultConfigPathName
	ConfigFileName = DefaultConfigFileName
)

// Options is a viper embedding.
type Options struct {
	*viper.Viper
}

// Init loads configuration first using defaults, then from a config file.
func Init() error {
	// load config from env variables
	options = &Options{viper.New()}

	// Set defaults for all env application settings
	initDefaults()

	// Bind viper names with env variables.
	bindEnvVars()

	// Use config file to override dafaults.
	if err := loadConfigFromFile(); err != nil {
		return err
	}

	return nil
}

func initDefaults() {
	// application settings
	options.Viper.SetDefault("application.name", DefaultApplicationName)

	// server settings
	options.Viper.SetDefault("server.http.address", DefaultServerHTTPAddress)
	options.Viper.SetDefault("server.http.port", DefaultServerHTTPPort)

	// logging settings
	options.Viper.SetDefault("log.level", DefaultLogLevel)
	options.Viper.SetDefault("log.format", DefaultLogFormat)
}

func bindEnvVars() {
	options.Viper.BindEnv("application.name", "APPLICATION_NAME")
	options.Viper.BindEnv("server.http.address", "SERVER_HTTP_ADDRESS")
	options.Viper.BindEnv("server.http.port", "SERVER_HTTP_PORT")
	options.Viper.BindEnv("log.level", "LOG_LEVEL")
	options.Viper.BindEnv("log.format", "LOG_FORMAT")
}

func loadConfigFromFile() error {
	options.Viper.AddConfigPath(ConfigPath)
	options.Viper.SetConfigName(ConfigFileName)
	options.Viper.SetConfigType("yaml")
	if err := options.Viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
