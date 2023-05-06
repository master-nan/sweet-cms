package initialize

import (
	"github.com/spf13/viper"
	"sweet-cms/global"
)

func Config() {
	filename := "config-debug.yaml"
	v := viper.New()
	v.SetConfigFile(filename)

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(global.ServerConf); err != nil {
		panic(err)
	}
}
