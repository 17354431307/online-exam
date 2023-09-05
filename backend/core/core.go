package core

import (
	"backend/core/internal"
	"backend/global"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

// InitViper 优先级: 命令行 > 环境变量 > 默认值
func InitViper(path ...string) *viper.Viper {
	var config string

	if len(path) == 0 {

		flag.StringVar(&config, "c", "", "请选择配置文件.")
		flag.Parse()

		if config == "" {
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDefaultFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigDefaultFile)
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigReleaseFile)
				case gin.TestMode:
					config = internal.ConfigTestFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigTestFile)
				}

			} else {
				config = configEnv
				fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", internal.ConfigEnv, config)
			}
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", config)
	}

	vip := viper.New()
	vip.SetConfigFile(config)
	vip.SetConfigType("yaml")
	err := vip.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败, err: %s \n", err))
	}

	vip.WatchConfig()

	vip.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件已变更:", e.Name)
		if err = vip.Unmarshal(&global.OE_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err = vip.Unmarshal(&global.OE_CONFIG); err != nil {
		fmt.Println(err)
	}

	fmt.Println("====1-viper====: viper init config success")

	return vip
}
