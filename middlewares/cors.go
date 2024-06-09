/**
 * @Author: Nan
 * @Date: 2023/3/13 22:26
 */

package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
	"io"
	"net/http"
)

func CorsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.S().Infof("CorsHandler start")
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		} else {
			if method == "POST" || method == "PUT" {
				var m interface{}
				err := c.ShouldBindBodyWith(&m, binding.JSON)
				if err == io.EOF {
					// 客户端请求体为空
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"success":      false,
						"errorCode":    http.StatusBadRequest,
						"errorMessage": "请求参数数据",
					})
					return
				}
			}
		}
		// 处理请求
		c.Next()
		zap.S().Infof("CorsHandler end")
	}
}
