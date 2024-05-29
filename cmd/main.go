package main

import (
	"fmt"
	"go.uber.org/zap"
	"sweet-cms/global"
	"sweet-cms/initialize"
)

func init() {
	initialize.Config()      // 初始化配置
	initialize.DB()          // 初始化db
	initialize.RedisClient() //初始化redis
	initialize.Logger()      // 初始化日志
	initialize.SF()

}
func main() {
	//initialize.Config()            // 初始化配置
	//initialize.DB()                // 初始化db
	//initialize.Logger()            // 初始化日志
	router := initialize.Routers() //初始化路由
	//templates := initialize.LoadTemplates()
	//router.HTMLRender = templates
	err := router.Run(fmt.Sprintf(":%d", global.ServerConf.Port))
	if err != nil {
		zap.S().Error("项目启动失败……")
	}
}
