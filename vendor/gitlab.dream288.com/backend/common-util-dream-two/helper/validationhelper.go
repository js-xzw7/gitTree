package helper

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
)

type ValidationHelper interface {
}
type defaultValidationHelper struct {
}

func NewValidationHelper() defaultValidationHelper {
	return defaultValidationHelper{}
}

/*验参函数*/
func (h defaultValidationHelper) Validation(req interface{}) string {
	//中文翻译器
	zhCh := zh.New()
	uni := ut.New(zhCh)
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	//验证器注册翻译器
	zh_translations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(req)
	if nil != err {
		errSlice := make([]string, 0)
		for _, err := range err.(validator.ValidationErrors) {
			errSlice = append(errSlice, err.Translate(trans))
		}
		return InterfaceHelperObject.ToString(errSlice)
	}
	return ""
}
