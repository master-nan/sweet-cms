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
	BlackCache               *cache.BlackCache
}

// Repository providers
var RepositoryProvider = wire.NewSet(
	impl.NewAccessLogRepositoryImpl,
	impl.NewLoginLogRepositoryImpl,
	impl.NewSysConfigureRepositoryImpl,
	impl.NewSysDictRepositoryImpl,
	impl.NewSysDictItemRepositoryImpl,
	impl.NewSysTableIndexFieldRepositoryImpl,
	impl.NewSysTableIndexRepositoryImpl,
	impl.NewSysTableRelationRepositoryImpl,
	impl.NewSysTableFieldRepositoryImpl,
	impl.NewSysTableRepositoryImpl,
	impl.NewSysUserRepositoryImpl,
	impl.NewSysMenuRepositoryImpl,
	impl.NewSysRoleRepositoryImpl,
	impl.NewSysRoleMenuButtonRepositoryImpl,
	impl.NewSysRoleMenuRepositoryImpl,
	impl.NewSysUserMenuDataPermissionRepositoryImpl,
	impl.NewSysUserRoleRepositoryImpl,

	impl.NewGeneralizationRepositoryImpl,
	impl.NewCasbinRuleRepositoryImpl,
	impl.NewBasicImpl,

	wire.Bind(new(repository.AccessLogRepository), new(*impl.AccessLogRepositoryImpl)),
	wire.Bind(new(repository.LoginLogRepository), new(*impl.LoginLogRepositoryImpl)),
	wire.Bind(new(repository.SysConfigureRepository), new(*impl.SysConfigureRepositoryImpl)),
	wire.Bind(new(repository.SysDictRepository), new(*impl.SysDictRepositoryImpl)),
	wire.Bind(new(repository.SysDictItemRepository), new(*impl.SysDictItemRepositoryImpl)),
	wire.Bind(new(repository.SysTableIndexFieldRepository), new(*impl.SysTableIndexFieldRepositoryImpl)),
	wire.Bind(new(repository.SysTableIndexRepository), new(*impl.SysTableIndexRepositoryImpl)),
	wire.Bind(new(repository.SysTableRelationRepository), new(*impl.SysTableRelationRepositoryImpl)),
	wire.Bind(new(repository.SysTableFieldRepository), new(*impl.SysTableFieldRepositoryImpl)),
	wire.Bind(new(repository.SysTableRepository), new(*impl.SysTableRepositoryImpl)),
	wire.Bind(new(repository.SysUserRepository), new(*impl.SysUserRepositoryImpl)),
	wire.Bind(new(repository.SysMenuRepository), new(*impl.SysMenuRepositoryImpl)),
	wire.Bind(new(repository.SysRoleRepository), new(*impl.SysRoleRepositoryImpl)),
	wire.Bind(new(repository.SysRoleMenuButtonRepository), new(*impl.SysRoleMenuButtonRepositoryImpl)),
	wire.Bind(new(repository.SysRoleMenuRepository), new(*impl.SysRoleMenuRepositoryImpl)),
	wire.Bind(new(repository.SysUserMenuDataPermissionRepository), new(*impl.SysUserMenuDataPermissionRepositoryImpl)),
	wire.Bind(new(repository.SysUserRoleRepository), new(*impl.SysUserRoleRepositoryImpl)),

	wire.Bind(new(repository.GeneralizationRepository), new(*impl.GeneralizationRepositoryImpl)),
	wire.Bind(new(repository.CasbinRuleRepository), new(*impl.CasbinRuleRepositoryImpl)),
	wire.Bind(new(repository.BasicRepository), new(*impl.BasicImpl)),
)

// Cache providers
var CacheProvider = wire.NewSet(
	cache.NewSysConfigureCache,
	cache.NewSysUserCache,
	cache.NewSysDictCache,
	cache.NewSysTableCache,
	cache.NewSysTableFieldCache,
	cache.NewGeneralizationCache,
	cache.NewBlackCache,
)

// Service providers
var ServiceProvider = wire.NewSet(
	service.NewLogServer,
	service.NewSysConfigureService,
	service.NewSysDictService,
	service.NewSysRoleService,
	service.NewSysMenuService,
	service.NewSysTableService,
	service.NewSysUserService,
	service.NewGeneralizationService,
	service.NewCasbinRuleService,
)

// Controller providers
var ControllerProvider = wire.NewSet(
	controller.NewDictController,
	controller.NewTableController,
	controller.NewUserController,
	controller.NewBasicController,
	controller.NewGeneralizationController,
)

var Providers = wire.NewSet(
	LoadConfig,
	InitDB,
	InitRedis,
	InitCasbin,
	InitSnowflake,
	InitValidators,
	utils.NewJWTTokenGen,
	utils.NewRedisUtil,
	wire.Bind(new(inter.CacheInterface), new(*utils.RedisUtil)),
	wire.Bind(new(inter.TokenGenerator), new(*utils.JWTTokenGen)),

	RepositoryProvider,
	CacheProvider,
	ServiceProvider,
	ControllerProvider,

	wire.Struct(new(App), "*"),
)

func InitializeApp() (*App, error) {
	wire.Build(Providers)
	return nil, nil
}
