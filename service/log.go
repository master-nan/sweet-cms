/**
 * @Author: Nan
 * @Date: 2023/3/19 14:47
 */

package service

import (
	"sweet-cms/model"
	"sweet-cms/repository"
)

type LogService struct {
	logRepository repository.LogRepository
}

func NewLogServer(logRepository repository.LogRepository) *LogService {
	return &LogService{logRepository}
}

func (ls *LogService) CreateLoginLog(log model.LoginLog) (int, error) {
	id, err := ls.logRepository.CreateLoginLog(log)
	return id, err
}

func (ls *LogService) CreateAccessLog(log model.AccessLog) (int, error) {
	id, err := ls.logRepository.CreateAccessLog(log)
	return id, err
}
