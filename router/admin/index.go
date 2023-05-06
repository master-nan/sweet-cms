package admin

import (
	"github.com/gin-gonic/gin"
	"sweet-cms/admin"
)

func InitIndex(router *gin.RouterGroup) {
	router.GET("", admin.NewIndex().Index)
	router.GET("index", admin.NewIndex().Index)
}
