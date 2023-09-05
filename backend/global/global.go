// global 存放全局变量
package global

import (
	"backend/config"
	"github.com/spf13/viper"
)

var (
	OE_CONFIG config.Configuration
	OE_VIPER  *viper.Viper
)
