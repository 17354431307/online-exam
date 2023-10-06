package example

import (
	"backend/global"
	"backend/model/common/response"
	"backend/model/example"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CustomerApi struct {
}

// CreateExaCustomer
// @Tags      ExaCustomer
// @Summary   创建客户
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.CreateExaCustomerRequest            true  "客户用户名, 客户手机号码"
// @Success   200   {object}  response.Response{msg=string}  "创建客户"
// @Router    /customer/customer [post]
func (e *CustomerApi) CreateExaCustomer(c *gin.Context) {
	var customer example.ExaCustomer

	err := c.ShouldBindJSON(&customer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = customerService.CreateExaCustomer(customer)
	if err != nil {
		global.OE_Log.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

func (e *CustomerApi) HelloWord(c *gin.Context) {
	response.OkWithMessage("hello, world", c)
}
