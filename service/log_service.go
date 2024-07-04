/**
 * @Author: Nan
 * @Date: 2023/3/19 14:47
 */

package service

import (
	"sweet-cms/model"
	"sweet-cms/repository"
	"sweet-cms/utils"
)

type LogService struct {
	logRepository repository.LogRepository
	sf            *utils.Snowflake
}

func NewLogServer(logRepository repository.LogRepository, sf *utils.Snowflake) *LogService {
	return &LogService{logRepository, sf}
}

func (ls *LogService) CreateLoginLog(log model.LoginLog) error {
	id, err := ls.sf.GenerateUniqueID()
	if err != nil {
		return err
	}
	log.Id = int(id)
	err = ls.logRepository.CreateLoginLog(log)
	return err
}

func (ls *LogService) CreateAccessLog(log model.AccessLog) error {
	id, err := ls.sf.GenerateUniqueID()
	if err != nil {
		return err
	}
	log.Id = int(id)
	err = ls.logRepository.CreateAccessLog(log)
	return err
}
