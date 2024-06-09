/**
 * @Author: Nan
 * @Date: 2024/5/17 上午11:30
 */

package service

import (
	"sweet-cms/repository"
	"sweet-cms/utils"
)

type SysTableService struct {
	sysTableRepo repository.SysTableRepository
	sf           *utils.Snowflake
}

func NewSysTableService(sysTableRepo repository.SysTableRepository, sf *utils.Snowflake) *SysTableService {
	return &SysTableService{
		sysTableRepo,
		sf,
	}
}
