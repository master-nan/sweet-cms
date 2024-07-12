/**
 * @Author: Nan
 * @Date: 2023/3/15 11:32
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"sweet-cms/form/response"
	"sweet-cms/inter"
	"sweet-cms/service"
)

const bearerLength = len("Bearer ")

func AuthHandler(tokenGenerator inter.TokenGenerator, userService *service.SysUserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		resp := response.NewResponse()
		c.Set("response", resp)
		zap.L().Info("AuthHandler start")
		if len(authorization) < bearerLength {
			e := &response.AdminError{
				Code:    http.StatusUnauthorized,
				Message: "请先登录",
			}
			c.Error(e)
			c.Abort()
			return
		}
		token := authorization[bearerLength:]
		id, err := tokenGenerator.ParseToken(token)
		if err != nil {
			e := &response.AdminError{
				Code:    http.StatusForbidden,
				Message: err.Error(),
			}
			c.Error(e)
			c.Abort()
			return
		}
		i, err := strconv.Atoi(id)
		if err != nil {
			e := &response.AdminError{
				Code:    http.StatusForbidden,
				Message: err.Error(),
			}
			c.Error(e)
			c.Abort()
			return
		}
		user, err := userService.GetById(i)
		if err != nil {
			e := &response.AdminError{
				Code:    http.StatusForbidden,
				Message: err.Error(),
			}
			c.Error(e)
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Set("id", id)
		c.Next()
		zap.L().Info("AuthHandler end")
	}
}
