package middleware

import (
	"database/sql"
	"fmt"
	"pandax/pkg/global"
	"testing"

	// "github.com/PandaXGO/PandaKit/casbin"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestAdapter(t *testing.T) {
	// 创建 GORM 适配器 .getTableInstance()
	t.Logf("\nglobal.Db= %v,\n global.Conf.Casbin.ModelPath= %v\n", global.Db, global.Conf.Casbin.ModelPath)

	adapter, err := gormadapter.NewAdapterByDB(global.Db) //会根据db数据库创建相对应的表，不需要手动建表casbin_rule
	if err != nil {
		t.Fatalf("failed to create gorm adapter: %v", err)
	}

	// 创建 Casbin Enforcer /home/eric/PandaX/resource/rbac_model.conf
	enforcer, err := casbin.NewSyncedEnforcer("/home/eric/PandaX/resource/rbac_model.conf", adapter)
	if err != nil {
		t.Fatalf("failed to create casbin enforcer: %v", err)
	}
	enforcer.LoadPolicy()
	// 添加策略
	// _, err = enforcer.AddPolicy("alice", "data1", "read")
	// if err != nil {
	// 	t.Fatalf("failed to add policy: %v", err)
	// }

	// 检查权限
	ok, err := enforcer.Enforce("admin", "/system/user", "PUT")
	if err != nil {
		t.Fatalf("failed to enforce policy: %v", err)
	}
	if ok {
		fmt.Printf("\nPermission granted\n")
	} else {
		fmt.Println("Permission denied")
	}
}

func TestGromPG(t *testing.T) {
	pg, err := sql.Open("postgres", "postgresql://v-root-pAdmin-dvCNXuxBQrC6lhVZzRWo-1744076990:eA7X-l1alSUo79Qz6Wqi@192.168.0.30:5432/odoo18?sslmode=disable")
	if err != nil {
		fmt.Println("错误：", err)
	}
	ormConfig := &gorm.Config{}
	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: pg,
	}), ormConfig)
	fmt.Println("连接成功：", gormDb)
}
