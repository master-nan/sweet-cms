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

func AuthHandler(jwt *utils.JWTTokenGen) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		resp := response.NewResponse()
		c.Set("response", resp)
		if len(authorization) < bearerLength {
			resp.SetErrorMessage("请先登录").SetErrorCode(http.StatusUnauthorized)
			return
		}
		token := authorization[bearerLength:]
		id, err := jwt.ParseToken(token)
		if err != nil {
			resp.SetErrorMessage(err.Error()).SetErrorCode(http.StatusForbidden)
			return
		}
		fmt.Println(id)
		c.Next()
	}
}
