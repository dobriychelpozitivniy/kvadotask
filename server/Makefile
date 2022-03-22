go:
	go run ./cmd/main.go

gen_proto:
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. pkg/grpc/proto/kvado.proto