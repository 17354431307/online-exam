package initialize_test

import (
	"backend/core"
	"backend/global"
	"backend/initialize"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGormPgSql(t *testing.T) {
	global.OE_VIPER = core.InitViper("../etc/config.yaml")
	assert.NotNil(t, global.OE_CONFIG.PGSQL)
	assert.Equal(t, global.OE_CONFIG.PGSQL.Port, "5432")
	initialize.GormPgSql()
}

func TestGormMysql(t *testing.T) {
	global.OE_VIPER = core.InitViper("../etc/config.yaml")
	assert.NotNil(t, global.OE_CONFIG.MySQL)
	assert.Equal(t, global.OE_CONFIG.MySQL.Port, "3306")
	initialize.GormMysql()
}
