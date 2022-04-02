package validation

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

type Result map[string][]string

var (
	uni      *ut.UniversalTranslator
	Validate *validator.Validate
	Trans    ut.Translator
)

// Init 初始化比表单验证模块.
func Init() {
	// 注册翻译器
	uni = ut.New(zh.New())
	Trans, _ = uni.GetTranslator("zh")
	// 获取gin的校验器
	Validate = binding.Validator.Engine().(*validator.Validate)
	Validate.RegisterTagNameFunc(func(f reflect.StructField) string {
		name := f.Tag.Get("label")
		if name == "-" {
			return ""
		}
		return name
	})
	// 注册翻译器
	_ = zh_translations.RegisterDefaultTranslations(Validate, Trans)
}

// Translate 翻译错误信息.
func Translate(err error) Result {
	var result = make(Result)
	errors := err.(validator.ValidationErrors)
	for _, err := range errors {
		field := err.Field()
		result[field] = append(result[field], err.Translate(Trans))
	}
	return result
}
