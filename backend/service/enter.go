package service

import (
	"backend/service/business"
	"backend/service/example"
)

type ServiceGroup struct {
	ExampleServiceGroup  example.ServiceGroup
	BusinessServiceGroup business.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
