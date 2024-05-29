package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/model"
	"sweet-cms/service"
	"sweet-cms/utils"
)

type ArticleContentApi struct {
}

func NewArticleContentApi() *ArticleContentApi {
	return &ArticleContentApi{}
}

func (ac ArticleContentApi) Get(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	rsp := response.NewRespData(ctx)
	var articleServer = service.NewArticleServer(ctx)
	acModel, err := articleServer.GetArticleContentByID(id)
	if err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	rsp.SetData(acModel).ReturnJson()
}

func (ac ArticleContentApi) Create(ctx *gin.Context) {
	var data request.ArticleChannelCreateReq
	rsp := response.NewRespData(ctx)
	var articleServer = service.NewArticleServer(ctx)
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		rsp.SetCode(http.StatusBadRequest).SetMsg(err.Error()).ReturnJson()
		return
	}
	var acModel model.ArticleContent
	utils.Assignment(&data, &acModel)
	id, err := articleServer.CreateArticleContent(acModel)
	if err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	rsp.SetData(id).ReturnJson()
}

func (ac ArticleContentApi) Update(ctx *gin.Context) {
	var upData request.ArticleContentUpdateReq
	rsp := response.NewRespData(ctx)
	var articleServer = service.NewArticleServer(ctx)
	err := ctx.ShouldBindJSON(&upData)
	if err != nil {
		rsp.SetCode(http.StatusBadRequest).SetMsg(err.Error()).ReturnJson()
		return
	}
	uid, err := uuid.Parse(ctx.Param("uuid"))
	if err != nil {
		rsp.SetCode(http.StatusBadRequest).SetMsg("资源id错误").ReturnJson()
		return
	}
	if err := articleServer.UpdateArticleContent(uid, upData); err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	rsp.ReturnJson()
}

func (ac ArticleContentApi) Delete(ctx *gin.Context) {
	var articleServer = service.NewArticleServer(ctx)
	rsp := response.NewRespData(ctx)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	err = articleServer.DeleteArticleContent(id)
	if err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	rsp.ReturnJson()
}
