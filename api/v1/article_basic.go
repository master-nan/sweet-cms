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

type ArticleBasicApi struct {
}

func NewArticleBasicApi() *ArticleBasicApi {
	return &ArticleBasicApi{}
}

func (ab ArticleBasicApi) GetList(ctx *gin.Context) {
	var data request.ArticleBasicQueryReq
	rsp := response.NewRespData(ctx)
	if err := ctx.ShouldBindQuery(&data); err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	var articleServer = server.NewArticleServer(ctx)
	d, total, err := articleServer.GetArticleBasic(data)
	if err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	rsp.SetTotal(total).SetData(d).ReturnJson()
}

func (ab ArticleBasicApi) Get(ctx *gin.Context) {
	uid := ctx.Param("uuid")
	rsp := response.NewRespData(ctx)
	var articleServer = server.NewArticleServer(ctx)
	abModel, err := articleServer.GetArticleBasicByID(uid)
	if err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	rsp.SetData(abModel).ReturnJson()
}

func (ab ArticleBasicApi) Create(ctx *gin.Context) {
	var data request.ArticleBasicCreateReq
	rsp := response.NewRespData(ctx)
	if err := ctx.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		rsp.SetCode(http.StatusBadRequest).SetMsg(err.Error()).ReturnJson()
		return
	}
	var abModel model.ArticleBasic
	utils.Assignment(&data, &abModel)
	var articleServer = server.NewArticleServer(ctx)
	uid, err := articleServer.CreateArticleBasic(abModel)
	if err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	rsp.SetData(uid).ReturnJson()
}

func (ab ArticleBasicApi) Delete(ctx *gin.Context) {
	uid := ctx.Param("uuid")
	rsp := response.NewRespData(ctx)
	var articleServer = server.NewArticleServer(ctx)
	err := articleServer.DeleteArticleBasic(uid)
	if err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	rsp.ReturnJson()
}

func (ab ArticleBasicApi) Update(ctx *gin.Context) {
	var upData request.ArticleBasicUpdateReq
	rsp := response.NewRespData(ctx)
	err := ctx.ShouldBindBodyWith(&upData, binding.JSON)
	if err != nil {
		rsp.SetCode(http.StatusBadRequest).SetMsg(err.Error()).ReturnJson()
		return
	}
	uid, err := uuid.Parse(ctx.Param("uuid"))
	if err != nil {
		rsp.SetCode(http.StatusBadRequest).SetMsg("资源id错误").ReturnJson()
		return
	}
	var articleServer = server.NewArticleServer(ctx)
	if err := articleServer.UpdateArticleBasic(uid, upData); err != nil {
		rsp.SetCode(http.StatusInternalServerError).SetMsg(err.Error()).ReturnJson()
		return
	}
	rsp.ReturnJson()
}
