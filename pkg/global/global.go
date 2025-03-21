package global

import (
	"context"
	"log"
	"pandax/pkg/cache"
	"pandax/pkg/config"
	"pandax/pkg/events"
	"pandax/pkg/tdengine"

	"github.com/PandaXGO/PandaKit/logger"
	"github.com/PandaXGO/PandaKit/rediscli"
	"github.com/PandaXGO/PandaKit/starter"
	vault "github.com/hashicorp/vault/api"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Log  *logrus.Logger // 日志
	Db   *gorm.DB       // gorm
	HrDb *gorm.DB       // hrp数据库 2025.2.5 add
	TdDb *tdengine.TdEngine
	Conf *config.Config
)
var EventEmitter = events.EventEmitter{}

// 2025-2-10 修改 本应在main.go中加载
func init() {
	// 读取配置文件
	// Conf = config.InitConfig(configFile)
	cc := getVault()
	if cc != nil {
		Conf = config.InitVaultConfig(cc["server"].(map[string]interface{}))
		Conf = config.InitVaultConfig(cc)
		// Log.Infof("解析:  %v", Conf)

		Log = logger.InitLog(Conf.Log.File.GetFilename(), Conf.Log.Level)
		dbGorm := starter.DbGorm{Type: Conf.Server.DbType}
		if Conf.Server.DbType == "mysql" {
			dbGorm.Dsn = Conf.Mysql.Dsn()
			dbGorm.MaxIdleConns = Conf.Mysql.MaxIdleConns
			dbGorm.MaxOpenConns = Conf.Mysql.MaxOpenConns
		} else {
			dbGorm.Dsn = Conf.Postgresql.PgDsn()
			dbGorm.MaxIdleConns = Conf.Postgresql.MaxIdleConns
			dbGorm.MaxOpenConns = Conf.Postgresql.MaxOpenConns
		}
		// Log.Infof("dbGorm=  %v", Conf.Postgresql)
		// Log.Infof("dbGorm type=  %v", dbGorm.Type)
		Db = dbGorm.GormInit()
		hrdbGorm := starter.DbGorm{Type: Conf.Server.DbType}
		hrdbGorm.Dsn = Conf.Hrdb.PgDsn()
		hrdbGorm.MaxIdleConns = Conf.Hrdb.MaxIdleConns
		hrdbGorm.MaxOpenConns = Conf.Hrdb.MaxOpenConns
		HrDb = hrdbGorm.GormInit()
		// Log.Infof("解析Db=  %v \n HrDb= %v", Db, HrDb)

		client, err := rediscli.NewRedisClient(Conf.Redis.Host, Conf.Redis.Password, Conf.Redis.Port, Conf.Redis.Db)
		if err != nil {
			Log.Panic("Redis连接错误")
		} else {
			Log.Info("Redis连接成功")
		}
		cache.RedisDb = client
		tDengine, err := tdengine.InitTd(Conf.Taos.Username, Conf.Taos.Password, Conf.Taos.Host, Conf.Taos.Database)
		if err != nil {
			Log.Panic("Tdengine连接错误")
		} else {
			Log.Info("Tdengine连接成功")
		}
		TdDb = tDengine
	} else {
		log.Panic("请检查远程配置config！")
	}
}

func getVault() map[string]interface{} {
	ctx := context.Background()
	// 连接到客户端
	config := vault.DefaultConfig()
	config.Address = "http://192.168.0.30:8200"

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}
	log.Print("远程配置端连接成功！")

	// Authenticate 这里以超级管理员进行token认证
	// WARNING: This quickstart uses the root token for our Vault dev server.
	// Don't do this in production!
	client.SetToken("hvs.rvdWltAtDl9DFUt5zaoPSYmr")
	log.Print("Authenticate success！")
	secret, err := client.KVv2("pandax05eric").Get(ctx, "pandax")
	if err != nil {
		log.Fatalf("unable to read secret: %v", err)
	}
	// server, _ := secret.Data["server"].(map[string]interface{})
	// taos, _ := server["taos"].(map[string]interface{})
	// log.Print(taos, taos["database"].(string), taos["host"].(string), taos["password"].(string))
	// log.Print(server)
	return secret.Data
}

// end
