/**
 * @Author: Nan
 * @Date: 2024/6/13 下午11:30
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"sweet-cms/service"
)

type GeneralizationController struct {
	generalizationService *service.GeneralizationService
}

func NewGeneralizationController(generalizationService *service.GeneralizationService) *GeneralizationController {
	return &GeneralizationController{
		generalizationService,
	}
}

func (gc *GeneralizationController) Query(ctx *gin.Context) {

}
