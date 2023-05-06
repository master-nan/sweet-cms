/**
 * @Author: Nan
 * @Date: 2023/3/15 11:32
 */

package middlewares

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sweet-cms/form/response"
	"sweet-cms/utils"
)

const bearerLength = len("Bearer ")

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if len(authorization) < bearerLength {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("请先登录"))
			return
		}
		token := authorization[bearerLength:]
		id, err := utils.NewJWTTokenGen().ParseToken(token)
		if err != nil {
			response.NewRespData(c).SetCode(http.StatusForbidden).SetMsg(err.Error()).AbortStatusJson()
			return
		}
		fmt.Println(id)
		c.Next()
	}
}
