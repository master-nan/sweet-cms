package server

import (
	"sweet-cms/global"
	"sweet-cms/model"
)

type SysServer struct {
}

func NewSysServer() *SysServer {
	return &SysServer{}
}

func (s *SysServer) GetSysUser(username string) (model.SysUser, error) {
	var user model.SysUser
	result := global.DB.Where(&model.SysUser{UserName: username}).First(&user)
	return user, result.Error
}
