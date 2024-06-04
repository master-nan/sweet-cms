/**
 * @Author: Nan
 * @Date: 2023/3/31 23:06
 */

package v1

import (
	"github.com/gin-gonic/gin"
)

func InitQuestionnaire(router *gin.RouterGroup) {
	//questionnaire := router.Group("questionnaire")
	//{
	//	questionnaire.GET("content", v1.NewQuestionnaireContentApi().GetList)
	//	questionnaire.GET("content/:uuid", v1.NewQuestionnaireContentApi().GetByUserName)
	//	questionnaire.PUT("content/:uuid", v1.NewQuestionnaireContentApi().UpdateSysDict)
	//	questionnaire.DELETE("content/:uuid", v1.NewQuestionnaireContentApi().DeleteSysDictById)
	//	questionnaire.POST("content", v1.NewQuestionnaireContentApi().Create)
	//
	//	questionnaire.GET("answer", v1.NewQuestionnaireAnswerApi().GetList)
	//	questionnaire.GET("answer/:id", v1.NewQuestionnaireAnswerApi().GetByUserName)
	//	questionnaire.PUT("answer/:id", v1.NewQuestionnaireAnswerApi().UpdateSysDict)
	//	questionnaire.DELETE("answer/:id", v1.NewQuestionnaireAnswerApi().DeleteSysDictById)
	//	questionnaire.POST("answer", v1.NewQuestionnaireAnswerApi().Create)
	//
	//}
}
