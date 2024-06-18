/**
 * @Author: Nan
 * @Date: 2023/3/15 11:32
 */

package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sweet-cms/form/response"
	"sweet-cms/utils"
)

const bearerLength = len("Bearer ")

func AuthHandler(jwt *utils.JWTTokenGen) gin.HandlerFunc {
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
		fmt.Println(id)
		c.Next()
	}
}
