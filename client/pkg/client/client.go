package client

import (
	"fmt"
	pb "taskclient/pkg/grpc/proto"
	"time"

	"google.golang.org/grpc"
)

type ClientConfig struct {
	ServerIP   string `mapstructure:"server_ip"`
	ServerPort string `mapstructure:"server_port"`
}

func NewClient(cfg ClientConfig) (pb.KvadoClient, error) {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.WithTimeout(time.Second*5))

	target := fmt.Sprintf("%s%s", cfg.ServerIP, cfg.ServerPort)

	fmt.Println(target)

	fmt.Println("Dial grpc")
	conn, err := grpc.Dial(target, opts...)
	if err != nil {
		return nil, err
	}

	fmt.Println("Grpc is dialed")

	client := pb.NewKvadoClient(conn)

	fmt.Println("Client is init")

	return client, nil
}
