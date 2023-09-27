package acl_test

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CasbinWithGorm(t *testing.T) {
	// 初始化 Gorm 适配器并在 Casbin 执行器中使用：
	// 适配器将使用名为 "casbin_demo_db" 的 MySQL 数据库。
	// 如果它不存在，适配器将自动创建它。
	// 您还可以将已经存在的 gorm 实例与 gormadapter 一起使用。NewAdapterByDB(gormInstance)
	//a, err := gormadapter.NewAdapter(
	//	"mysql",
	//	"root:123456@tcp(127.0.0.1:3306)/casbin_demo_db?charset=utf8mb4&parseTime=true&loc=Local",
	//)

	// 或者您可以使用现有的数据库 "abc", 如下所示:
	// 适配器将使用名为 "casbin_rule" 的表.
	// 如果不存在，适配器将自动创建它.
	// a := gormaapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)
	a, err := gormadapter.NewAdapter(
		"mysql",
		"root:123456@tcp(127.0.0.1:3306)/casbin_demo_db?charset=utf8mb4&parseTime=true&loc=Local",
		true,
	)
	assert.NoError(t, err)

	// 初始化 cabin 执行器
	e, err := casbin.NewEnforcer("./model.conf", a)
	assert.NoError(t, err)

	// 从数据库 加载 policy。
	err = e.LoadPolicy()
	assert.NoError(t, err)

	// 检查 permission
	ok, err := e.Enforce("alice", "data1", "read")
	assert.NoError(t, err)
	assert.Equal(t, false, ok)

	// 添加 policy
	// e.AddPolicy(...)
	ok, err = e.AddPolicy("he", "api/v1/nb", "GET")
	assert.NoError(t, err)
	assert.True(t, ok)

	// 删除 policy
	//e.RemovePolicy(...)
	ok, err = e.RemovePolicy("he", "api/v1/nb", "GET")
	assert.NoError(t, err)
	assert.True(t, ok)

	// 将 policy 保存回数据库。
	policys := e.GetPolicy()
	for _, policy := range policys {
		t.Log(policy)
		if policy[0] == "alice" {
			policy[2] = "POST"
		}
	}
	err = e.SavePolicy()
	assert.NoError(t, err)
}
