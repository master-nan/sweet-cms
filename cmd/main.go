package main

import (
	"fmt"
	"go.uber.org/zap"
	_ "sweet-cms/docs"
	"sweet-cms/initialize"
)

// @title 试验性项目
// @version 0.1
// @description 基于gin+gorm的后台管理项目，实现部分低代码

// @contact.name 南
// @contact.email maxdwy@gmail.com

// @BasePath  /sweet/admin

func main() {
	initialize.Logger() // 初始化日志
	app, err := initialize.InitializeApp()
	if err != nil {
		zap.L().Error("failed to initialize application", zap.Error(err))
	}
	router := initialize.InitRouter(app)
	port := app.Config.Port
	zap.L().Info("Starting server on port", zap.Int("port", port))
	err = router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		zap.L().Error("failed to start server: %v", zap.Error(err))
	}
}
