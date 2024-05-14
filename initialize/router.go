package initialize

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"sweet-cms/global"
	"sweet-cms/middlewares"
	"sweet-cms/router/admin"
	"sweet-cms/router/api/v1"
)

func Routers() *gin.Engine {
	router := gin.Default()
	store := cookie.NewStore([]byte(global.ServerConf.Session.Secret))
	router.
		Use(middlewares.Cors()).
		Use(middlewares.AccessLog()).
		Use(sessions.Sessions("sweet-cms-session", store))
	// 总路由
	routerGroup := router.Group("/sweet")
	// 后台非验证路由
	adminBaseGroup := routerGroup.Group("/admin")
	admin.InitBasic(adminBaseGroup)

	// 后台验证路由
	adminGroup := routerGroup.Group("/admin")
	adminGroup.Use(middlewares.Auth())
	admin.InitIndex(adminGroup)
	admin.InitArticle(adminGroup)

	apiBaseV1 := routerGroup.Group("/api/v1")
	v1.InitBase(apiBaseV1)
	// api接口
	//apiV1 := routerGroup.Group("/api/v1")
	//v1.InitArticle(apiV1)
	//v1.InitQuestionnaire(apiV1)

	return router
}
