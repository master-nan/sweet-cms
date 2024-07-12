/**
 * @Author: Nan
 * @Date: 2024/5/25 下午4:22
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
	"sweet-cms/form/response"
)

func ResponseHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("ResponseHandler start")
		c.Next()
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				var err *response.AdminError
				switch {
				case errors.As(e.Err, &err):
					// 处理自定义API错误
					err.Success = false
					c.JSON(err.ErrorCode, err)
				default:
					// 处理未知错误
					newErr := &response.AdminError{
						ErrorCode:    err.ErrorCode,
						ErrorMessage: e.Error(),
						Success:      false,
					}
					c.JSON(http.StatusInternalServerError, newErr)
				}
				return
			}
		} else {
			if resp, exists := c.Get("response"); exists {
				c.JSON(http.StatusOK, resp)
			} else {
				c.JSON(http.StatusOK, gin.H{})
			}
		}
		zap.L().Info("ResponseHandler end")
	}
}
