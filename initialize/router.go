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
		adminBaseGroup.GET("/test", app.BasicController.Test)

	}

	//后台验证路由
	adminGroup := routerGroup.Group("/admin")
	adminGroup.Use(middleware.AuthHandler(app.JWT, app.UserService))
	{
		// dict
		adminGroup.GET("/dict/id/:id", app.DictController.GetSysDictById)
		adminGroup.GET("/dict/code/:code", app.DictController.GetSysDictByCode)
		adminGroup.GET("/dict/query", app.DictController.QuerySysDict)
		adminGroup.POST("/dict", app.DictController.CreateSysDict)
		adminGroup.PUT("/dict/:id", app.DictController.UpdateSysDict)
		adminGroup.DELETE("/dict/:id", app.DictController.DeleteSysDictById)

		// dict_item
		adminGroup.GET("/dict/items/:id", app.DictController.GetSysDictItemsByDictId)
		adminGroup.GET("/dict/item/:id", app.DictController.GetSysDictItemById)
		adminGroup.POST("/dict/item", app.DictController.CreateSysDictItem)
		adminGroup.PUT("/dict/item/:id", app.DictController.UpdateSysDictItem)
		adminGroup.DELETE("/dict/item/:id", app.DictController.DeleteSysDictItemById)

		// table
		adminGroup.GET("/table/id/:id", app.TableController.GetTableByID)
		adminGroup.GET("/table/code/:code", app.TableController.GetTableByCode)
		adminGroup.GET("/table/query", app.TableController.QueryTable)
		adminGroup.POST("/table", app.TableController.CreateTable)
		adminGroup.PUT("/table/:id", app.TableController.UpdateTable)
		adminGroup.DELETE("/table/:id", app.TableController.DeleteTableById)

		// table_field
		adminGroup.GET("/table/fields/:id", app.TableController.GetTableFieldsByTableId)
		adminGroup.GET("/table/field/:id", app.TableController.GetTableFieldById)
		adminGroup.POST("/table/field", app.TableController.CreateTableField)
		adminGroup.PUT("/table/field/:id", app.TableController.UpdateTableField)
		adminGroup.DELETE("/table/field/:id", app.TableController.DeleteTableFieldById)

		adminGroup.GET("/table/init/:code", app.TableController.InitTable)

		// table_index
		adminGroup.GET("/table/indexes/:id", app.TableController.GetTableIndexesByTableId)
		adminGroup.GET("/table/index/:id", app.TableController.GetTableIndexById)
		adminGroup.POST("/table/index", app.TableController.CreateTableIndex)
		adminGroup.PUT("/table/index/:id", app.TableController.UpdateTableIndex)
		adminGroup.DELETE("/table/index/:id", app.TableController.DeleteTableIndexById)

		// table_relation
		adminGroup.GET("/table/relations/:id", app.TableController.GetTableRelationsByTableId)
		adminGroup.GET("/table/relation/:id", app.TableController.GetTableRelationById)
		adminGroup.POST("/table/relation", app.TableController.CreateTableRelation)
		adminGroup.PUT("/table/relation/:id", app.TableController.UpdateTableRelation)
		adminGroup.DELETE("/table/relation/:id", app.TableController.DeleteTableRelationById)

		// menu
		adminGroup.GET("/menu/id/:id", app.MenuController.GetSysMenuById)
		adminGroup.GET("/menu/query", app.MenuController.QuerySysMenu)
		adminGroup.POST("/menu", app.MenuController.CreateSysMenu)
		adminGroup.PUT("/menu/:id", app.MenuController.UpdateSysMenu)
		adminGroup.DELETE("/menu/:id", app.MenuController.DeleteSysMenuById)
		adminGroup.GET("/menu/my", app.MenuController.GetMyMenus)
		adminGroup.GET("/menu/role", app.MenuController.GetRoleMenus)

		// user
		adminGroup.GET("/user/me", app.UserController.GetMe)
		adminGroup.GET("/user/query", app.UserController.QuerySysUser)
		adminGroup.POST("/user", app.UserController.CreateSysUser)
		adminGroup.PUT("/user/:id", app.UserController.UpdateSysUser)
		adminGroup.DELETE("/user/:id", app.UserController.DeleteSysUser)

		// generalization
		adminGroup.GET("/generalization/query/:id", app.GeneralizationController.Query)

	}
	return router
}
