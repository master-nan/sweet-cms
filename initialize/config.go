package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"reflect"
	"strings"
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

	// 绑定环境变量
	v.AutomaticEnv()
	v.SetEnvPrefix("APP")
	bindEnvs(v, &config.Server{})

	var cfg config.Server
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// 绑定结构体的所有字段到环境变量
func bindEnvs(v *viper.Viper, s interface{}, prefix ...string) {
	ps := ""
	if len(prefix) > 0 {
		ps = prefix[0] + "_"
	}
	fields := reflect.TypeOf(s).Elem()
	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		envKey := ps + strings.ToUpper(field.Name)
		if field.Type.Kind() == reflect.Struct {
			bindEnvs(v, reflect.New(field.Type).Interface(), envKey)
		} else {
			v.BindEnv(envKey)
		}
	}
}
