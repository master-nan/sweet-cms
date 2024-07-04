# sweet-cms

## 项目使用wire管理依赖注入

如果增加了控制器、服务器层及数据层代码，需要修改initialize目录下wire.go代码并在此目录下执行`wire`重新生成依赖注入代码

## 生成swagger

在根目录下执行 `swag init -g cmd/main.go -d . --parseDependency --parseInternal` 会在根目录下重新生成docs目录以及相应内容

### 重新生成各类代码完需要重启项目使生效


### 代码规范：返回前端json格式字段都使用驼峰命名