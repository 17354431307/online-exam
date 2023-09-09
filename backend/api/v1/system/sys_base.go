package system

import (
	"backend/model/common/response"
	systemReq "backend/model/system/request"
	"github.com/gin-gonic/gin"
)

type BaseApi struct {
}

func (receiver *BaseApi) Login(c *gin.Context) {
	var l systemReq.Login
	err := c.ShouldBindJSON(&l)
	//key := c.ClientIP()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//err = utils.Verify(l, utils.LoginVerify)

}
