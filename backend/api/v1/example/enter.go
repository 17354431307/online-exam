package example

import "backend/service/example"

type ApiGroup struct {
	CustomerApi
}

var (
	customerService = example.CustomerService{}
)
