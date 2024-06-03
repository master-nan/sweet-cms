// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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

// Injectors from wire.go:

func InitializeApp() (*App, error) {
	server, err := LoadConfig()
	if err != nil {
		return nil, err
	}
	db, err := InitDB(server)
	if err != nil {
		return nil, err
	}
	client, err := InitRedis(server)
	if err != nil {
		return nil, err
	}
	snowflake, err := InitSnowflake()
	if err != nil {
		return nil, err
	}
	jwtTokenGen := utils.NewJWTTokenGen()
	sysDictRepositoryImpl := impl.NewSysDictRepositoryImpl()
	sysDictService := service.NewSysDictService(sysDictRepositoryImpl)
	dictController := controller.NewDictController(sysDictService)
	sysConfigureRepositoryImpl := impl.NewSysConfigureRepositoryImpl(db)
	redisUtil := utils.NewRedisUtil(client)
	sysConfigureCache := cache.NewSysConfigureCache(redisUtil)
	sysConfigureService := service.NewSysConfigureService(sysConfigureRepositoryImpl, sysConfigureCache)
	logRepositoryImpl := impl.NewLogRepositoryImpl(db)
	logService := service.NewLogServer(logRepositoryImpl, snowflake)
	sysUserRepositoryImpl := impl.NewSysUserRepositoryImpl(db)
	sysUserService := service.NewSysUserService(sysUserRepositoryImpl)
	basicController := controller.NewBasicController(jwtTokenGen, server, sysConfigureService, logService, sysUserService)
	app := &App{
		Config:          server,
		DB:              db,
		Redis:           client,
		SF:              snowflake,
		JWT:             jwtTokenGen,
		DictController:  dictController,
		BasicController: basicController,
		LogService:      logService,
	}
	return app, nil
}

// wire.go:

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
	InitSnowflake, utils.NewJWTTokenGen, utils.NewRedisUtil, impl.NewSysConfigureRepositoryImpl, impl.NewSysDictRepositoryImpl, impl.NewLogRepositoryImpl, impl.NewSysUserRepositoryImpl, wire.Bind(new(inter.CacheInterface), new(*utils.RedisUtil)), wire.Bind(new(inter.TokenGenerator), new(*utils.JWTTokenGen)), wire.Bind(new(repository.LogRepository), new(*impl.LogRepositoryImpl)), wire.Bind(new(repository.SysConfigureRepository), new(*impl.SysConfigureRepositoryImpl)), wire.Bind(new(repository.SysDictRepository), new(*impl.SysDictRepositoryImpl)), wire.Bind(new(repository.SysUserRepository), new(*impl.SysUserRepositoryImpl)), cache.NewSysConfigureCache, service.NewSysConfigureService, service.NewSysDictService, service.NewSysUserService, service.NewLogServer, controller.NewDictController, controller.NewBasicController, wire.Struct(new(App), "*"),
)
