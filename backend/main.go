package main

import (
	"backend/core"
	"backend/global"
	"backend/initialize"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const AppMode = "debug"

func main() {
	gin.SetMode(AppMode)

	// TODO: 1. 配置初始化
	global.OE_VIPER = core.InitViper()

	// TODO: 2. 日志
	global.OE_Log = core.InitializeZap()
	zap.ReplaceGlobals(global.OE_Log)

	global.OE_Log.Info("server run success on ", zap.String("zap_log", "zap_log"))

	// TODO: 3. 数据库连接
	global.OE_DB = initialize.Gorm()
	// TODO: 4. 其他初始化
	initialize.OtherInit()

	// TODO: 5. 启动服务
	core.RunServer()

}
