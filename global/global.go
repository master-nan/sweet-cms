package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"sweet-cms/config"
	"sweet-cms/utils"
)

var (
	ServerConf  = &config.Server{}
	DB          *gorm.DB
	RedisClient *redis.Client
	SF          *utils.Snowflake
)
