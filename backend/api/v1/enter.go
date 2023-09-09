package v1

import (
	"backend/api/v1/example"
	"backend/api/v1/system"
)

// ApiGroup 把所有的接口统一管理起来并暴露出去
type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
