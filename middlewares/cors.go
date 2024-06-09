/**
 * @Author: Nan
 * @Date: 2023/3/13 22:26
 */

package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func CorsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("CorsHandler start")
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
		zap.L().Info("CorsHandler end")
	}
}
