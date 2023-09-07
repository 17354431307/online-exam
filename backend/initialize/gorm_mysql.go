package initialize

import (
	"backend/global"
	"backend/initialize/internal"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormMysql 初始化 MySQL 数据库
func GormMysql() *gorm.DB {
	m := global.OE_CONFIG.MySQL
	if m.DbName == "" {
		return nil
	}

	// 创建 mysql.config 实例，其中包含了连接数据库所需的信息，比如 DSN，字符串类型字段的默认长度以及自动根据版本进行初始化等操作
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}

	// 打开数据库连接
	db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular))
	if err != nil {
		panic(err)
		return nil
	}

	// 将引擎设置为我们配置的引擎，并设置每个连接的最大空闲数和最大连接数。
	db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(m.MaxIdleConns)
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)

	fmt.Println("====3-gorm====: gorm link mysql success")
	return db
}
