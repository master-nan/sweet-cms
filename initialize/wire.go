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
	"sweet-cms/cache"
	"sweet-cms/config"
	"sweet-cms/controller"
	"sweet-cms/inter"
	"sweet-cms/repository"
	"sweet-cms/repository/impl"
	"sweet-cms/service"
	"sweet-cms/utils"
)

type App struct {
	Config          *config.Server
	DB              *gorm.DB
	Redis           *redis.Client
	SF              *utils.Snowflake
	JWT             *utils.JWTTokenGen
	DictController  *controller.DictController
	BasicController *controller.BasicController
	LogService      *service.LogService
}

var Providers = wire.NewSet(
	LoadConfig,
	InitDB,
	InitRedis,
	InitSnowflake,
	InitValidators,
	utils.NewJWTTokenGen,
	utils.NewRedisUtil,
	impl.NewSysConfigureRepositoryImpl,
	impl.NewSysDictRepositoryImpl,
	impl.NewLogRepositoryImpl,
	impl.NewSysUserRepositoryImpl,

	wire.Bind(new(inter.CacheInterface), new(*utils.RedisUtil)),
	wire.Bind(new(inter.TokenGenerator), new(*utils.JWTTokenGen)),
	wire.Bind(new(repository.LogRepository), new(*impl.LogRepositoryImpl)),
	wire.Bind(new(repository.SysConfigureRepository), new(*impl.SysConfigureRepositoryImpl)),
	wire.Bind(new(repository.SysDictRepository), new(*impl.SysDictRepositoryImpl)),
	wire.Bind(new(repository.SysUserRepository), new(*impl.SysUserRepositoryImpl)),

	cache.NewSysConfigureCache,

	service.NewSysConfigureService,
	service.NewSysDictService,
	service.NewSysUserService,
	service.NewLogServer,

	controller.NewDictController,
	controller.NewBasicController,

	wire.Struct(new(App), "*"),
)

func InitializeApp() (*App, error) {
	wire.Build(Providers)
	return nil, nil
}
