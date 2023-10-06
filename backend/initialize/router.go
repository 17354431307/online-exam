package initialize

import (
	"backend/docs"
	_ "backend/docs"
	"backend/global"
	"backend/router"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example

	docs.SwaggerInfo.BasePath = global.OE_CONFIG.App.RouterPrefix
	// 注册 swagger
	Router.GET(global.OE_CONFIG.App.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	global.OE_Log.Info("register swagger handler")

	// 方便统一添加路由前缀 多服务器上线使用
	PublicGroup := Router.Group(global.OE_CONFIG.App.RouterPrefix)
	{
		// 健康检测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		exampleRouter.InitCustomerRouter(PublicGroup)
	}

	return Router
}
