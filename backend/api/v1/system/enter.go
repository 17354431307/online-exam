package system

import "backend/service"

type ApiGroup struct {
	BaseApi
}

var (
	jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService
)
