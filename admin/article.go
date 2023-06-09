package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sweet-cms/utils"
)

type Article struct {
}

func NewArticle() *Article {
	return &Article{}
}

func (a Article) GetList(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "article_list.html", gin.H{
		"menu": utils.GetMenu(),
		"url":  ctx.Request.URL.String(),
	})
}

func (a Article) GetChannelList(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "article_channel_list.html", gin.H{
		"menu": utils.GetMenu(),
		"url":  ctx.Request.URL.String(),
	})
}
