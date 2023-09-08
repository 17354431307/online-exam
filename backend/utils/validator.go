package utils

import (
	"github.com/go-playground/validator/v10"
	"log"
	"regexp"
)

type Validator interface {
	GetMessages() ValidatorMessages
}

type ValidatorMessages map[string]string

// GetErrorMsg 获取错误信息
func GetErrorMsg(request any, err error) string {
	if _, isValidatorErrors := err.(validator.ValidationErrors); isValidatorErrors {
		_, isValidator := request.(Validator)

		for _, v := range err.(validator.ValidationErrors) {
			// 若 request 结构体实现 Validator 接口即可实现自定义错误信息
			if isValidator {
				log.Println(v.Field() + "." + v.Tag())
				if message, exists := request.(Validator).GetMessages()[v.Field()+"."+v.Tag()]; exists {
					return message
				}
			}

			return v.Error()
		}
	}

	return "Parameter error"
}

// ValidateMobile 检验手机号
func ValidateMobile(f1 validator.FieldLevel) bool {
	mobile := f1.Field().String()
	ok, _ := regexp.MatchString(`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`, mobile)
	if !ok {
		return false
	}

	return true
}
