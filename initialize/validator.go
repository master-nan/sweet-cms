/**
 * @Author: Nan
 * @Date: 2024/6/7 下午10:54
 */

package initialize

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entrans "github.com/go-playground/validator/v10/translations/en"
	zhtrans "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

func InitValidators() (map[string]ut.Translator, error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		translators := make(map[string]ut.Translator)
		// 注册自定义字段名称提取函数
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		// 创建翻译器
		zhT := zh.New()
		enT := en.New()
		uni := ut.New(enT, zhT)

		// 获取翻译器
		zhTrans, _ := uni.GetTranslator("zh")
		enTrans, _ := uni.GetTranslator("en")

		// 注册翻译
		_ = zhtrans.RegisterDefaultTranslations(v, zhTrans)
		_ = entrans.RegisterDefaultTranslations(v, enTrans)
		translators["zh"] = zhTrans
		translators["en"] = enTrans
		return translators, nil
	}
	return map[string]ut.Translator{}, nil
}
