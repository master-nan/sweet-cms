/**
 * @Author: Nan
 * @Date: 2024/5/25 下午4:22
 */

package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			// 这里你可以根据错误的类型或内容定制不同的响应
			err := c.Errors.Last().Err
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   true,
				"message": err.Error(),
			})
			c.Abort()
		}
		if resp, exists := c.Get("response"); exists {
			c.JSON(http.StatusOK, resp)
			c.Abort()
		}
	}
}
