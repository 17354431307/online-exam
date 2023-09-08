package initialize

import (
	"backend/global"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func Redis() {

	client := redis.NewClient(&redis.Options{
		Addr:     global.OE_CONFIG.Redis.Addr,
		Password: global.OE_CONFIG.Redis.Password, // no password set
		DB:       global.OE_CONFIG.Redis.DB,       // use default DB
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.OE_Log.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		fmt.Println("====4-redis====: redis init success")
		global.OE_Log.Info("redis connect ping response:", zap.String("pong", pong))
		global.OE_REDIS = client
	}
}
