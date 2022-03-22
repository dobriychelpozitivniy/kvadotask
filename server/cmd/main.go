package main

import (
	"flag"
	"taskserver/pkg/config"
	"taskserver/pkg/grpc/server"
	"taskserver/pkg/repository"

	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	flag.String("config", "configs/local", "display colorized output")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		log.Panic().Msgf("Error parse flag config path: %s", err)
	}

	cfg, err := config.Load(viper.GetString("config"))
	if err != nil {
		log.Panic().Msgf("Error init config: %s", err.Error())
	}

	r, err := repository.NewRepository(&repository.PgDBConfig{
		DBHost:     cfg.DBHost,
		DBPort:     cfg.DBPort,
		DBUsername: cfg.DBUsername,
		DBPass:     cfg.DBPassword,
		DBName:     cfg.DBName,
	})
	if err != nil {
		log.Panic().Msgf("Error init repository: %s", err)
	}

	s := server.NewKvadoServer(r)

	server.StartGRPCServer(s, cfg.Port)
}
