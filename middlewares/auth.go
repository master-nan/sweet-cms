/**
 * @Author: Nan
 * @Date: 2023/3/15 11:32
 */

package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sweet-cms/form/response"
	"sweet-cms/utils"
)

const bearerLength = len("Bearer ")

func Auth(jwt *utils.JWTTokenGen) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		resp := response.NewResponse()
		c.Set("response", resp)
		if len(authorization) < bearerLength {
			resp.SetMsg("请先登录").SetCode(http.StatusUnauthorized)
			return
		}
		token := authorization[bearerLength:]
		id, err := jwt.ParseToken(token)
		if err != nil {
			resp.SetMsg(err.Error()).SetCode(http.StatusForbidden)
			return
		}
		fmt.Println(id)
		c.Next()
	}
}
