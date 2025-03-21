package middleware

import (
	"pandax/pkg/global"

	"github.com/PandaXGO/PandaKit/biz"
	// "github.com/PandaXGO/PandaKit/casbin"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/PandaXGO/PandaKit/token"

	"github.com/golang-jwt/jwt/v5"

	casbint "github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/lib/pq"
)

// 鉴权中间件————每次请求路由时都需要执行判断 后期如果要修改odoo用户权限应该也在这里实现
func PermissionHandler(rc *restfulx.ReqCtx) error {
	permission := rc.RequiredPermission
	// 如果需要的权限信息不为空，并且不需要token，则不返回错误，继续后续逻辑
	if permission != nil && !permission.NeedToken {
		return nil
	}
	tokenStr := rc.Request.Request.Header.Get("X-TOKEN")
	// header不存在则从查询参数token中获取
	if tokenStr == "" {
		tokenStr = rc.Request.QueryParameter("token")
	}
	if tokenStr == "" {
		return biz.PermissionErr
	}
	// global.Log.Infof("已获取token=  %v", tokenStr)
	j := token.NewJWT("", []byte(global.Conf.Jwt.Key), jwt.SigningMethodHS256)
	loginAccount, err := j.ParseToken(tokenStr)
	if err != nil || loginAccount == nil {
		return biz.PermissionErr
	}
	// global.Log.Infof("获取登录账号=  %v", rc.ResData)
	// global.Log.Infof("获取登录账号sub=  %v", loginAccount.RoleKey)
	rc.LoginAccount = loginAccount

	if !permission.NeedCasbin {
		return nil
	}
	// global.Log.Infof("登录账号同步给restful api=  %v", rc.Request.Request)
	// ca := casbin.CasbinService{ModelPath: global.Conf.Casbin.ModelPath}
	// e := ca.GetCasbinEnforcer()     //BUG!!!!!!!!!!!

	adapter, err := gormadapter.NewAdapterByDB(global.Db) //会根据db数据库创建相对应的表，不需要手动建表casbin_rule
	if err != nil {
		global.Log.Fatalf("failed to create gorm adapter: %v", err)
	}

	// 创建 Casbin Enforcer /home/eric/PandaX/resource/rbac_model.conf
	e, err := casbint.NewSyncedEnforcer("/home/eric/PandaX/resource/rbac_model.conf", adapter)
	// 加载权限
	err = e.LoadPolicy()
	if err != nil {
		global.Log.Fatalf("Failed to load policy: %v", err)
	}

	success, err := e.Enforce(loginAccount.RoleKey, rc.Request.Request.URL.Path, rc.Request.Request.Method)
	// global.Log.Infof("判断策略中是否存在=  %v,  %v", success, err) //output：false,  <nil>
	if !success || err != nil {
		return biz.CasbinErr
	}

	return nil
}
