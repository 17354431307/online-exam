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

func (e *CustomerApi) CreateExaCustomer(c *gin.Context) {
	var customer example.ExaCusmoter
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
