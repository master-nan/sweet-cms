/**
 * @Author: Nan
 * @Date: 2023/3/30 23:46
 */

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"net/http"
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/model"
	"sweet-cms/server"
	"sweet-cms/utils"
)

type QuestionnaireContentApi struct {
}

func NewQuestionnaireContentApi() *QuestionnaireContentApi {
	return &QuestionnaireContentApi{}
}

func (qc *QuestionnaireContentApi) GetList(ctx *gin.Context) {
	var data request.QuestionnaireContentQueryReq
	rsp := response.NewRespData(ctx)
	if err := ctx.ShouldBindQuery(&data); err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	questionnaireServer := server.NewQuestionnaireServer(ctx)
	d, total, err := questionnaireServer.GetQuestionnaireContent(data)
	if err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	rsp.SetTotal(total).SetData(d).ReturnJson()
}

func (qc *QuestionnaireContentApi) Get(ctx *gin.Context) {
	uid := ctx.Param("uuid")
	rsp := response.NewRespData(ctx)
	questionnaireServer := server.NewQuestionnaireServer(ctx)
	abModel, err := questionnaireServer.GetQuestionnaireContentById(uid)
	if err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	rsp.SetData(abModel).ReturnJson()
}

func (qc *QuestionnaireContentApi) Create(ctx *gin.Context) {
	var data request.QuestionnaireContentCreateReq
	rsp := response.NewRespData(ctx)
	if err := ctx.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		rsp.SetCode(http.StatusBadRequest).SetMsg(err.Error()).ReturnJson()
		return
	}
	var qcModel model.QuestionnaireContent
	utils.Assignment(&data, &qcModel)
	questionnaireServer := server.NewQuestionnaireServer(ctx)
	uid, err := questionnaireServer.CreateQuestionnaireContent(qcModel)
	if err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	rsp.SetData(uid).ReturnJson()
}

func (qc *QuestionnaireContentApi) Delete(ctx *gin.Context) {
	uid := ctx.Param("uuid")
	rsp := response.NewRespData(ctx)
	questionnaireServer := server.NewQuestionnaireServer(ctx)
	err := questionnaireServer.DeleteQuestionnaireContent(uid)
	if err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	rsp.ReturnJson()
}

func (qc *QuestionnaireContentApi) Update(ctx *gin.Context) {
	var data request.QuestionnaireContentUpdateReq
	rsp := response.NewRespData(ctx)
	if err := ctx.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		rsp.SetCode(http.StatusBadRequest).SetMsg(err.Error()).ReturnJson()
		return
	}
	uid, err := uuid.Parse(ctx.Param("uuid"))
	if err != nil {
		rsp.SetCode(http.StatusBadRequest).SetMsg("资源id错误").ReturnJson()
		return
	}
	questionnaireServer := server.NewQuestionnaireServer(ctx)
	err = questionnaireServer.UpdateQuestionnaireContent(uid, data)
	if err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	rsp.ReturnJson()
}
