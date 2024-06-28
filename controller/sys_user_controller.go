/**
 * @Author: Nan
 * @Date: 2024/6/28 上午11:46
 */

package controller

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
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

func (user *UserController) GetMyUserInfo(c *gin.Context) {

}

func (user *UserController) GetSysUserList(c *gin.Context) {}

func (user *UserController) GetSysUserByUserName(c *gin.Context) {}

func (user *UserController) GetSysUserById(c *gin.Context) {}

func (user *UserController) InsertSysUser(c *gin.Context) {}

func (user *UserController) UpdateSysUser(c *gin.Context) {}

func (user *UserController) DeleteSysUser(c *gin.Context) {}
