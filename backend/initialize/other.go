package initialize

import (
	"backend/global"
	"backend/utils"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"reflect"
	"strings"
)

func OtherInit() {
	//initializeValidator()

	dr, err := utils.ParseDuration(global.OE_CONFIG.Jwt.ExpiresTime)
	if err != nil {
		panic(err)
	}

	_, err = utils.ParseDuration(global.OE_CONFIG.Jwt.BufferTime)
	if err != nil {
		panic(err)
	}

	global.BlackCahe = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)

	fmt.Println(" ===== Other init ===== ")

}

func initializeValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义验证器
		_ = v.RegisterValidation("mobile", utils.ValidateMobile)

		// 注册自定义 json tag 函数
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}

			return name
		})

	}

}
