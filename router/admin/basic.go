package admin

import (
	"github.com/gin-gonic/gin"
	"sweet-cms/admin"
)

func InitBasic(router *gin.RouterGroup) {
	router.GET("captcha", admin.NewBasic().Captcha)
	router.GET("login", admin.NewBasic().Login)
	router.POST("login", admin.NewBasic().Login)
}
