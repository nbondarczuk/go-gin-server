package config

import (
	"io"
	"os"

	"github.com/spf13/viper"
)

var (
	options        *Options
	ConfigPath     string
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
	options.Viper.AutomaticEnv()

	// Set defaults for all env application settings
	initDefaults()

	// Use config file to override dafaults.
	if err := loadConfigFromFile(); err != nil {
		return err
	}

	return nil
}

func initDefaults() {
	options.Viper.SetDefault("application.name", DefaultApplicationName)
	options.Viper.SetDefault("server.http.address", DefaultServerHTTPAddress)
	options.Viper.SetDefault("server.http.port", DefaultServerHTTPPort)

	// logging settings
	options.Viper.SetDefault("log.level", DefaultLogLevel)
	options.Viper.SetDefault("log.format", DefaultLogFormat)
}

func loadConfigFromFile() error {
	v := options.Viper

	// Testing may override this path.
	if ConfigPath == "" {
		ConfigPath, err := os.Getwd()
		if err != nil {
			return errNoWorkingDir
		}
		ConfigPath += "/config"
	}

	v.AddConfigPath(ConfigPath)
	v.SetConfigName(ConfigFileName)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

// InitForTest does the same as Init but it does not use config file.
func InitForTest(input io.Reader) error {
	options = &Options{viper.New()}
	options.Viper.AutomaticEnv()

	// Set defaults for all env application settings
	initDefaults()

	// Read config from provided reader.
	if input != nil {
		if err := viper.ReadConfig(input); err != nil {
			return err
		}
	}

	return nil
}
