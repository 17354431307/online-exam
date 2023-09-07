package config

import "fmt"

type MySQL struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`                               // 服务器地址
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                               // MySQL 访问端口
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                         // 高级配置
	DbName       string `mapstructure:"db_name" json:"db_name" yaml:"db_name"`                      // 数据库名称
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                   // 用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                   // 密码
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 全局表前缀，单独定义 TableName 不生效
	Singular     bool   `mapstructure:"singular" json:"singular" yaml:"singular"`                   // 是否开启全局禁用复数，true 表示不开启
	Engine       string `mapstructure:"engine" json:"engine" yaml:"engine"`                         // 引擎，默认 InnoDB
	MaxIdleConns int    `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"` // 最大空闲连接数
	MaxOpenConns int    `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"` // 最大连接数
	LogMode      string `mapstructure:"log_mode" json:"log_mode" yaml:"log_mode"`                   // 日志级别
	LogZap       bool   `mapstructure:"log_zap" json:"log_zap" yaml:"log_zap"`                      // 是否通过 zap 库写日志文件
}

func (m *MySQL) Dsn() string {
	// username:password@protocol(address)/dbname?param=value
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", m.Username, m.Password, m.Host, m.Port, m.DbName, m.Config)
}

func (m *MySQL) GetLogMode() string {
	return m.LogMode
}
