package global

import (
	"gorm.io/gorm"
	"sweet-cms/config"
	"sweet-cms/utils"
)

var (
	ServerConf = &config.Server{}
	DB         *gorm.DB
	SF         *utils.Snowflake
)
