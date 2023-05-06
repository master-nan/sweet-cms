/**
 * @Author: Nan
 * @Date: 2023/3/31 23:06
 */

package v1

import (
	"github.com/gin-gonic/gin"
	"sweet-cms/api/v1"
)

func InitQuestionnaire(router *gin.RouterGroup) {
	questionnaire := router.Group("questionnaire")
	{
		questionnaire.GET("content", v1.NewQuestionnaireContentApi().GetList)
		questionnaire.GET("content/:uuid", v1.NewQuestionnaireContentApi().Get)
		questionnaire.PUT("content/:uuid", v1.NewQuestionnaireContentApi().Update)
		questionnaire.DELETE("content/:uuid", v1.NewQuestionnaireContentApi().Delete)
		questionnaire.POST("content", v1.NewQuestionnaireContentApi().Create)

		questionnaire.GET("answer", v1.NewQuestionnaireAnswerApi().GetList)
		questionnaire.GET("answer/:id", v1.NewQuestionnaireAnswerApi().Get)
		questionnaire.PUT("answer/:id", v1.NewQuestionnaireAnswerApi().Update)
		questionnaire.DELETE("answer/:id", v1.NewQuestionnaireAnswerApi().Delete)
		questionnaire.POST("answer", v1.NewQuestionnaireAnswerApi().Create)

	}
}
