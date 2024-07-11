// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package initialize

import (
	"github.com/casbin/casbin/v2"
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
	enforcer, err := InitCasbin(db)
	if err != nil {
		return nil, err
	}
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
	sysUserCache := cache.NewSysUserCache(redisUtil)
	sysUserService := service.NewSysUserService(sysUserRepositoryImpl, snowflake, sysUserCache)
	basicController := controller.NewBasicController(jwtTokenGen, server, sysConfigureService, logService, sysUserService, v)
	basicImpl := impl.NewBasicImpl(db)
	sysTableRepositoryImpl := impl.NewSysTableRepositoryImpl(db, basicImpl)
	sysTableCache := cache.NewSysTableCache(redisUtil)
	sysTableFieldCache := cache.NewSysTableFieldCache(redisUtil)
	sysTableService := service.NewSysTableService(sysTableRepositoryImpl, snowflake, sysTableCache, sysTableFieldCache, server)
	tableController := controller.NewTableController(sysTableService, v)
	userController := controller.NewUserController(sysUserService, v)
	generalizationRepositoryImpl := impl.NewGeneralizationRepositoryImpl(db)
	generalizationService := service.NewGeneralizationService(generalizationRepositoryImpl)
	generalizationController := controller.NewGeneralizationController(generalizationService, sysTableService)
	blackCache := cache.NewBlackCache(redisUtil)
	app := &App{
		Config:                   server,
		DB:                       db,
		Redis:                    client,
		SF:                       snowflake,
		JWT:                      jwtTokenGen,
		Enforcer:                 enforcer,
		DictController:           dictController,
		BasicController:          basicController,
		TableController:          tableController,
		UserController:           userController,
		GeneralizationController: generalizationController,
		LogService:               logService,
		UserService:              sysUserService,
		BlackCache:               blackCache,
	}
	return app, nil
}

// wire.go:

type App struct {
	Config                   *config.Server
	DB                       *gorm.DB
	Redis                    *redis.Client
	SF                       *utils.Snowflake
	JWT                      *utils.JWTTokenGen
	Enforcer                 *casbin.Enforcer
	DictController           *controller.DictController
	BasicController          *controller.BasicController
	TableController          *controller.TableController
	UserController           *controller.UserController
	GeneralizationController *controller.GeneralizationController
	LogService               *service.LogService
	UserService              *service.SysUserService
	BlackCache               *cache.BlackCache
}

var Providers = wire.NewSet(
	LoadConfig,
	InitDB,
	InitRedis,
	InitCasbin,
	InitSnowflake,
	InitValidators, utils.NewJWTTokenGen, utils.NewRedisUtil, impl.NewLogRepositoryImpl, impl.NewSysConfigureRepositoryImpl, impl.NewSysDictRepositoryImpl, impl.NewSysTableRepositoryImpl, impl.NewSysUserRepositoryImpl, impl.NewGeneralizationRepositoryImpl, impl.NewCasbinRuleRepositoryImpl, impl.NewBasicImpl, wire.Bind(new(inter.CacheInterface), new(*utils.RedisUtil)), wire.Bind(new(inter.TokenGenerator), new(*utils.JWTTokenGen)), wire.Bind(new(repository.LogRepository), new(*impl.LogRepositoryImpl)), wire.Bind(new(repository.SysConfigureRepository), new(*impl.SysConfigureRepositoryImpl)), wire.Bind(new(repository.SysDictRepository), new(*impl.SysDictRepositoryImpl)), wire.Bind(new(repository.SysTableRepository), new(*impl.SysTableRepositoryImpl)), wire.Bind(new(repository.SysUserRepository), new(*impl.SysUserRepositoryImpl)), wire.Bind(new(repository.GeneralizationRepository), new(*impl.GeneralizationRepositoryImpl)), wire.Bind(new(repository.CasbinRuleRepository), new(*impl.CasbinRuleRepositoryImpl)), wire.Bind(new(repository.BasicRepository), new(*impl.BasicImpl)), cache.NewSysConfigureCache, cache.NewSysUserCache, cache.NewSysDictCache, cache.NewSysTableCache, cache.NewSysTableFieldCache, cache.NewGeneralizationCache, cache.NewBlackCache, service.NewLogServer, service.NewSysConfigureService, service.NewSysDictService, service.NewSysTableService, service.NewSysUserService, service.NewGeneralizationService, service.NewCasbinRuleService, controller.NewDictController, controller.NewTableController, controller.NewUserController, controller.NewBasicController, controller.NewGeneralizationController, wire.Struct(new(App), "*"),
)
