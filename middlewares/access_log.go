/**
 * @Author: Nan
 * @Date: 2023/3/18 17:04
 */

package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
	"sweet-cms/model"
	"sweet-cms/server"
	"time"
)

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		var body any
		var query = c.Request.URL.Query()
		_ = c.ShouldBindBodyWith(&body, binding.JSON)
		var accessLog = model.AccessLog{
			BasicModel: model.BasicModel{},
			Method:     c.Request.Method,
			Ip:         c.ClientIP(),
			Locality:   "",
			Url:        c.Request.URL.Path,
			Data:       fmt.Sprintf("body:%v，query:%v", body, query),
		}
		logServer := server.NewLogServer(c)
		_, err := logServer.CreateAccessLog(accessLog)
		if err != nil {
			zap.S().Error("日志存储异常。。。。")
		}
		c.Next()
		duration := time.Since(startTime)
		zap.S().Infof("用户访问日志:", zap.String("uri", c.Request.URL.Path), zap.Any("method", c.Request.Method), zap.Any("queryParaList", c.Request.URL.Query()), zap.String("ip", c.ClientIP()), zap.Any("duration", duration))
	}
}
