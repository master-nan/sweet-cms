/**
 * @Author: Nan
 * @Date: 2024/8/2 上午10:18
 */

package controller

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"net/http"
	"strconv"
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/service"
	"sweet-cms/utils"
)

type RoleController struct {
	sysRoleService *service.SysRoleService
	translators    map[string]ut.Translator
}

func NewRoleController(sysRoleService *service.SysRoleService, translators map[string]ut.Translator) *RoleController {
	return &RoleController{
		sysRoleService: sysRoleService,
		translators:    translators,
	}
}

func (r *RoleController) QueryRole(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var data request.Basic
	translator, _ := r.translators["zh"]
	err := utils.ValidatorQuery[request.Basic](ctx, &data, translator)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	result, err := r.sysRoleService.GetRoleList(data)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	resp.SetData(result.Data).SetTotal(result.Total)
	return
}

func (r *RoleController) GetRoleById(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	role, err := r.sysRoleService.GetRoleById(id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	resp.SetData(role)
	return
}

func (r *RoleController) CreateRole(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var data request.RoleCreateReq
	translator, _ := r.translators["zh"]
	err := utils.ValidatorBody[request.RoleCreateReq](ctx, &data, translator)
	err = r.sysRoleService.CreateRole(ctx, data)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (r *RoleController) UpdateRole(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var data request.RoleUpdateReq
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		e := &response.AdminError{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: err.Error(),
		}
		_ = ctx.Error(e)
		return
	}
	data.Id = id
	translator, _ := r.translators["zh"]
	err = utils.ValidatorBody[request.RoleUpdateReq](ctx, &data, translator)
	err = r.sysRoleService.UpdateRole(ctx, data)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
}

func (r *RoleController) DeleteRoleById(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = r.sysRoleService.DeleteRole(ctx, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (r *RoleController) GetRoleMenus(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	roleId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	result, err := r.sysRoleService.GetRoleMenus(roleId)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	resp.SetData(result)
	return
}

func (r *RoleController) CreateRoleMenu(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	roleId, err := strconv.Atoi(ctx.Param("roleId"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	menuId, err := strconv.Atoi(ctx.Param("menuId"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	data := request.RoleMenuCreateReq{
		MenuId: menuId,
		RoleId: roleId,
	}
	err = r.sysRoleService.CreateRoleMenu(ctx, data)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (r *RoleController) GetRoleMenuButtons(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	roleId, err := strconv.Atoi(ctx.Param("roleId"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	menuId, err := strconv.Atoi(ctx.Param("menuId"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	result, err := r.sysRoleService.GetRoleMenuButtons(roleId, menuId)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	resp.SetData(result)
	return
}
