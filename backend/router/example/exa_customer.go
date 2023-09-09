package example

import (
	v1 "backend/api/v1"
	"github.com/gin-gonic/gin"
)

type CustomerRouter struct {
}

func (e *CustomerRouter) InitCustomerRouter(Router *gin.RouterGroup) gin.IRouter {
	customerRouter := Router.Group("customer")

	exaCustomerApi := v1.ApiGroupApp.ExampleApiGroup.CustomerApi
	{
		customerRouter.POST("customer", exaCustomerApi.CreateExaCustomer)
	}

	return customerRouter
}
