/**
 * @Author: Nan
 * @Date: 2023/3/19 14:47
 */

package service

import (
	"github.com/gin-gonic/gin"
	"sweet-cms/global"
	"sweet-cms/model"
)

type LogServer struct {
	ctx *gin.Context
}

func NewLogServer(ctx *gin.Context) *LogServer {
	return &LogServer{ctx: ctx}
}

func (s *LogServer) CreateLoginLog(log model.LoginLog) (int, error) {
	err := global.DB.Omit("gmt_delete").Create(&log).Error
	if err != nil {
		return 0, err
	}
	return log.ID, err
}

func (s *LogServer) CreateAccessLog(log model.AccessLog) (int, error) {
	err := global.DB.Omit("gmt_delete").Create(&log).Error
	if err != nil {
		return 0, err
	}
	return log.ID, err
}
