package system

import (
	"backend/global"
	"backend/model/system/request"
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"strconv"
	"sync"
)

type CasbinService struct {
}

var CasbinServiceApp = new(CasbinService)

// UpdateCasbin 更新 Casbin 权限
func (c *CasbinService) UpdateCasbin(AuthorityID uint, casbinInfos []request.CasbinInfo) error {
	authorityId := strconv.Itoa(int(AuthorityID))
	c.ClearCasbin(0, authorityId)
	rules := [][]string{}

	// 做权限去重处理
	deduplicateMap := make(map[string]bool)
	for _, v := range casbinInfos {
		key := authorityId + v.Path + v.Method
		if _, ok := deduplicateMap[key]; !ok {
			deduplicateMap[key] = true
			rules = append(rules, []string{authorityId, v.Path, v.Method})
		}
	}
	e := c.Casbin()
	success, _ := e.AddPolicies(rules)
	if !success {
		return errors.New("存在相同api, 添加失败, 请联系管理员")
	}
	return nil
}

// UpdateCasbinApi API更新随动
func (c *CasbinService) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := global.OE_DB.Model(&gormadapter.CasbinRule{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]any{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	if err != nil {
		return err
	}

	e := c.Casbin()
	return e.LoadPolicy()
}

// ClearCasbin 清除匹配的权限
func (c *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := c.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

// GetPolicyPathByAuthority 获取权限列表
func (c *CasbinService) GetPolicyPathByAuthority(AuthorityID uint) (pathMaps []request.CasbinInfo) {
	e := c.Casbin()
	authorityId := strconv.Itoa(int(AuthorityID))
	list := e.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})

	}
	return pathMaps
}

var (
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
	once                 sync.Once
)

// Casbin 持久化到数据库  引入自定义规则
func (c *CasbinService) Casbin() *casbin.SyncedCachedEnforcer {
	once.Do(func() {
		a, err := gormadapter.NewAdapterByDB(global.OE_DB)
		if err != nil {
			// TODO 这里可能有错误
			global.OE_Log.Error("适配数据库失败请检查casbin表是否为InnoDB引擎!", zap.Error(err))
			//zap.L().Error("适配数据库失败请检查casbin表是否为InnoDB引擎!", zap.Error(err))
			return
		}

		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
	
		[role_definition]
		g = _, _

		[policy_effect]
		e = some(where (p.eft == allow))

		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`

		m, err := model.NewModelFromString(text)
		if err != nil {
			// TODO 这里可能有错误
			global.OE_Log.Error("字符串加载模型失败!", zap.Error(err))
			return
		}

		syncedCachedEnforcer, _ = casbin.NewSyncedCachedEnforcer(m, a)
		syncedCachedEnforcer.SetExpireTime(60 * 60)
		_ = syncedCachedEnforcer.LoadPolicy()

	})
	return syncedCachedEnforcer
}
