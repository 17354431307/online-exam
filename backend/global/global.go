// global 存放全局变量
package global

import (
	"backend/config"
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

/*
	singleflight.Group{} 是 Go 语言的并发控制机制，它用于在并发环境中对相同的操作进行合并和共享结果，以降低重复操作的成本
	在并发编程中，有时候多个并发任务可能会同时发起相同的请求或操作，这样就会导致重复的计算或查询，浪费计算资源或网络带宽。
	singleflight.Group 通过合并相同的请求，只执行一次实际的操作，然后将结果返回给所有并发的请求者
*/

var (
	OE_CONFIG config.Configuration
	OE_VIPER  *viper.Viper
	OE_Log    *zap.Logger
	OE_DB     *gorm.DB
	OE_REDIS  *redis.Client

	OE_Concurrency_Control = &singleflight.Group{}

	BlackCahe local_cache.Cache
)
