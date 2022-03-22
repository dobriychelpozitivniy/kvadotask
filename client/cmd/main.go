package main

import (
	"flag"
	"fmt"
	"taskclient/pkg/client"
	"taskclient/pkg/config"
	"taskclient/pkg/handler"

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

	c, err := client.NewClient(client.ClientConfig{
		ServerIP:   cfg.ServerIP,
		ServerPort: cfg.ServerPort,
	})

	if err != nil {
		log.Panic().Msgf("Error init grpc client: %s", err.Error())
	}

	h := handler.NewHandler(c)
	g := h.InitRoutes()

	fmt.Println("Server start ")
	g.Run(cfg.HostPort)
}
