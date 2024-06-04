package initialize

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "sweet-cms/docs"
	"sweet-cms/middlewares"
)

func InitRouter(app *App) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	store := cookie.NewStore([]byte(app.Config.Session.Secret))
	router.
		Use(middlewares.Cors()).
		Use(middlewares.ErrorHandler()).
		Use(middlewares.JSONResponse()).
		Use(middlewares.AccessLog(app.LogService)).
		Use(sessions.Sessions("sweet-cms-session", store))
	//总路由
	routerGroup := router.Group("/sweet")
	//后台非验证路由
	adminBaseGroup := routerGroup.Group("/admin")
	{
		adminBaseGroup.GET("captcha", app.BasicController.Captcha)
		adminBaseGroup.GET("login", app.BasicController.Login)
		adminBaseGroup.GET("configure", app.BasicController.Configure)
		adminBaseGroup.POST("login", app.BasicController.Login)
	}

	//后台验证路由
	adminGroup := routerGroup.Group("/admin")
	//adminGroup.Use(middlewares.Auth(app.JWT))
	{
		adminGroup.GET("dict/id/{id}", app.DictController.GetSysDictById)
		adminGroup.GET("dict/code/{code}", app.DictController.GetSysDictByCode)
		adminGroup.GET("dict/query", app.DictController.QuerySysDict)
		adminGroup.POST("dict", app.DictController.InsertSysDict)
		adminGroup.PUT("dict/{id}", app.DictController.UpdateSysDict)
		adminGroup.DELETE("dict/{id}", app.DictController.DeleteSysDictById)
	}
	//apiBaseV1 := routerGroup.Group("/api/v1")
	//v1.InitBase(apiBaseV1)
	// api接口
	//apiV1 := routerGroup.Group("/api/v1")
	//v1.InitArticle(apiV1)
	//v1.InitQuestionnaire(apiV1)
	return router
}
