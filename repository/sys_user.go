/**
 * @Author: Nan
 * @Date: 2024/6/3 下午6:07
 */

package repository

import (
	"gorm.io/gorm"
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/model"
)

type SysUserRepository interface {
	BasicRepository
	GetByUserName(string) (model.SysUser, error)
	GetById(int) (model.SysUser, error)
	Update(*gorm.DB, request.UserUpdateReq) error
	DeleteById(*gorm.DB, int) error
	GetList(request.Basic) (response.ListResult[model.SysUser], error)
	Insert(*gorm.DB, model.SysUser) error
	GetByEmployeeID(int) (model.SysUser, error)
}
