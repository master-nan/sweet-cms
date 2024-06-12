// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package initialize

import (
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
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

import (
	_ "sweet-cms/docs"
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
	sysDictRepositoryImpl := impl.NewSysDictRepositoryImpl(db)
	redisUtil := utils.NewRedisUtil(client)
	sysDictCache := cache.NewSysDictCache(redisUtil)
	sysDictService := service.NewSysDictService(sysDictRepositoryImpl, snowflake, sysDictCache)
	v, err := InitValidators()
	if err != nil {
		return nil, err
	}
	dictController := controller.NewDictController(sysDictService, v)
	sysConfigureRepositoryImpl := impl.NewSysConfigureRepositoryImpl(db)
	sysConfigureCache := cache.NewSysConfigureCache(redisUtil)
	sysConfigureService := service.NewSysConfigureService(sysConfigureRepositoryImpl, sysConfigureCache)
	logRepositoryImpl := impl.NewLogRepositoryImpl(db)
	logService := service.NewLogServer(logRepositoryImpl, snowflake)
	sysUserRepositoryImpl := impl.NewSysUserRepositoryImpl(db)
	sysUserService := service.NewSysUserService(sysUserRepositoryImpl)
	basicController := controller.NewBasicController(jwtTokenGen, server, sysConfigureService, logService, sysUserService, v)
	sysTableRepositoryImpl := impl.NewSysTableRepositoryImpl(db)
	sysTableCache := cache.NewSysTableCache(redisUtil)
	sysTableService := service.NewSysTableService(sysTableRepositoryImpl, snowflake, sysTableCache)
	tableController := controller.NewTableController(sysTableService, v)
	app := &App{
		Config:          server,
		DB:              db,
		Redis:           client,
		SF:              snowflake,
		JWT:             jwtTokenGen,
		DictController:  dictController,
		BasicController: basicController,
		TableController: tableController,
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
	TableController *controller.TableController
	LogService      *service.LogService
}

var Providers = wire.NewSet(
	LoadConfig,
	InitDB,
	InitRedis,
	InitSnowflake,
	InitValidators, utils.NewJWTTokenGen, utils.NewRedisUtil, impl.NewLogRepositoryImpl, impl.NewSysConfigureRepositoryImpl, impl.NewSysDictRepositoryImpl, impl.NewSysTableRepositoryImpl, impl.NewSysUserRepositoryImpl, wire.Bind(new(inter.CacheInterface), new(*utils.RedisUtil)), wire.Bind(new(inter.TokenGenerator), new(*utils.JWTTokenGen)), wire.Bind(new(repository.LogRepository), new(*impl.LogRepositoryImpl)), wire.Bind(new(repository.SysConfigureRepository), new(*impl.SysConfigureRepositoryImpl)), wire.Bind(new(repository.SysDictRepository), new(*impl.SysDictRepositoryImpl)), wire.Bind(new(repository.SysTableRepository), new(*impl.SysTableRepositoryImpl)), wire.Bind(new(repository.SysUserRepository), new(*impl.SysUserRepositoryImpl)), cache.NewSysConfigureCache, cache.NewSysDictCache, cache.NewSysTableCache, service.NewLogServer, service.NewSysConfigureService, service.NewSysDictService, service.NewSysTableService, service.NewSysUserService, controller.NewDictController, controller.NewTableController, controller.NewBasicController, wire.Struct(new(App), "*"),
)
