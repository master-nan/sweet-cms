package main

import (
	"fmt"
	"go.uber.org/zap"
	_ "sweet-cms/docs"
	"sweet-cms/initialize"
)

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
