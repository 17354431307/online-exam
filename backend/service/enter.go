package service

import (
	"backend/service/business"
	"backend/service/example"
	"backend/service/system"
)

type ServiceGroup struct {
	ExampleServiceGroup  example.ServiceGroup
	BusinessServiceGroup business.ServiceGroup
	SystemServiceGroup   system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
