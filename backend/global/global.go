// global 存放全局变量
package global

import (
	"backend/config"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	OE_CONFIG config.Configuration
	OE_VIPER  *viper.Viper
	OE_Log    *zap.Logger
	OE_DB     *gorm.DB
	OE_REDIS  *redis.Client
)
