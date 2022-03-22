package server

import (
	"context"
	"errors"
	"net"
	pb "taskserver/pkg/grpc/proto"
	"taskserver/pkg/repository"
	mock_repository "taskserver/pkg/repository/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/proto"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {

}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestKvadoServer_FindBooksByAuthorID(t *testing.T) {
	type args struct {
		ctx context.Context
		req *pb.FindBooksByAuthorIDRequest
	}
	tests := []struct {
		name    string
		args    args
		mock    func(mock *mock_repository.MockBooksAuthors)
		want    *pb.FindBooksResponse
		wantErr bool
	}{
		{
			name: "ok",
			mock: func(mock *mock_repository.MockBooksAuthors) {
				res := []*pb.FindBooksResponse_Book{{BookId: 1, Genre: "comedy", Year: 2000, Name: "bookname"}, {BookId: 2, Genre: "comedy", Year: 2006, Name: "bookname"}}

				mock.EXPECT().FindBooksByAuthorID(gomock.Eq(int(2))).Times(1).Return(res, nil)
			},
			args:    args{ctx: context.Background(), req: &pb.FindBooksByAuthorIDRequest{AuthorId: 2}},
			want:    &pb.FindBooksResponse{Books: []*pb.FindBooksResponse_Book{{BookId: 1, Genre: "comedy", Year: 2000, Name: "bookname"}, {BookId: 2, Genre: "comedy", Year: 2006, Name: "bookname"}}},
			wantErr: false,
		},
		{
			name: "author with this id not found",
			mock: func(mock *mock_repository.MockBooksAuthors) {

				mock.EXPECT().FindBooksByAuthorID(gomock.Eq(int(2))).Times(1).Return(nil, errors.New("Error"))
			},
			args:    args{ctx: context.Background(), req: &pb.FindBooksByAuthorIDRequest{AuthorId: 2}},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			BooksAuthorsMock := mock_repository.NewMockBooksAuthors(ctrl)
			tt.mock(BooksAuthorsMock)

			r := repository.Repository{BooksAuthors: BooksAuthorsMock}

			startTestServer(&r)

			ctx := context.Background()
			conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
			if err != nil {
				t.Fatalf("Failed to dial bufnet: %v", err)
			}
			defer conn.Close()
			client := pb.NewKvadoClient(conn)

			got, err := client.FindBooksByAuthorID(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("KvadoServer.FindBooksByAuthorID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !proto.Equal(got.ProtoReflect().Interface(), tt.want.ProtoReflect().Interface()) {
				t.Errorf("KvadoServer.FindBooksByAuthorID() = %v, want %v", got, tt.want)

			}
		})
	}
}

func TestKvadoServer_FindBooksByAuthorName(t *testing.T) {
	type args struct {
		ctx context.Context
		req *pb.FindBooksByAuthorNameRequest
	}
	tests := []struct {
		name    string
		args    args
		mock    func(mock *mock_repository.MockBooksAuthors)
		want    *pb.FindBooksResponse
		wantErr bool
	}{
		{
			name: "ok",
			mock: func(mock *mock_repository.MockBooksAuthors) {
				res := []*pb.FindBooksResponse_Book{{BookId: 1, Genre: "comedy", Year: 2000, Name: "bookname"}, {BookId: 2, Genre: "comedy", Year: 2006, Name: "bookname"}}

				mock.EXPECT().FindBooksByAuthorName(gomock.Eq("author")).Times(1).Return(res, nil)
			},
			args:    args{ctx: context.Background(), req: &pb.FindBooksByAuthorNameRequest{AuthorName: "author"}},
			want:    &pb.FindBooksResponse{Books: []*pb.FindBooksResponse_Book{{BookId: 1, Genre: "comedy", Year: 2000, Name: "bookname"}, {BookId: 2, Genre: "comedy", Year: 2006, Name: "bookname"}}},
			wantErr: false,
		},
		{
			name: "author with this id not found",
			mock: func(mock *mock_repository.MockBooksAuthors) {

				mock.EXPECT().FindBooksByAuthorName(gomock.Eq("author")).Times(1).Return(nil, errors.New("Error"))
			},
			args:    args{ctx: context.Background(), req: &pb.FindBooksByAuthorNameRequest{AuthorName: "author"}},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			BooksAuthorsMock := mock_repository.NewMockBooksAuthors(ctrl)
			tt.mock(BooksAuthorsMock)

			r := repository.Repository{BooksAuthors: BooksAuthorsMock}

			startTestServer(&r)

			ctx := context.Background()
			conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
			if err != nil {
				t.Fatalf("Failed to dial bufnet: %v", err)
			}
			defer conn.Close()
			client := pb.NewKvadoClient(conn)

			got, err := client.FindBooksByAuthorName(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("KvadoServer.FindBooksByAuthorID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !proto.Equal(got.ProtoReflect().Interface(), tt.want.ProtoReflect().Interface()) {
				t.Errorf("KvadoServer.FindBooksByAuthorID() = %v, want %v", got, tt.want)

			}
		})
	}
}

func TestKvadoServer_FindAuthorsByBookID(t *testing.T) {
	type args struct {
		ctx context.Context
		req *pb.FindAuthorsByBookIDRequest
	}
	tests := []struct {
		name    string
		args    args
		mock    func(mock *mock_repository.MockBooksAuthors)
		want    *pb.FindAuthorsResponse
		wantErr bool
	}{
		{
			name: "ok",
			mock: func(mock *mock_repository.MockBooksAuthors) {
				res := []*pb.FindAuthorsResponse_Author{{AuthorId: 1, Name: "John"}, {AuthorId: 2, Name: "Oleg"}}

				mock.EXPECT().FindAuthorsByBookID(gomock.Eq(int(2))).Times(1).Return(res, nil)
			},
			args:    args{ctx: context.Background(), req: &pb.FindAuthorsByBookIDRequest{BookId: 2}},
			want:    &pb.FindAuthorsResponse{Authors: []*pb.FindAuthorsResponse_Author{{AuthorId: 1, Name: "John"}, {AuthorId: 2, Name: "Oleg"}}},
			wantErr: false,
		},
		{
			name: "book with this id not found",
			mock: func(mock *mock_repository.MockBooksAuthors) {

				mock.EXPECT().FindAuthorsByBookID(gomock.Eq(int(2))).Times(1).Return(nil, errors.New("Error"))
			},
			args:    args{ctx: context.Background(), req: &pb.FindAuthorsByBookIDRequest{BookId: 2}},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			BooksAuthorsMock := mock_repository.NewMockBooksAuthors(ctrl)
			tt.mock(BooksAuthorsMock)

			r := repository.Repository{BooksAuthors: BooksAuthorsMock}

			startTestServer(&r)

			ctx := context.Background()
			conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
			if err != nil {
				t.Fatalf("Failed to dial bufnet: %v", err)
			}
			defer conn.Close()
			client := pb.NewKvadoClient(conn)

			got, err := client.FindAuthorsByBookID(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("KvadoServer.FindBooksByAuthorID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !proto.Equal(got.ProtoReflect().Interface(), tt.want.ProtoReflect().Interface()) {
				t.Errorf("KvadoServer.FindBooksByAuthorID() = %v, want %v", got, tt.want)

			}
		})
	}
}

func TestKvadoServer_FindAuthorsByBookName(t *testing.T) {
	type args struct {
		ctx context.Context
		req *pb.FindAuthorsByBookNameRequest
	}
	tests := []struct {
		name    string
		args    args
		mock    func(mock *mock_repository.MockBooksAuthors)
		want    *pb.FindAuthorsResponse
		wantErr bool
	}{
		{
			name: "ok",
			mock: func(mock *mock_repository.MockBooksAuthors) {
				res := []*pb.FindAuthorsResponse_Author{{AuthorId: 1, Name: "John"}, {AuthorId: 2, Name: "Oleg"}}

				mock.EXPECT().FindAuthorByBookName(gomock.Eq("book")).Times(1).Return(res, nil)
			},
			args:    args{ctx: context.Background(), req: &pb.FindAuthorsByBookNameRequest{BookName: "book"}},
			want:    &pb.FindAuthorsResponse{Authors: []*pb.FindAuthorsResponse_Author{{AuthorId: 1, Name: "John"}, {AuthorId: 2, Name: "Oleg"}}},
			wantErr: false,
		},
		{
			name: "book with this name not found",
			mock: func(mock *mock_repository.MockBooksAuthors) {

				mock.EXPECT().FindAuthorByBookName(gomock.Eq("book")).Times(1).Return(nil, errors.New("Error"))
			},
			args:    args{ctx: context.Background(), req: &pb.FindAuthorsByBookNameRequest{BookName: "book"}},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			BooksAuthorsMock := mock_repository.NewMockBooksAuthors(ctrl)
			tt.mock(BooksAuthorsMock)

			r := repository.Repository{BooksAuthors: BooksAuthorsMock}

			startTestServer(&r)

			ctx := context.Background()
			conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
			if err != nil {
				t.Fatalf("Failed to dial bufnet: %v", err)
			}
			defer conn.Close()
			client := pb.NewKvadoClient(conn)

			got, err := client.FindAuthorsByBookName(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("KvadoServer.FindBooksByAuthorID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !proto.Equal(got.ProtoReflect().Interface(), tt.want.ProtoReflect().Interface()) {
				t.Errorf("KvadoServer.FindBooksByAuthorID() = %v, want %v", got, tt.want)

			}
		})
	}
}

func startTestServer(r *repository.Repository) {
	lis = bufconn.Listen(bufSize)
	s := NewKvadoServer(r)
	grpcServer := grpc.NewServer()
	pb.RegisterKvadoServer(grpcServer, s)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal().Msgf("Server exited with error: %v", err)
		}
	}()
}
