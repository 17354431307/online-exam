package example

import (
	"backend/service"
)

type ApiGroup struct {
	CustomerApi
}

var (
	customerService = service.ServiceGroupApp.ExampleServiceGroup.CustomerService
	jwtService      = service.ServiceGroupApp.SystemServiceGroup.JwtService
)
