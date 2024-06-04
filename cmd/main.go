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
		zap.S().Errorf("failed to initialize application: %v", err)
	}
	router := initialize.InitRouter(app)
	port := app.Config.Port
	zap.S().Infof("Starting server on port %d", port)
	err = router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		zap.S().Errorf("failed to start server: %v", err)
	}
}
