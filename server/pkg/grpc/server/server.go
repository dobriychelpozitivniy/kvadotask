package server

import (
	"context"
	"fmt"
	"net"
	middlewareLogger "taskserver/pkg/grpc/logger"
	pb "taskserver/pkg/grpc/proto"
	"taskserver/pkg/repository"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type KvadoServer struct {
	pb.UnimplementedKvadoServer
	r *repository.Repository
}

// Get book's id from request, gets authors from the repository by this id and puts them to response pb.FindAuthorsResponse
func (s *KvadoServer) FindAuthorsByBookID(ctx context.Context, req *pb.FindAuthorsByBookIDRequest) (*pb.FindAuthorsResponse, error) {
	authors, err := s.r.FindAuthorsByBookID(int(req.GetBookId()))
	if err != nil {
		return nil, fmt.Errorf("Error from db: %s", err)
	}

	res := &pb.FindAuthorsResponse{Authors: authors}

	return res, nil
}

// Get author's id from request, gets books from the repository by this id and puts them to response pb.FindBooksReponse
func (s *KvadoServer) FindBooksByAuthorID(ctx context.Context, req *pb.FindBooksByAuthorIDRequest) (*pb.FindBooksResponse, error) {
	books, err := s.r.FindBooksByAuthorID(int(req.GetAuthorId()))
	if err != nil {
		return nil, fmt.Errorf("Error from db: %s", err)
	}

	res := &pb.FindBooksResponse{Books: books}

	return res, nil
}

// Get author's name from request, gets books from the repository by this name and puts them to response pb.FindBooksResponse
func (s *KvadoServer) FindBooksByAuthorName(ctx context.Context, req *pb.FindBooksByAuthorNameRequest) (*pb.FindBooksResponse, error) {
	books, err := s.r.FindBooksByAuthorName(req.GetAuthorName())
	if err != nil {
		return nil, fmt.Errorf("Error from db: %s", err)
	}

	res := &pb.FindBooksResponse{Books: books}

	return res, nil
}

// Get book's name from request, gets authors from the repository by this name and puts them to response pb.FindAuthorsResponse
func (s *KvadoServer) FindAuthorsByBookName(ctx context.Context, req *pb.FindAuthorsByBookNameRequest) (*pb.FindAuthorsResponse, error) {
	authors, err := s.r.FindAuthorByBookName(req.GetBookName())
	if err != nil {
		return nil, fmt.Errorf("Error from db: %s", err)
	}

	res := &pb.FindAuthorsResponse{Authors: authors}

	return res, nil
}

func NewKvadoServer(r *repository.Repository) *KvadoServer {
	return &KvadoServer{r: r}
}

// Init grpc server and start him
func StartGRPCServer(server *KvadoServer, port string) {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Panic().Msgf("Error in listen %s", err.Error())
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			middlewareLogger.NewUnaryServerInterceptor(),
		)),
	)

	pb.RegisterKvadoServer(grpcServer, server)

	fmt.Println("Start server")

	if err := grpcServer.Serve(listen); err != nil {
		log.Panic().Msgf("error serve grpc: %s", err.Error())
	}
}
