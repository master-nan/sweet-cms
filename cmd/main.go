package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	_ "sweet-cms/docs"
	"sweet-cms/initialize"
	"syscall"
	"time"
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
	// 使用一个独立的 goroutine 启动服务器
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}
	zap.L().Info("Starting server on port", zap.Int("port", port))
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zap.L().Fatal("listen: ", zap.Error(err))
		}
	}()
	// 创建一个通道，用于接收退出信号
	quit := make(chan os.Signal, 1)
	// 接收 SIGINT 和 SIGTERM 信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // 阻塞，直到接收到信号
	zap.L().Info("Shutting down server...")
	// 创建一个5秒的超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 优雅关闭服务器
	if err := server.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server forced to shutdown:", zap.Error(err))
	}
	zap.L().Info("Server exiting")
}
