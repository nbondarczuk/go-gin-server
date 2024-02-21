package config

import "fmt"

func ApplicationName() string {
	return string(Options.Viper.Get("application.name").(string))
}

func ServerHTTPAddress() string {
	return string(Options.Viper.Get("server.http.address").(string))
}

func ServerHTTPPort() string {
	return fmt.Sprintf("%d", Options.Viper.Get("server.http.port").(int))
}

func LogLevel() string {
	return string(Options.Viper.Get("log.level").(string))
}

func LogFormat() string {
	return string(Options.Viper.Get("log.format").(string))
}
