package config

import "fmt"

func ApplicationName() string {
	return string(options.Viper.Get("application.name").(string))
}

func ServerHTTPAddress() string {
	return string(options.Viper.Get("server.http.address").(string))
}

func ServerHTTPPort() string {
	return fmt.Sprintf("%d", options.Viper.GetInt("server.http.port"))
}

func LogLevel() string {
	return string(options.Viper.Get("log.level").(string))
}

func LogFormat() string {
	return string(options.Viper.Get("log.format").(string))
}
