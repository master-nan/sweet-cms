//go:build wireinject
// +build wireinject

/**
 * @Author: Nan
 * @Date: 2024/6/1 下午9:54
 */
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
}

var Providers = wire.NewSet(
	LoadConfig,
	InitDB,
	InitRedis,
	InitCasbin,
	InitSnowflake,
	InitValidators,
	utils.NewJWTTokenGen,
	utils.NewRedisUtil,
	impl.NewLogRepositoryImpl,
	impl.NewSysConfigureRepositoryImpl,
	impl.NewSysDictRepositoryImpl,
	impl.NewSysTableRepositoryImpl,
	impl.NewSysUserRepositoryImpl,
	impl.NewGeneralizationRepositoryImpl,
	impl.NewCasbinRuleRepositoryImpl,

	wire.Bind(new(inter.CacheInterface), new(*utils.RedisUtil)),
	wire.Bind(new(inter.TokenGenerator), new(*utils.JWTTokenGen)),
	wire.Bind(new(repository.LogRepository), new(*impl.LogRepositoryImpl)),
	wire.Bind(new(repository.SysConfigureRepository), new(*impl.SysConfigureRepositoryImpl)),
	wire.Bind(new(repository.SysDictRepository), new(*impl.SysDictRepositoryImpl)),
	wire.Bind(new(repository.SysTableRepository), new(*impl.SysTableRepositoryImpl)),
	wire.Bind(new(repository.SysUserRepository), new(*impl.SysUserRepositoryImpl)),
	wire.Bind(new(repository.GeneralizationRepository), new(*impl.GeneralizationRepositoryImpl)),
	wire.Bind(new(repository.CasbinRuleRepository), new(*impl.CasbinRuleRepositoryImpl)),

	cache.NewSysConfigureCache,
	cache.NewSysUserCache,
	cache.NewSysDictCache,
	cache.NewSysTableCache,
	cache.NewSysTableFieldCache,
	cache.NewGeneralizationCache,

	service.NewLogServer,
	service.NewSysConfigureService,
	service.NewSysDictService,
	service.NewSysTableService,
	service.NewSysUserService,

	service.NewGeneralizationService,
	service.NewCasbinRuleService,

	controller.NewDictController,
	controller.NewTableController,
	controller.NewUserController,
	controller.NewBasicController,
	controller.NewGeneralizationController,

	wire.Struct(new(App), "*"),
)

func InitializeApp() (*App, error) {
	wire.Build(Providers)
	return nil, nil
}
