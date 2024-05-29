/**
 * @Author: Nan
 * @Date: 2023/3/15 11:32
 */

package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sweet-cms/utils"
)

const bearerLength = len("Bearer ")

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		resp := NewResponse()
		if len(authorization) < bearerLength {
			resp.SetMsg("请先登录").SetCode(http.StatusUnauthorized)
			return
		}
		token := authorization[bearerLength:]
		id, err := utils.NewJWTTokenGen().ParseToken(token)
		if err != nil {
			resp.SetMsg(err.Error()).SetCode(http.StatusForbidden)
			return
		}
		fmt.Println(id)
		c.Next()
	}
}
