package middleware

import (
	"backend/global"
	"backend/model/common/response"
	"backend/service"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.OE_CONFIG.App.Env != "develop" {
			waitUse, _ := utils.GetClaims(c)
			// 获取请求的 Path
			path := c.Request.URL.Path
			obj := strings.TrimPrefix(path, global.OE_CONFIG.App.RouterPrefix)
			// 获取请求方法
			act := c.Request.Method
			// 获取用户的角色
			sub := strconv.Itoa(int(waitUse.AuthorityId))

			// 判断策略中是否存在
			e := casbinService.Casbin()
			success, _ := e.Enforce(sub, obj, act)
			if !success {
				response.FailWithDetailed(gin.H{}, "权限不足", c)
				c.Abort()
				return
			}
			c.Next()
		}
	}
}
