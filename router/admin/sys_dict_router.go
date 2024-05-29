/**
 * @Author: Nan
 * @Date: 2024/5/24 下午5:56
 */

package admin

import (
	"github.com/gin-gonic/gin"
	"sweet-cms/controller"
	"sweet-cms/repository/impl"
	"sweet-cms/service"
)

func InitDictRouter(router *gin.RouterGroup) {
	sysDictRepoImpl := impl.NewSysDictRepositoryImpl()
	sysDictService := service.NewSysDictServer(sysDictRepoImpl)
	dictController := controller.NewDictController(sysDictService)
	router.GET("dict/{id}", dictController.Get)
	router.GET("dict/query", dictController.Query)
	router.POST("dict", dictController.Insert)
	router.PUT("dict", dictController.Update)
	router.DELETE("dict", dictController.Delete)

	//router.GET("dict/items", controller.NewDictController().Get)
	//router.GET("configure", controller.NewBasic().Configure)
	//router.POST("login", controller.NewBasic().Login)
}
