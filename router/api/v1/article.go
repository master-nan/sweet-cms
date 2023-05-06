package v1

import (
	"github.com/gin-gonic/gin"
	"sweet-cms/api/v1"
)

func InitArticle(router *gin.RouterGroup) {
	articleRouter := router.Group("article")

	articleRouter.POST("basic", v1.NewArticleBasicApi().Create)
	articleRouter.DELETE("basic/:uuid", v1.NewArticleBasicApi().Delete)
	articleRouter.GET("basic/:uuid", v1.NewArticleBasicApi().Get)
	articleRouter.GET("basic", v1.NewArticleBasicApi().GetList)
	articleRouter.PUT("basic/:uuid", v1.NewArticleBasicApi().Update)

	articleRouter.POST("channel", v1.NewArticleChannelApi().Create)
	articleRouter.DELETE("channel/:id", v1.NewArticleChannelApi().Delete)
	articleRouter.PUT("channel/:id", v1.NewArticleChannelApi().Update)
	articleRouter.GET("channel/:id", v1.NewArticleChannelApi().Get)
	articleRouter.GET("channel", v1.NewArticleChannelApi().GetList)

	articleRouter.POST("content", v1.NewArticleContentApi().Create)
	articleRouter.GET("content/:uuid", v1.NewArticleContentApi().Get)
	articleRouter.DELETE("content/:uuid", v1.NewArticleContentApi().Delete)
	articleRouter.PUT("content/:uuid", v1.NewArticleContentApi().Update)

}
