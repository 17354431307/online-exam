package internal

import (
	"backend/global"
	"fmt"
	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

func NewWriter(w logger.Writer) *writer {
	return &writer{w}
}

// Printf 格式化打印日志
func (w *writer) Printf(message string, data ...any) {
	var logZap bool
	switch global.OE_CONFIG.App.DbType {
	case "mysql":
		logZap = global.OE_CONFIG.MySQL.LogZap
	}

	if logZap {
		global.OE_Log.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
