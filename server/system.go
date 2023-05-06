package server

import (
	"sweet-cms/global"
	"sweet-cms/model"
)

type SystemServer struct {
}

func NewSystemServer() *SystemServer {
	return &SystemServer{}
}

func (s *SystemServer) GetSysUser(username string) (model.SysUser, error) {
	var user model.SysUser
	result := global.DB.Where(&model.SysUser{UserName: username}).First(&user)
	return user, result.Error
}
