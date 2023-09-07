//go:build e2e

package core

import (
	"backend/global"
	"go.uber.org/zap"
	"testing"
)

func TestInitializeZap(t *testing.T) {
	global.OE_VIPER = InitViper("../etc/config.yaml")

	global.OE_Log = InitializeZap()

	zap.ReplaceGlobals(global.OE_Log)

	global.OE_Log.Info("server run success on ", zap.String("zap_log", "zap_log"))
}
