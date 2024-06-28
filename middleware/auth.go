/**
 * @Author: Nan
 * @Date: 2023/3/15 11:32
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sweet-cms/form/response"
	"sweet-cms/service"
	"sweet-cms/utils"
)

const bearerLength = len("Bearer ")

func AuthHandler(jwt *utils.JWTTokenGen, userService *service.SysUserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		resp := response.NewResponse()
		c.Set("response", resp)
		if len(authorization) < bearerLength {
			e := &response.AdminError{
				Code:    http.StatusUnauthorized,
				Message: "请先登录",
			}
			c.Error(e)
			return
		}
		token := authorization[bearerLength:]
		id, err := jwt.ParseToken(token)
		if err != nil {
			e := &response.AdminError{
				Code:    http.StatusForbidden,
				Message: err.Error(),
			}
			c.Error(e)
			return
		}
		i, err := strconv.Atoi(id)
		if err != nil {
			e := &response.AdminError{
				Code:    http.StatusForbidden,
				Message: err.Error(),
			}
			c.Error(e)
			return
		}
		user, err := userService.GetByUserId(i)
		if err != nil {
			e := &response.AdminError{
				Code:    http.StatusForbidden,
				Message: err.Error(),
			}
			c.Error(e)
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
