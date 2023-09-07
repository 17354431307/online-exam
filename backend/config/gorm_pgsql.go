package config

import "fmt"

type PGSQL struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`                               // 服务器地址
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                               // PostgresSQL 访问端口
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                         // 高级配置
	DbName       string `mapstructure:"db_name" json:"db_name" yaml:"db_name"`                      // 数据库名称
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                   // 用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                   // 密码
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 全局表前缀，单独定义 TableName 不生效
	Singular     bool   `mapstructure:"singular" json:"singular" yaml:"singular"`                   // 是否开启全局禁用复数，true 表示不开启
	Engine       string `mapstructure:"engine" json:"engine" yaml:"engine"`                         // 引擎，默认 InnoDB
	MaxIdleConns int    `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"` // 最大空闲连接数
	MaxOpenConns int    `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"` // 最大连接数
	LogMode      string `mapstructure:"log_mode" json:"log_mode" yaml:"log_mode"`                   // 是否开启Gorm全局日志
	LogZap       bool   `mapstructure:"log_zap" json:"log_zap" yaml:"log_zap"`                      // 是否通过 zap 库写日志文件
}

// Dsn 基于配置文件或者 dsn
func (p *PGSQL) Dsn() string {
	// host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s %s", p.Host, p.Username, p.Password, p.DbName, p.Port, p.Config)
}

// LinkDsb 根据 dbname 生成 dsn
func (p *PGSQL) LinkDsb(dbName string) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s %s", p.Host, p.Username, p.Password, dbName, p.Port, p.Config)
}

func (p *PGSQL) GetLogMode() string {
	return p.LogMode
}
