package initialize

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "sweet-cms/docs"
	"sweet-cms/middleware"
)

func InitRouter(app *App) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	store := cookie.NewStore([]byte(app.Config.Session.Secret))
	router.
		Use(middleware.CorsHandler()).
		Use(middleware.LogHandler(app.LogService)).
		Use(middleware.ResponseHandler()).
		Use(sessions.Sessions("sweet-cms-session", store))
	//总路由
	routerGroup := router.Group("/sweet")
	//后台非验证路由
	adminBaseGroup := routerGroup.Group("/admin")
	{
		adminBaseGroup.POST("/login", app.BasicController.Login)
		adminBaseGroup.GET("/captcha", app.BasicController.Captcha)
		adminBaseGroup.GET("/configure", app.BasicController.Configure)
		adminBaseGroup.POST("/logout", app.BasicController.Logout)

	}

	//后台验证路由
	adminGroup := routerGroup.Group("/admin")
	adminGroup.Use(middleware.AuthHandler(app.JWT, app.UserService))
	{
		// dict
		adminGroup.GET("/dict/id/:id", app.DictController.GetSysDictById)
		adminGroup.GET("/dict/code/:code", app.DictController.GetSysDictByCode)
		adminGroup.GET("/dict/query", app.DictController.QuerySysDict)
		adminGroup.POST("/dict", app.DictController.InsertSysDict)
		adminGroup.PUT("/dict/:id", app.DictController.UpdateSysDict)
		adminGroup.DELETE("/dict/:id", app.DictController.DeleteSysDictById)

		// dict_item
		adminGroup.GET("/dict/items/:id", app.DictController.GetSysDictItemsByDictId)
		adminGroup.GET("/dict/item/:id", app.DictController.GetSysDictItemById)
		adminGroup.POST("/dict/item", app.DictController.InsertSysDictItem)
		adminGroup.PUT("/dict/item/:id", app.DictController.UpdateSysDictItem)
		adminGroup.DELETE("/dict/item/:id", app.DictController.DeleteSysDictItemById)

		// table
		adminGroup.GET("/table/id/:id", app.TableController.GetSysTableByID)
		adminGroup.GET("/table/code/:code", app.TableController.GetSysTableByCode)
		adminGroup.GET("/table/query", app.TableController.QuerySysTable)
		adminGroup.POST("/table", app.TableController.InsertSysTable)
		adminGroup.PUT("/table/:id", app.TableController.UpdateSysTable)
		adminGroup.DELETE("/table/:id", app.TableController.DeleteSysTableById)

		// table_field
		adminGroup.GET("/table/fields/:id", app.TableController.GetSysTableFieldsByTableId)
		adminGroup.GET("/table/field/:id", app.TableController.GetSysTableFieldById)
		adminGroup.POST("/table/field", app.TableController.InsertSysTableField)
		adminGroup.PUT("/table/field/:id", app.TableController.UpdateSysTableField)
		adminGroup.DELETE("/table/field/:id", app.TableController.DeleteSysTableFieldById)

		adminGroup.GET("/generalization/query/:id", app.GeneralizationController.Query)

	}
	return router
}
