package admin

import (
	"github.com/gin-gonic/gin"
	"sweet-cms/controller"
)

func InitBasicRouter(router *gin.RouterGroup) {
	router.GET("captcha", controller.NewBasic().Captcha)
	router.GET("login", controller.NewBasic().Login)
	router.GET("configure", controller.NewBasic().Configure)
	router.POST("login", controller.NewBasic().Login)
}
