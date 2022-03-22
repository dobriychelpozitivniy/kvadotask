package middlewareLogger

import (
	"context"
	"crypto/rand"
	"crypto/sha1"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func NewUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return NewUnaryServerInterceptorWithLogger(&log.Logger)
}

// Logger middleware for grpc. Logging request before handler and logging response after handler.
func NewUnaryServerInterceptorWithLogger(log *zerolog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		data := make([]byte, 10)

		_, _ = rand.Read(data)

		fmt.Println(info.FullMethod)

		id := fmt.Sprintf("%x", sha1.Sum(data))

		log.Info().Interface("Request", req).Str("ID", id).Send()

		resp, err := handler(ctx, req)
		if err != nil {
			log.Error().Interface("Response", resp).Str("ID", id).Send()
		} else {
			log.Info().Interface("Response", resp).Str("ID", id).Send()
		}

		return resp, err
	}
}
