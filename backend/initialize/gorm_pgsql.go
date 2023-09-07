package initialize

import (
	"backend/global"
	"backend/initialize/internal"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GormPgSql() *gorm.DB {
	p := global.OE_CONFIG.PGSQL

	if p.DbName == "" {
		return nil
	}

	pgsqlConfig := postgres.Config{
		DSN:                  p.Dsn(), // data source name
		PreferSimpleProtocol: false,   // 启用 prepare statement 缓存
	}

	db, err := gorm.Open(postgres.New(pgsqlConfig), internal.Gorm.Config(p.Prefix, p.Singular))
	if err != nil {
		panic(err)
	}

	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(p.MaxIdleConns)
	sqlDb.SetMaxOpenConns(p.MaxOpenConns)
	fmt.Println("====3-gorm====: gorm link PostgreSQL success")
	return db
}
