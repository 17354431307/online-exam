package initialize

import (
	"backend/global"
	"backend/model/business"
	"backend/model/example"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

// Gorm 初始化数据库并产生数据库全局变量
func Gorm() *gorm.DB {
	switch global.OE_CONFIG.App.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	default:
		return GormMysql()
	}
}

// RegisterTables 注册数据库表专用
func RegisterTables() {
	db := global.OE_DB
	err := db.AutoMigrate(
		example.ExaCusmoter{},
		business.User{},
	)

	if err != nil {
		global.OE_Log.Error("register table failed", zap.Error(err))
		os.Exit(1)
	}
	global.OE_Log.Info("register table success")
}
