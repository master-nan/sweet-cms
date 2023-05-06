package global

import (
	"gorm.io/gorm"
	"sweet-cms/config"
)

var (
	ServerConf  = &config.Server{}
	DBConf      = &config.DB{}
	RedisConf   = &config.Redis{}
	DB          *gorm.DB
	SessionConf = &config.Session{}
)
