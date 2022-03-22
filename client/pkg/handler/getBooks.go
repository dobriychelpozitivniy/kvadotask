package handler

import (
	"fmt"
	"net/http"
	"time"

	pb "taskclient/pkg/grpc/proto"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

type AuthorIDURI struct {
	AuthorID uint `json:"author_id" uri:"author_id"`
}

type AuthorNameURI struct {
	AuthorName string `json:"author_name" uri:"author_name"`
}

func (h *Handler) GetBooksByAuthorID(c *gin.Context) {
	uri := AuthorIDURI{}

	if err := c.BindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		fmt.Println(err)

		return
	}

	req := &pb.FindBooksByAuthorIDRequest{AuthorId: uint32(uri.AuthorID)}

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1000)
	defer cancel()

	res, err := h.Client.FindBooksByAuthorID(ctx, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		fmt.Println(err)

		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetBooksByAuthorName(c *gin.Context) {
	uri := AuthorNameURI{}

	if err := c.BindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		fmt.Println(err)

		return
	}

	req := &pb.FindBooksByAuthorNameRequest{AuthorName: uri.AuthorName}

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1000)
	defer cancel()

	res, err := h.Client.FindBooksByAuthorName(ctx, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		fmt.Println(err)

		return
	}

	c.JSON(http.StatusOK, res)
}
