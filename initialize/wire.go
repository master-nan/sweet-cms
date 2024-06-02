//go:build wireinject
// +build wireinject

/**
 * @Author: Nan
 * @Date: 2024/6/1 下午9:54
 */
package initialize

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/gorm"
	"sweet-cms/config"
	"sweet-cms/controller"
	"sweet-cms/repository/impl"
	"sweet-cms/service"
)

type App struct {
	Config *config.Server
	DB     *gorm.DB
	Redis  *redis.Client
	//Router         *gin.Engine
	DictController  *controller.DictController
	BasicController *controller.BasicController
}

var Providers = wire.NewSet(
	LoadConfig,
	InitDB,
	InitRedis,
	//InitRouter,
	impl.NewSysDictRepositoryImpl,
	service.NewSysDictService,
	controller.NewDictController,
	controller.NewBasicController,
	wire.Struct(new(App), "*"),
)

// InitializeApp is just declared here, it will be implemented in wire_gen.go
func InitializeApp() (*App, error) {
	wire.Build(Providers)
	return nil, nil
}
