package casbin

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func initEnforcer(t *testing.T) (*casbin.Enforcer, error) {
	adapter, err := gormadapter.NewAdapter(
		"mysql",
		"root:123456@tcp(127.0.0.1:3306)/casbin_demo_db?charset=utf8mb4&parseTime=true&loc=Local",
		true,
	)
	assert.NoError(t, err)

	enforcer, err := casbin.NewEnforcer("./model.conf", adapter)
	assert.NoError(t, err)
	return enforcer, nil
}

func Test_Add_Policy(t *testing.T) {
	enforcer, err := initEnforcer(t)
	assert.NoError(t, err)

	// AddPolicy 添加单个 policy, 如果不存在则添加并返回 true；存在则不添加并返回 false
	/*ok, err := enforcer.AddPolicy("alice", "data1", "read")
	assert.NoError(t, err)
	if ok {
		t.Log("添加策略成功")
	} else {
		t.Log("策略已经存在")
	}*/

	// AddPolicies 添加多个 policys， 都不存在则添加并返回true；如果有一个存在则添加失败并返回 false
	/*rules := [][]string{
		[]string{"bob", "data2", "write"},
		[]string{"data2_admin", "data2", "read"},
		[]string{"data2_admin", "data2", "write"},
	}
	ok, err := enforcer.AddPolicies(rules)
	assert.NoError(t, err)
	if ok {
		t.Log("批量添加策略成功")
	} else {
		t.Log("批量添加策略失败")
	}*/

	// AddPoliciesEx() 添加多个 policys，和 AddPolicies 是一样的行为，但是有点区别，当添加的策略中如果有已经
	// 策略的话，不仅仅返回 false，还有返回 数据库主键冲突错误 error
	/*rules := [][]string{
		[]string{"data2_admin", "data2", "read"},
		[]string{"bob", "data2", "write"},
		[]string{"data2_admin", "data2", "write"},
	}
	ok, err := enforcer.AddPoliciesEx(rules)
	assert.NoError(t, err)
	if ok {
		t.Log("批量添加策略成功")
	} else {
		t.Log("批量添加策略失败")
	}*/

	// AddNamedPolicy 添加单个策略的底层方法，和 AddPolicy 没什么区别，只不过第一个形参变为了 ptype, 策略的类型
	// 策略的类型目前已经知道的有 p, g; p 代表着一条策略，g 代表着角色
	/*ok, err := enforcer.AddNamedPolicy("p", "eve", "data3,", "read")
	assert.NoError(t, err)
	if ok {
		t.Log("添加策略成功")
	} else {
		t.Log("策略已经存在")
	}*/

	// AddNamedPolicies 批量添加策略的底层方法，和 AddPolicys 没什么区别，只不过第一个形参变成了 ptype
	/*rules := [][]string{
		[]string{"data2_admin", "data2", "read"},
		[]string{"bob", "data2", "write"},
		[]string{"data2_admin", "data2", "write"},
	}
	ok, err := enforcer.AddNamedPolicies("p", rules)
	assert.NoError(t, err)
	if ok {
		t.Log("批量添加策略成功")
	} else {
		t.Log("批量添加策略失败")
	}*/

	// AddNamedPoliciesEx 是 AddPoliciesEx 的底层方法，区别不大，只不过第一个形参变成了 ptype
	/*rules := [][]string{
		[]string{"data2_admin", "data2", "read"},
		[]string{"bob", "data2", "write"},
		[]string{"data2_admin", "data2", "write"},
	}
	ok, err := enforcer.AddNamedPoliciesEx("p", rules)
	assert.NoError(t, err)
	if ok {
		t.Log("批量添加策略成功")
	} else {
		t.Log("批量添加策略失败")
	}*/

	// SelfAddPolicy 会将授权规则添加到当前已命名的策略中，并禁用自动通知观察者功能。
	// 其他的行为就和 addPolicy 差不多，但是这个非常便捷，可以添加 p(策略类型) 和 g(角色类型)
	/*ok, err := enforcer.SelfAddPolicy("g", "g", []string{"alice", "data1_admin"})
	assert.NoError(t, err)
	if ok {
		t.Log("添加角色成功")
	} else {
		t.Log("添加角色失败")
	}*/

	// SelfAddPolicies 和 SelfAddPolicy 差不多的行为，只不过是批量的，全都不存在则添加成功返回 true，有一个存在则添加失败，返回 false
	/*rules := [][]string{
		[]string{"data2_admin", "data2", "read"},
		[]string{"bob", "data2", "write"},
		[]string{"data2_admin", "data2", "write"},
	}
	ok, err := enforcer.SelfAddPolicies("p", "p", rules)
	assert.NoError(t, err)
	if ok {
		t.Log("批量添加策略成功")
	} else {
		t.Log("批量添加策略失败")
	}
	ok, err = enforcer.SelfAddPolicies("g", "g", [][]string{[]string{"bob", "data2_admin"}, {"alice", "data2_admin"}})
	assert.NoError(t, err)
	if ok {
		t.Log("批量添加角色成功")
	} else {
		t.Log("批量添加角色失败")
	}*/

	// SelfAddPoliciesEx 和 SelfAddPolicies 差不多的行为，只不过添加失败的是否会返回数据库错误
	/*rules := [][]string{
		[]string{"data2_admin", "data2", "read"},
		[]string{"bob", "data2", "write"},
		[]string{"data2_admin", "data2", "write"},
	}
	ok, err := enforcer.SelfAddPoliciesEx("p", "p", rules)
	assert.NoError(t, err)
	if ok {
		t.Log("批量添加策略成功")
	} else {
		t.Log("批量添加策略失败")
	}
	ok, err = enforcer.SelfAddPoliciesEx("g", "g", [][]string{[]string{"bob", "data2_admin"}, {"alice", "data2_admin"}})
	assert.NoError(t, err)
	if ok {
		t.Log("批量添加角色成功")
	} else {
		t.Log("批量添加角色失败")
	}*/

	// AddGroupingPolicy 向当前策略添加角色继承规则。 如果规则已经存在，函数返回false，并且不会添加规则。 否则，函数通过添加新规则并返回true。
	/*	ok, err := enforcer.AddGroupingPolicy("alice", "data1_admin")
		assert.NoError(t, err)
		if ok {
			t.Log("添加角色成功")
		} else {
			t.Log("添加角色失败")
		}*/

	// AddGroupingPolicies 和 AddPolicies 行为差不多，只不过操作对象变成了角色
	/*rules := [][]string{
		[]string{"bob", "data1_admin1"},
		[]string{"alice", "data2_damin2"},
	}
	ok, err := enforcer.AddGroupingPolicies(rules)
	assert.NoError(t, err)
	if ok {
		t.Log("批量添加角色成功")
	} else {
		t.Log("批量添加角色失败")
	}*/

	// AddGroupingPoliciesEx 和 AddPoliciesEx 行为差不多，只不过操作对象变成了角色
	/*rules := [][]string{
		[]string{"bob", "data1_admin1"},
		[]string{"alice", "data2_damin2"},
	}
	ok, err := enforcer.AddGroupingPoliciesEx(rules)
	assert.NoError(t, err)
	if ok {
		t.Log("批量添加策略成功")
	} else {
		t.Log("批量添加策略失败")
	}*/

	// AddNamedGroupingPolicy, AddGroupingPolicies, AddNamedGroupingPoliciesEx 也是同样的行为，不做测试了

	_ = enforcer
}

func Test_Get_Policy(t *testing.T) {
	enforcer, err := initEnforcer(t)
	assert.NoError(t, err)

	// 只测试几个经典的

	// GetPolicy 获取所有的策略
	/*	policys := enforcer.GetPolicy()
		for _, policy := range policys {
			t.Log(policy)
		}*/

	// GetFilteredPolicy 获取策略中的所有授权规则，可以指定字段筛选器。
	// findIndex 表示第几个个字段(sub, obj, act), fieldValues 表示从指定的字段中匹配对应的内容
	/*	policys := enforcer.GetFilteredPolicy(1, "data2")
		for _, policy := range policys {
			t.Log(policy)
		}*/

	// GetNamedPolicy 获取命名策略中的所有授权规则。其实就是 GetPolicy 的底层方法啦
	/*policys := enforcer.GetNamedPolicy("p")
	for _, policy := range policys {
		t.Log(policy)
	}*/

	// GetGroupXXX 也有相同的一套

	// HasPolicy 确定是否存在授权规则。
	/*ok := enforcer.HasPolicy("data2_admin", "data2", "read")
	if ok {
		t.Log("存在该授权规则")
	} else {
		t.Log("不存在该授权规则")
	}*/

	// HasNamedPolicy 确定是否存在命名授权规则。
	/*ok := enforcer.HasNamedPolicy("p", "data2_admin", "data2", "read")
	if ok {
		t.Log("存在该授权规则")
	} else {
		t.Log("不存在该授权规则")
	}*/

	// HasGroupXXX 也有相同的一套

	_ = enforcer
}

func Test_Remove_Policy(t *testing.T) {
	enforcer, err := initEnforcer(t)
	assert.NoError(t, err)

	// 删除 policy
	// RemovePolicy 从当前策略中删除授权规则
	/*ok, err := enforcer.RemovePolicy("tom", "data1", "write")
	assert.NoError(t, err)
	if ok {
		t.Log("删除策略成功")
	} else {
		t.Log("删除策略失败")
	}*/

	// RemovePolicies 从当前策略中删除授权规则。
	// 该操作本质上是原子的 因此，如果授权规则由不符合现行政策的规则组成， 函数返回 false ，当前政策中没有任何政策规则被删除。
	// 如果所有授权规则都符合政策规则，则函数返回true，每项政策规则都从现行政策中删除。
	/*rules := [][]string{
		[]string{
			"tom", "data1", "write",
		},
		[]string{
			"tom", "data1", "read",
		},
	}
	ok, err := enforcer.RemovePolicies(rules)
	assert.NoError(t, err)
	if ok {
		t.Log("批量删除策略成功")
	} else {
		t.Log("批量删除策略失败")
	}*/

	// RemoveFilteredPolicy 移除当前策略中的授权规则，可以指定字段筛选器。 RemovePolicy 从当前策略中删除授权规则。
	/*ok, err := enforcer.RemoveFilteredPolicy(0, "tom", "data1", "write")
	assert.NoError(t, err)
	if ok {
		t.Log("删除策略成功")
	} else {
		t.Log("删除策略失败")
	}*/

	// RemoveNamedPolicy()
	// RemoveNamedPolicies()
	// RemoveFilteredNamedPolicy()
	// 这三个的原理和 AddXXX 的原理类似，都是各自的底层方法

	// 删除 role
	// RemoveGroupingPolicy 从当前策略中删除角色继承规则。
	/*ok, err := enforcer.RemoveGroupingPolicy("tom", "data2_admin")
	assert.NoError(t, err)
	if ok {
		t.Log("删除角色成功")
	} else {
		t.Log("删除角色失败")
	}*/

	// RemoveGroupingPolicy 从当前策略中删除角色继承规则。
	// 该操作本质上是原子的 因此，如果授权规则由不符合现行政策的规则组成， 函数返回 false ，当前政策中没有任何政策规则被删除。
	// 如果所有授权规则都符合政策规则，则函数返回true，每项政策规则都从现行政策中删除。
	/*rules := [][]string{
		[]string{
			"tom", "data2_admin",
		},
		[]string{
			"tom", "data3_admin",
		},
	}
	ok, err := enforcer.RemovePolicies(rules)
	assert.NoError(t, err)
	if ok {
		t.Log("批量删除角色成功")
	} else {
		t.Log("批量删除角色失败")
	}*/

	// RemoveFilteredGroupingPolicy 从当前策略中移除角色继承规则，可以指定字段筛选器。
	/*ok, err := enforcer.RemoveFilteredGroupingPolicy(0, "tom", "data2_admin")
	assert.NoError(t, err)
	if ok {
		t.Log("删除角色成功")
	} else {
		t.Log("删除角色失败")
	}*/

	// RemoveNamedGroupingPolicy()
	// RemoveNamedGroupingPolicies()
	// RemoveFilteredNamedGroupingPolicy()
	// 对应的取得 Named 名字方法的底层方法

	//67	g	tom	data2_admin
	_ = enforcer
}

func Test_Update_Policy(t *testing.T) {
	enforcer, err := initEnforcer(t)
	assert.NoError(t, err)

	// UpdatePolicy 把旧的策略更新到新的策略
	/*ok, err := enforcer.UpdatePolicy([]string{"tom", "data4", "read"}, []string{"tom1", "data5", "write"})
	assert.NoError(t, err)
	if ok {
		t.Log("修改 policy 成功")
	} else {
		t.Log("修改 policy 失败")
	}*/

	// UpdatePolicies 将所有的旧政策更新到新政策
	/*ok, err := enforcer.UpdatePolicies([][]string{
		[]string{"tom1", "data5", "write"},
		[]string{"tom2", "data5", "write"},
	}, [][]string{
		[]string{"tom1", "data4", "read"},
		[]string{"tom2", "data4", "read"},
	})
	assert.NoError(t, err)
	if ok {
		t.Log("批量修改 policy 成功")
	} else {
		t.Log("批量修改 policy 失败")
	}*/

	// UpdateGroupingPolicy 在 g段更新oldRule到newRule
	/*ok, err := enforcer.UpdateGroupingPolicy([]string{
		"alice", "data1_admin",
	}, []string{
		"alice", "data2_admin",
	})
	assert.NoError(t, err)
	if ok {
		t.Log("修改 role 成功")
	} else {
		t.Log("修改 role 失败")
	}*/

	ok, err := enforcer.UpdateGroupingPolicies(
		[][]string{
			[]string{
				"alice", "data2_admin",
			},
			[]string{
				"tom", "data2_admin",
			},
		},
		[][]string{
			[]string{
				"alice", "data3_admin",
			},
			[]string{
				"tom", "data3_admin",
			},
		},
	)
	assert.NoError(t, err)
	if ok {
		t.Log("批量修改 role 成功")
	} else {
		t.Log("批量修改 role 失败")
	}
}
