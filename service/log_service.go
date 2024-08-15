/**
 * @Author: Nan
 * @Date: 2023/3/19 14:47
 */

package service

import (
	"github.com/gin-gonic/gin"
	"sweet-cms/model"
	"sweet-cms/repository"
	"sweet-cms/utils"
)

type LogService struct {
	loginLogRepository  repository.LoginLogRepository
	accessLogRepository repository.AccessLogRepository
	sf                  *utils.Snowflake
}

func NewLogServer(loginLogRepository repository.LoginLogRepository, accessLogRepository repository.AccessLogRepository, sf *utils.Snowflake) *LogService {
	return &LogService{
		loginLogRepository,
		accessLogRepository,
		sf,
	}
}

func (ls *LogService) CreateLoginLog(ctx *gin.Context, log model.LoginLog) error {
	id, err := ls.sf.GenerateUniqueID()
	if err != nil {
		return err
	}
	log.Id = int(id)
	err = ls.loginLogRepository.Create(ls.loginLogRepository.DBWithContext(ctx), &log)
	return err
}

func (ls *LogService) CreateAccessLog(ctx *gin.Context, log model.AccessLog) error {
	id, err := ls.sf.GenerateUniqueID()
	if err != nil {
		return err
	}
	log.Id = int(id)
	err = ls.accessLogRepository.Create(ls.loginLogRepository.DBWithContext(ctx), &log)
	return err
}
