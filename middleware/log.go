/**
 * @Author: Nan
 * @Date: 2023/3/18 17:04
 */

package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
	"sweet-cms/form/response"
	"sweet-cms/model"
	"sweet-cms/service"
	"sweet-cms/utils"
	"time"
)

func LogHandler(logService *service.LogService) gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Access Log start")
		startTime := time.Now()
		var body interface{}
		var query = c.Request.URL.Query()
		_ = c.ShouldBindBodyWith(&body, binding.JSON)

		blw := &response.BufferedResponseWriter{
			ResponseWriter: c.Writer,
			Body:           bytes.NewBufferString(""),
		}
		c.Writer = blw

		c.Next()
		duration := time.Since(startTime)
		responseBody := blw.Body.String()

		queryStr, _ := json.Marshal(query)
		bodyStr, _ := json.Marshal(body)

		//responseStatus := blw.Status()
		var accessLog = model.AccessLog{
			Basic:    model.Basic{},
			Method:   c.Request.Method,
			Ip:       c.ClientIP(),
			Locality: "",
			Url:      c.Request.URL.Path,
			Body:     utils.SanitizeInput(string(bodyStr)),
			Query:    utils.SanitizeInput(string(queryStr)),
			Response: utils.SanitizeInput(responseBody),
		}
		err := logService.CreateAccessLog(c, accessLog)
		if err != nil {
			zap.L().Error("日志存储异常。。。。", zap.Error(err))
		}
		zap.L().Info("用户访问日志:",
			zap.String("uri", c.Request.URL.Path),
			zap.String("method", c.Request.Method),
			zap.Any("query", accessLog.Query),
			zap.Any("body", accessLog.Body),
			zap.String("response", accessLog.Response),
			zap.String("ip", c.ClientIP()),
			zap.String("duration", fmt.Sprintf("%.4f seconds", duration.Seconds())))
		zap.L().Info("Access Log end")
	}
}
