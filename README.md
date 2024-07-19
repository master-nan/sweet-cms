# Sweet-CMS

## 概述
基于 Gin, GORM, Redis, Casbin, Viper 等强大技术栈构建的高性能后台管理系统。此项目提供动态数据表操作、权限管理、配置管理以及强大的日志功能。

## 技术栈

- **Gin**: 高性能Web框架。
- **GORM**: 用于数据库操作的ORM工具。
- **Redis**: 提供缓存解决方案。
- **Casbin**: 强大的访问控制库。
- **Viper**: 配置解决方案。
- **JWT**: 安全的用户认证。
- **Zap**: 高性能日志。
- **Wire**: 依赖注入。
- **Swaggo**: 自动API文档生成。

### 项目结构说明
```
    ├── Dockerfile
    ├── LICENSE
    ├── README.md
    ├── cache                                   缓存
    │   ├── black_cache.go
    │   ├── generalization_cache.go
    │   ├── sys_configure_cache.go
    │   ├── sys_dict_cache.go
    │   ├── sys_table_cache.go   
    │   ├── sys_table_field_cache.go
    │   └── sys_user_cache.go
    ├── casbin_model.conf                        casbin权限配置
    ├── cmd
    │   └── main.go                              入口
    ├── config
    │   └── config.go                            全局配置
    ├── config-debug.yaml                        配置文件
    ├── config-local.yaml                        配置文件
    ├── config-pro.yaml                          配置文件
    ├── k8s-deployment.yaml
    ├── k8s-secret.yaml
    ├── controller                               控制器
    │   ├── basic_controller.go                  
    │   ├── generalization_controller.go
    │   ├── sys_dict_controller.go
    │   ├── sys_menu_controller.go
    │   ├── sys_table_controller.go
    │   └── sys_user_controller.go
    ├── docs                                     生成的swagger文档
    │   ├── docs.go
    │   ├── swagger.json
    │   └── swagger.yaml
    ├── enum                                     所有枚举
    │   └── enum.go
    ├── form                                    请求和响应的模型
    │   ├── request                             请求模型
    │   │   ├── basic.go
    │   │   ├── signin_req.go
    │   │   ├── sys_configure_req.go
    │   │   ├── sys_dict_req.go
    │   │   ├── sys_table_req.go
    │   │   └── sys_user_req.go
    │   └── response                            响应模型
    │       ├── base_res.go
    │       ├── signin_res.go
    │       └── user_res.go
    ├── go.mod
    ├── go.sum
    ├── initialize                              初始化目录
    │   ├── casbin.go                           权限初始化
    │   ├── config.go                           配置初始化
    │   ├── db.go                               数据库初始化
    │   ├── logger.go                           日志初始化
    │   ├── redis.go                            redis初始化
    │   ├── router.go                           路由初始化
    │   ├── sonwflake.go                        雪花算法初始化
    │   ├── validator.go                        验证初始化
    │   ├── wire.go                             依赖注入配置项
    │   └── wire_gen.go                         依赖注入程序自动生成无需修改
    ├── inter                                   接口
    │   ├── cache.go                            缓存接口
    │   └── token_generator.go                  token接口                   
    ├── middleware                              中间件
    │   ├── auth.go                             权限验证
    │   ├── cors.go                             跨域验证
    │   ├── logger.go                              日志记录
    │   └── response.go                         统一返回
    ├── model                                   系统模型
    │   ├── basic.go                            基础统一模型
    │   ├── log.go                              日志模型
    │   └── sys.go                              系统模型
    ├── repository                              数据层接口
    │   ├── basic.go                            基础类接口
    │   ├── casbin_rule.go
    │   ├── generalization.go
    │   ├── log.go
    │   ├── sys_configure.go
    │   ├── sys_dict.go
    │   ├── sys_table.go
    │   ├── sys_user.go
    │   ├── impl                                数据层接口实现
    │   │   ├── basic_impl.go                   基础接口实现
    │   │   ├── casbin_rule_impl.go
    │   │   ├── generalization_impl.go  
    │   │   ├── log_impl.go
    │   │   ├── sys_configure_impl.go
    │   │   ├── sys_dict_impl.go
    │   │   ├── sys_table_impl.go
    │   │   └── sys_user_impl.go
    │   └── util                                数据层工具类
    │       └── query.go
    ├── service                                 服务层
    │   ├── casbin_rule_service.go 
    │   ├── generalization_service.go
    │   ├── log_service.go
    │   ├── sys_configure_service.go
    │   ├── sys_dict_service.go
    │   ├── sys_table_service.go
    │   └── sys_user_service.go
    ├── test
    │   └── main.go
    └── utils                                   公共工具类
        ├── jwt.go                              jwt工具类
        ├── redis.go                            redis工具类
        ├── tools.go                            普通工具类
        └── uniqueid.go                         雪花算法工具类
```

## 开发和部署

### 依赖注入
修改`initialize/wire.go`后，在`initialize`目录下执行以下命令重新生成依赖注入代码。
```bash
wire
```

### 生成Swagger文档
在项目根目录执行：
```bash
swag init -g cmd/main.go -d . --parseDependency --parseInternal
```

### 快速开始
1. 进入`initialize`目录，执行 `wire` 生成依赖
2. 运行`go run cmd/main.go`


