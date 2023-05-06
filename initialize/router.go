package initialize

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"sweet-cms/global"
	"sweet-cms/middlewares"
	"sweet-cms/router/admin"
	"sweet-cms/router/api/v1"
	"time"
)

func Routers() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:    []string{"*"},
		ExposeHeaders:   []string{"Authorization"},
		MaxAge:          12 * time.Hour,
	})).Use(middlewares.AccessLog())
	store := cookie.NewStore([]byte(global.ServerConf.Session.Secret))
	router.Use(sessions.Sessions("sweet-cms-session", store))
	router.Static("/assets", "./static")
	// 总路由
	routerGroup := router.Group("/sweet")

	// 后台非验证路由
	adminBaseGroup := routerGroup.Group("/admin")
	admin.InitBasic(adminBaseGroup)

	// 后台验证路由
	adminGroup := routerGroup.Group("/admin")
	admin.InitIndex(adminGroup)
	admin.InitArticle(adminGroup)

	apiBaseV1 := routerGroup.Group("/api/v1")
	v1.InitBase(apiBaseV1)
	// api接口
	apiV1 := routerGroup.Group("/api/v1")
	apiV1.Use(middlewares.Auth())
	v1.InitArticle(apiV1)
	v1.InitQuestionnaire(apiV1)

	return router
}
