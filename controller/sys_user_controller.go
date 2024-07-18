/**
 * @Author: Nan
 * @Date: 2024/6/28 上午11:46
 */

package controller

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"net/http"
	"strconv"
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/model"
	"sweet-cms/service"
	"sweet-cms/utils"
)

type UserController struct {
	sysUserService *service.SysUserService
	translators    map[string]ut.Translator
}

func NewUserController(sysUserService *service.SysUserService, translators map[string]ut.Translator) *UserController {
	return &UserController{
		sysUserService,
		translators,
	}
}

func (u *UserController) QuerySysUser(ctx *gin.Context) {}

func (u *UserController) GetMe(ctx *gin.Context) {
	if data, exists := ctx.Get("user"); exists {
		resp := response.NewResponse()
		ctx.Set("response", resp)
		user, _ := data.(model.SysUser)
		var userRes response.UserRes
		utils.Assignment(&user, &userRes)
		resp.SetData(userRes)
		return
	} else {
		e := &response.AdminError{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "请求参数错误",
		}
		_ = ctx.Error(e)
		return
	}
}

func (u *UserController) GetSysUserByUserName(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	username := ctx.Param("username")
	data, err := u.sysUserService.GetByUserName(username)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	var userRes response.UserRes
	utils.Assignment(&data, &userRes)
	resp.SetData(userRes)
	return
}

func (u *UserController) GetSysUserById(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	data, err := u.sysUserService.GetById(id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	resp.SetData(data)
	return
}

func (u *UserController) InsertSysUser(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var data request.UserCreateReq
	translator, _ := u.translators["zh"]
	err := utils.ValidatorBody[request.UserCreateReq](ctx, &data, translator)
	err = u.sysUserService.Insert(ctx, data)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (u *UserController) UpdateSysUser(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var data request.UserUpdateReq
	translator, _ := u.translators["zh"]
	err := utils.ValidatorBody[request.UserUpdateReq](ctx, &data, translator)
	err = u.sysUserService.Update(ctx, data)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (u *UserController) DeleteSysUser(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = u.sysUserService.Delete(ctx, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}
