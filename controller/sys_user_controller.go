/**
 * @Author: Nan
 * @Date: 2024/6/28 上午11:46
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/service"
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

func (u *UserController) QueryUser(ctx *gin.Context) {}

func (u *UserController) GetMyUserInfo(ctx *gin.Context) {
	if data, exists := ctx.Get("u"); exists {
		resp := response.NewResponse()
		ctx.Set("response", resp)
		resp.SetData(data)
		return
	} else {
		e := &response.AdminError{
			Code:    http.StatusBadRequest,
			Message: "请求参数错误",
		}
		ctx.Error(e)
		return
	}
}

func (u *UserController) GetSysUserByUserName(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	username := ctx.Param("username")
	data, err := u.sysUserService.GetByUserName(username)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.SetData(data)
	return
}

func (u *UserController) GetSysUserById(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		return
	}
	data, err := u.sysUserService.GetByUserId(id)
	if err != nil {
		ctx.Error(err)
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
	if err := ctx.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		if err == io.EOF {
			// 客户端请求体为空
			e := &response.AdminError{
				Code:    http.StatusBadRequest,
				Message: "请求参数错误",
			}
			ctx.Error(e)
			return
		}
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			// 如果是验证错误，则翻译错误信息
			var errorMessages []string
			for _, e := range ve {
				errMsg := e.Translate(translator)
				errorMessages = append(errorMessages, errMsg)
			}
			e := &response.AdminError{
				Code:    http.StatusBadRequest,
				Message: strings.Join(errorMessages, ","),
			}
			ctx.Error(e)
			return
		}
		ctx.Error(err)
		return
	}
	err := u.sysUserService.Insert(data)
	if err != nil {
		ctx.Error(err)
		return
	}
	return
}

func (u *UserController) UpdateSysUser(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
}

func (u *UserController) DeleteSysUser(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
}
