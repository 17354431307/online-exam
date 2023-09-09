package initialize

import (
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

	//docs.SwaggerInfo.BasePath = global.OE_CONFIG.App.

	// 注册 swagger
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	global.OE_Log.Info("register swagger handler")

	PublicGroup := Router.Group("")
	{
		// 健康检测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由
		exampleRouter.InitCustomerRouter(PublicGroup)
	}

	return Router
}
