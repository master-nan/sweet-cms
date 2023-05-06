package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sweet-cms/utils"
)

type Index struct {
}

func NewIndex() *Index {
	return &Index{}
}

func (i Index) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"menu": utils.GetMenu(),
		"url":  ctx.Request.URL.String(),
	})
}

// get user
