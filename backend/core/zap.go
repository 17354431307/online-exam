package core

import (
	"backend/core/internal"
	"backend/global"
	"backend/utils"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// InitializeZap Zap 获取 zap.Logger
func InitializeZap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.OE_CONFIG.Zap.Director); !ok { // 判断是否有 Director 文件夹
		fmt.Printf("create %v directory\n", global.OE_CONFIG.Zap.Director)
		_ = os.Mkdir(global.OE_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.OE_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	fmt.Println("====2-zap====: zap log init success")
	return logger
}
