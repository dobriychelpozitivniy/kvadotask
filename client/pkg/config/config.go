package config

import (
	"path"

	"github.com/spf13/viper"
)

type Config struct {
	HostPort   string `mapstructure:"host_port"`
	ServerIP   string `mapstructure:"server_ip"`
	ServerPort string `mapstructure:"server_port"`
}

func Load(cfgPath string) (*Config, error) {
	var config Config

	viper.AddConfigPath(path.Dir(cfgPath))
	viper.SetConfigName(path.Base(cfgPath))

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
