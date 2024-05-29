/**
 * @Author: Nan
 * @Date: 2023/3/13 13:09
 */

package v1

import (
	"github.com/gin-gonic/gin"
	v1 "sweet-cms/api/v1"
)

func InitBase(router *gin.RouterGroup) {
	baseRouter := router.Group("auth")
	baseRouter.POST("login", v1.NewAuthApi().Login)
	//baseRouter.POST("question/answer", v1.NewQuestionnaireAnswerApi().Create)
}
