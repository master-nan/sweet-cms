package admin

import (
	"github.com/gin-gonic/gin"
	"sweet-cms/admin"
)

func InitArticle(router *gin.RouterGroup) {
	router.GET("article/index", admin.NewArticle().GetList)
	router.GET("article/channel", admin.NewArticle().GetChannelList)
}
