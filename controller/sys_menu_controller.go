/**
 * @Author: Nan
 * @Date: 2024/5/17 上午11:12
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

type MenuController struct {
	sysMenuService *service.SysMenuService
	translators    map[string]ut.Translator
}

func NewMenuController(sysMenuService *service.SysMenuService, translators map[string]ut.Translator) *MenuController {
	return &MenuController{
		sysMenuService,
		translators,
	}
}

func (m *MenuController) GetSysMenuById(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	data, err := m.sysMenuService.GetMenuById(id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	resp.SetData(data)
	return
}

func (m *MenuController) InsertSysMenu(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var data request.MenuCreateReq
	translator, _ := m.translators["zh"]
	err := utils.ValidatorBody[request.MenuCreateReq](ctx, &data, translator)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = m.sysMenuService.InsertMenu(ctx, data)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (m *MenuController) UpdateSysMenu(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var data request.MenuUpdateReq
	translator, _ := m.translators["zh"]
	err := utils.ValidatorBody[request.MenuUpdateReq](ctx, &data, translator)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = m.sysMenuService.UpdateMenu(ctx, data)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (m *MenuController) DeleteSysMenu(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = m.sysMenuService.DeleteMenu(ctx, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (m *MenuController) QuerySysMenu(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	//var data request.Basic
	//translator, _ := m.translators["zh"]
	//err := utils.ValidatorQuery[request.Basic](ctx, &data, translator)
	//if err != nil {
	//	_ = ctx.Error(err)
	//	return
	//}
	result, err := m.sysMenuService.GetMenuTree()
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	resp.SetData(result).SetTotal(len(result))
	return
}

func (m *MenuController) GetUserMenus(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	result, err := m.sysMenuService.GetUserMenus(id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	resp.SetData(result)
	return
}

func (m *MenuController) GetUserMenuPermissions(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	result, err := m.sysMenuService.GetUserMenuPermissions(id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	resp.SetData(result)
	return
}

func (m *MenuController) GetRoleMenus(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	roleId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	result, err := m.sysMenuService.GetRoleMenus(roleId)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	resp.SetData(result)
	return
}

// GetMyMenus 获取我的菜单
func (m *MenuController) GetMyMenus(ctx *gin.Context) {
	if data, exists := ctx.Get("user"); exists {
		resp := response.NewResponse()
		ctx.Set("response", resp)
		user, _ := data.(model.SysUser)
		result, err := m.sysMenuService.GetUserMenus(user.Id)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
		resp.SetData(result)
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
