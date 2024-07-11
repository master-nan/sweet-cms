# sweet-cms

## 项目使用wire管理依赖注入

如果增加了控制器、服务器层及数据层代码，需要修改initialize目录下wire.go代码并在此目录下执行`wire`重新生成依赖注入代码

## 生成swagger

在根目录下执行 `swag init -g cmd/main.go -d . --parseDependency --parseInternal` 会在根目录下重新生成docs目录以及相应内容

### 重新生成各类代码完需要重启项目使生效

### 新增、删除、修改在构造函数中会自动添加创建人、修改人以及删除人信息，但要保证将控制器中的gin.Context传递到db中
    1：如果使用事务，那在创建事务的时候传递ctx
    2：如果不是用事务，那在创建db时将ctx加入，例如使用DBWithContext返回的db传入对应的方法


### 代码规范：返回前端json格式字段都使用驼峰命名