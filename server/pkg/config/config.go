package config

import (
	"path"

	"github.com/spf13/viper"
)

type Config struct {
	Port       string `mapstructure:"port"`
	PgDBConfig `mapstructure:"pgdb_config"`
}

type PgDBConfig struct {
	DBHost     string `mapstructure:"db_host"`
	DBUsername string `mapstructure:"db_username"`
	DBPassword string `mapstructure:"db_password"`
	DBPort     string `mapstructure:"db_port"`
	DBName     string `mapstructure:"db_name"`
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
