package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	Port string
}

func NewConfiguration() (*Config, error) {

	pflag.String("port", "8080", "port for server")

	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		return nil, err
	}
	viper.AutomaticEnv()

	viper.AutomaticEnv()

	return &Config{
		Port: viper.GetString("port"),
	}, nil
}
