package handler

import (
	"fmt"
	"net/http"
	pb "taskclient/pkg/grpc/proto"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

type BookIDURI struct {
	BookID uint `json:"book_id" uri:"book_id"`
}

type BookNameURI struct {
	BookName string `json:"book_name" uri:"book_name"`
}

func (h *Handler) GetAuthorsByBookID(c *gin.Context) {
	uri := BookIDURI{}

	if err := c.BindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		fmt.Println(err)

		return
	}

	req := &pb.FindAuthorsByBookIDRequest{BookId: uint32(uri.BookID)}

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1000)
	defer cancel()

	res, err := h.Client.FindAuthorsByBookID(ctx, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		fmt.Println(err)

		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetAuthorsByBookName(c *gin.Context) {
	uri := BookNameURI{}

	if err := c.BindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		fmt.Println(err)

		return
	}

	req := &pb.FindAuthorsByBookNameRequest{BookName: uri.BookName}

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1000)
	defer cancel()

	res, err := h.Client.FindAuthorsByBookName(ctx, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		fmt.Println(err)

		return
	}

	c.JSON(http.StatusOK, res)
}
