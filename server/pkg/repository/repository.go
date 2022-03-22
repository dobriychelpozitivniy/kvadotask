package repository

import (
	"context"
	"fmt"
	pb "taskserver/pkg/grpc/proto"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type BooksAuthors interface {
	FindAuthorsByBookID(bookID int) ([]*pb.FindAuthorsResponse_Author, error)
	FindBooksByAuthorID(authorID int) ([]*pb.FindBooksResponse_Book, error)
	FindBooksByAuthorName(authorName string) ([]*pb.FindBooksResponse_Book, error)
	FindAuthorByBookName(bookName string) ([]*pb.FindAuthorsResponse_Author, error)
}

type Repository struct {
	BooksAuthors
}

type PgDBConfig struct {
	DBHost     string
	DBPort     string
	DBUsername string
	DBPass     string
	DBName     string
}

func NewRepository(cfg *PgDBConfig) (*Repository, error) {
	conn, err := connectPG(cfg)
	if err != nil {
		return nil, err
	}

	return &Repository{
		BooksAuthors: NewBooksAuthors(conn),
	}, nil
}

func connectPG(cfg *PgDBConfig) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DBUsername, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := pgxpool.Connect(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("Error connect to db: %s", err)
	}

	return conn, nil
}
