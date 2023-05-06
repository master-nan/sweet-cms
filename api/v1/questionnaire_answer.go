/**
 * @Author: Nan
 * @Date: 2023/3/30 23:46
 */

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/model"
	"sweet-cms/server"
	"sweet-cms/utils"
)

type QuestionnaireAnswerApi struct {
}

func NewQuestionnaireAnswerApi() *QuestionnaireAnswerApi {
	return &QuestionnaireAnswerApi{}
}

func (qa *QuestionnaireAnswerApi) GetList(ctx *gin.Context) {
	var data request.QuestionnaireAnswerQueryReq
	rsp := response.NewRespData(ctx)
	if err := ctx.ShouldBindQuery(&data); err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	questionnaireServer := server.NewQuestionnaireServer(ctx)
	d, total, err := questionnaireServer.GetQuestionnaireAnswer(data)
	if err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	rsp.SetTotal(total).SetData(d).ReturnJson()
}

func (qa *QuestionnaireAnswerApi) Get(ctx *gin.Context) {
	uid, err := strconv.Atoi(ctx.Param("id"))
	rsp := response.NewRespData(ctx)
	questionnaireServer := server.NewQuestionnaireServer(ctx)
	abModel, err := questionnaireServer.GetQuestionnaireAnswerById(uid)
	if err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	rsp.SetData(abModel).ReturnJson()
}

func (qa *QuestionnaireAnswerApi) Create(ctx *gin.Context) {
	var data request.QuestionnaireAnswerCreateReq
	rsp := response.NewRespData(ctx)
	if err := ctx.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		rsp.SetCode(http.StatusBadRequest).SetMsg(err.Error()).ReturnJson()
		return
	}
	var qcModel model.QuestionnaireAnswer
	utils.Assignment(&data, &qcModel)
	questionnaireServer := server.NewQuestionnaireServer(ctx)
	uid, err := questionnaireServer.CreateQuestionnaireAnswer(qcModel)
	if err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	rsp.SetData(uid).ReturnJson()
}

func (qa *QuestionnaireAnswerApi) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	rsp := response.NewRespData(ctx)
	if err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	questionnaireServer := server.NewQuestionnaireServer(ctx)
	err = questionnaireServer.DeleteQuestionnaireAnswer(id)
	if err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	rsp.ReturnJson()
}

func (qa *QuestionnaireAnswerApi) Update(ctx *gin.Context) {
	var data request.QuestionnaireAnswerUpdateReq
	rsp := response.NewRespData(ctx)
	if err := ctx.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		rsp.SetCode(http.StatusBadRequest).SetMsg(err.Error()).ReturnJson()
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	questionnaireServer := server.NewQuestionnaireServer(ctx)
	err = questionnaireServer.UpdateQuestionnaireAnswer(id, data)
	if err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	rsp.ReturnJson()
}
