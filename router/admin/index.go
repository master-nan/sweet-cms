package admin

import (
	"github.com/gin-gonic/gin"
	"sweet-cms/controller"
)

func InitIndex(router *gin.RouterGroup) {
	router.GET("", controller.NewIndex().Index)
	router.GET("index", controller.NewIndex().Index)
}
