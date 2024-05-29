package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Index struct {
}

func NewIndex() *Index {
	return &Index{}
}

func (i Index) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"url": ctx.Request.URL.String(),
	})
}

// get user
