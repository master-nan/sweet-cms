package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"sweet-cms/config"
)

func LoadConfig() (*config.Server, error) {
	environment := os.Getenv("APP_ENV")
	if environment == "" {
		environment = "local" // 默认使用本地环境的配置
	}
	filename := fmt.Sprintf("config-%s.yaml", environment)
	v := viper.New()
	v.SetConfigFile(filename)
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %w", filename, err)
	}
	var cfg config.Server
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
