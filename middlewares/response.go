/**
 * @Author: Nan
 * @Date: 2024/5/25 下午4:22
 */

package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"sweet-cms/form/response"
)

func ResponseHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				var err *response.AdminError
				switch {
				case errors.As(e.Err, &err):
					// 处理自定义API错误
					c.JSON(err.Code, gin.H{
						"success":      false,
						"errorCode":    err.Code,
						"errorMessage": err.Message,
					})
				default:
					// 处理未知错误
					c.JSON(http.StatusInternalServerError, gin.H{
						"success":      false,
						"errorCode":    http.StatusInternalServerError,
						"errorMessage": e.Error(),
					})
				}
				c.Abort()
				return
			}
		}
		if resp, exists := c.Get("response"); exists {
			c.JSON(http.StatusOK, resp)
			c.Abort()
		}
	}
}
