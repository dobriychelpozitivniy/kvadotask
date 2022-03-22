package handler

import "github.com/gin-gonic/gin"
import pb "taskclient/pkg/grpc/proto"

type Handler struct {
	Client pb.KvadoClient
}

func NewHandler(client pb.KvadoClient) *Handler {
	return &Handler{Client: client}
}

func (h *Handler) InitRoutes() *gin.Engine {
	g := gin.New()

	g.GET("/books/by_author_id/:author_id", h.GetBooksByAuthorID)
	g.GET("/authors/by_book_id/:book_id", h.GetAuthorsByBookID)
	g.GET("/authors/by_book_name/:book_name", h.GetAuthorsByBookName)
	g.GET("/books/by_author_name/:author_name", h.GetBooksByAuthorName)

	return g
}
