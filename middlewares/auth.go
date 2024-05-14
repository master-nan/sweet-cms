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

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if len(authorization) < bearerLength {
			//c.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("请先登录"))
			response.NewRespData(c).SetCode(http.StatusUnauthorized).SetMsg("请先登录").ReturnJson()
		}
		token := authorization[bearerLength:]
		id, err := utils.NewJWTTokenGen().ParseToken(token)
		if err != nil {
			response.NewRespData(c).SetCode(http.StatusForbidden).SetMsg(err.Error()).AbortStatusJson()
		}
		fmt.Println(id)
		c.Next()
	}
}
