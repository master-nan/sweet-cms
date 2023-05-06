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

type ArticleChannelApi struct {
}

func NewArticleChannelApi() *ArticleChannelApi {
	return &ArticleChannelApi{}
}

func (ac ArticleChannelApi) GetList(ctx *gin.Context) {
	var query request.ArticleChannelQueryReq
	rsp := response.NewRespData(ctx)
	err := ctx.ShouldBindQuery(&query)
	if err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	var articleServer = server.NewArticleServer(ctx)
	d, count, err := articleServer.GetArticleChannel(query)
	if err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	rsp.SetTotal(count).SetData(d).ReturnJson()
}

func (ac ArticleChannelApi) Get(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	rsp := response.NewRespData(ctx)
	var articleServer = server.NewArticleServer(ctx)
	acModel, err := articleServer.GetArticleChannelByID(id)
	if err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	rsp.SetData(acModel).ReturnJson()
}

func (ac ArticleChannelApi) Create(ctx *gin.Context) {
	var data request.ArticleChannelCreateReq
	rsp := response.NewRespData(ctx)
	var articleServer = server.NewArticleServer(ctx)
	err := ctx.ShouldBindBodyWith(&data, binding.JSON)
	if err != nil {
		rsp.SetCode(http.StatusBadRequest).SetMsg(err.Error()).ReturnJson()
		return
	}
	var acModel model.ArticleChannel
	utils.Assignment(&data, &acModel)
	id, err := articleServer.CreateArticleChannel(acModel)
	if err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	rsp.SetData(id).ReturnJson()
}

func (ac ArticleChannelApi) Update(ctx *gin.Context) {
	var upData request.ArticleChannelUpdateReq
	rsp := response.NewRespData(ctx)
	var articleServer = server.NewArticleServer(ctx)
	err := ctx.ShouldBindBodyWith(&upData, binding.JSON)
	if err != nil {
		rsp.SetCode(http.StatusBadRequest).SetMsg(err.Error()).ReturnJson()
		return
	}
	if err := articleServer.UpdateArticleChannel(upData); err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	rsp.ReturnJson()
}

func (ac ArticleChannelApi) Delete(ctx *gin.Context) {
	var articleServer = server.NewArticleServer(ctx)
	rsp := response.NewRespData(ctx)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	err = articleServer.DeleteArticleChannel(id)
	if err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	rsp.ReturnJson()
}
