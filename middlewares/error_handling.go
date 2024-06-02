/**
 * @Author: Nan
 * @Date: 2024/5/30 下午11:21
 */

package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // 处理请求
		// 检查是否有错误信息
		if len(c.Errors) > 0 {
			// 这里你可以根据错误的类型或内容定制不同的响应
			err := c.Errors.Last().Err
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   true,
				"message": err.Error(),
			})
			c.Abort()
		}
	}
}
