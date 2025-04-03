package global

import (
	"context"
	"fmt"
	"log"
	"pandax/pkg/cache"
	"pandax/pkg/config"
	"pandax/pkg/events"
	"pandax/pkg/tdengine"
	"strings"
	"time"

	"github.com/PandaXGO/PandaKit/logger"
	"github.com/PandaXGO/PandaKit/rediscli"
	"github.com/PandaXGO/PandaKit/starter"
	"github.com/hashicorp/vault/api"
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
const (
	expiryThreshold = 5 * time.Minute // 过期前5分钟触发续租
	// renewalThreshold   = 3 * 24 * time.Hour // 3天有效期阈值
	// shortLeaseDuration = 1 * time.Hour      // 短租约续期时长
)

func init() {
	// 读取配置文件
	// Conf = config.InitConfig(configFile)
	cc := getVault()
	if cc != nil {
		Conf = config.InitVaultConfig(cc["server"].(map[string]interface{}))
		Conf = config.InitVaultConfig(cc)
		// Log.Infof("解析:  %+v", Conf.Server)

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
		// Log.Infof("dbGorm=  %+v", Conf.Postgresql)
		// Log.Infof("dbGorm type=  %v", dbGorm.Type)
		Db = dbGorm.GormInit()
		hrdbGorm := starter.DbGorm{Type: Conf.Server.DbType}
		hrdbGorm.Dsn = Conf.Hrdb.PgDsn()
		hrdbGorm.MaxIdleConns = Conf.Hrdb.MaxIdleConns
		hrdbGorm.MaxOpenConns = Conf.Hrdb.MaxOpenConns
		HrDb = hrdbGorm.GormInit()
		// Log.Infof("解析Db=  %v \n HrDb= %+v", Db, HrDb)

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

// 获取vault服务器中pandax配置信息
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
		log.Fatalf("unable to read config: %v", err)
	}
	// server, _ := secret.Data["server"].(map[string]interface{})
	// taos, _ := server["taos"].(map[string]interface{})
	// log.Print(taos, taos["database"].(string), taos["host"].(string), taos["password"].(string))
	// log.Print(server)
	return secret.Data
}

// 获取vault pandax数据库凭证
func getDatabaseCredential() {
	// ctx := context.Background()
	// 连接到客户端
	config := vault.DefaultConfig()
	config.Address = "http://192.168.0.30:8200"

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to initialize Vault client: %v", err)
	}
	// log.Print("远程配置端连接成功！")
	client.SetToken("hvs.rvdWltAtDl9DFUt5zaoPSYmr")
	// log.Print("Authenticate success！")

	// 获取数据库连接字符串,创建凭证
	db, ee := client.Logical().Read("database/config/localhost_odoo18")
	if ee != nil {
		log.Fatalf("获取数据库连接: %v", ee)
	}
	fmt.Printf("获取数据库连接: %+v\n", getConnectionURL(db.Data))
	conn_url := getConnectionURL(db.Data)

	/*&{RequestID:5d791561-396a-32cf-ccda-6c383450b96a LeaseID: LeaseDuration:0 Renewable:false
	Data:map[allowed_roles:[pAdmin] connection_details:map[backend:database
	connection_url:postgresql://{{username}}:{{password}}@192.168.0.30:5432/odoo18?sslmode=disable
	// max_connection_lifetime:0s max_idle_connections:0 max_open_connections:5 username:lindsay]
	// password_policy: plugin_name:postgresql-database-plugin
	// plugin_version: root_credentials_rotate_statements:[]] Warnings:[] Auth:<nil> WrapInfo:<nil> MountType:database}*/

	// 读取数据库凭证,或者也叫生成数据库凭据，有新的租约和数据库账户信息生成 等同于CLI命令：vault read database/creds/pAdmin
	secret, err := client.Logical().Read("database/creds/pAdmin")
	if err != nil {
		log.Fatalf("获取凭证失败: %v", err)
	}
	fmt.Printf("返回结果secret： %+v\n", secret)
	// 获取动态角色的租约信息
	username := secret.Data["username"].(string)
	password := secret.Data["password"].(string)
	// 替换 {{username}} 和 {{password}}
	conn_url = strings.Replace(conn_url, "{{username}}", username, -1)
	conn_url = strings.Replace(conn_url, "{{password}}", password, -1)
	/*&{RequestID:61778435-0231-4d18-791e-d89055f6cb67 LeaseID:database/creds/pAdmin/fY20y0kTdEdqgwNTAtDGARzW
	LeaseDuration:3600 Renewable:true Data:map[password:lvd-GhGus2r3vVrZEape
	username:v-root-pAdmin-24Js0mFty7VkFx670gci-1743658141]
	Warnings:[] Auth:<nil> WrapInfo:<nil> MountType:database}*/

	// 检查租约状态
	needsRenewal, err := checkLeaseExpiry(client, secret.LeaseID)
	if err != nil {
		fmt.Errorf("检查租约状态失败: %v", err)
	}

	// 根据检查结果处理
	if needsRenewal {
		if secret.Renewable {
			// 续租（续期1小时）
			renewSecret, err := client.Sys().Renew(secret.LeaseID, 3600)
			if err != nil {
				log.Printf("续租失败: %v，将创建新凭证", err)
			}
			log.Printf("租约已成功续期: %+v\n", renewSecret)
		} else {
			// 不可续租则创建新凭证
			log.Println("租约不可续期，创建新凭证")
		}
	}

	// 获取凭证信息,lookup需要指定LeaseID  等同于CLI命令：vault lease lookup database/creds/role2/XJZY2161epFKnHev4GLZkMPu
	// info, e := client.Sys().Lookup("database/creds/pAdmin/fY20y0kTdEdqgwNTAtDGARzW")
	// if e != nil {
	// 	log.Fatalf("查看租约信息失败: %v", e)
	// }
	// fmt.Printf("lookup查看租约信息结果： %+v\n", info)

	/*Data属性中才是真正的返回数据
	&{RequestID:10ebe392-aa73-333c-010f-5b4e25bada9e LeaseID: LeaseDuration:0 Renewable:false
	Data:map[
		expire_time:2025-04-03T14:29:01.3561928+08:00 id:database/creds/pAdmin/fY20y0kTdEdqgwNTAtDGARzW
		issue_time:2025-04-03T13:29:01.3561928+08:00 last_renewal:<nil> renewable:true ttl:3038]
	Warnings:[] Auth:<nil> WrapInfo:<nil> MountType:system}*/

	// 续订租约
	// renew, r := client.Sys().Renew("database/creds/pAdmin/fY20y0kTdEdqgwNTAtDGARzW", 3600)
	// if e != nil {
	// 	log.Fatalf("续订租约失败: %v", r)
	// }
	// fmt.Printf("续订租约结果： %+v\n", renew)
	/*&{RequestID:dc144db1-fc20-a2eb-40fd-c2f06b5874c8 LeaseID:database/creds/pAdmin/fY20y0kTdEdqgwNTAtDGARzW
	LeaseDuration:3600 Renewable:true Data:map[] Warnings:[] Auth:<nil> WrapInfo:<nil> MountType:system}*/
}

// 续租或轮换逻辑
func renewOrRotateLease(leaseID string) (*api.Secret, error) {
	// 1. 查询租约剩余时间
	leaseInfo, err := vaultClient.Sys().Lookup(leaseID)
	if err != nil {
		return nil, err
	}

	// 2. 计算剩余时间
	expireTime, _ := time.Parse(time.RFC3339, leaseInfo.Data["expire_time"].(string))
	remaining := time.Until(expireTime)

	// 3. 判断处理逻辑
	switch {
	case remaining > renewalThreshold:
		// 3天内：续租
		renewed, err := vaultClient.Sys().Renew(leaseID, int(shortLeaseDuration.Seconds()))
		if err != nil {
			return nil, err
		}
		log.Printf("租约续期成功，新过期时间: %v", expireTime.Add(shortLeaseDuration))
		return renewed, nil

	default:
		// 超过3天：撤销旧租约
		if err := vaultClient.Sys().Revoke(leaseID); err != nil {
			log.Printf("租约撤销失败: %v", err)
		}
		log.Println("租约已撤销（超过3天）")
		return nil, nil // 触发创建新租约
	}
}

// 检查租约是否即将过期
func checkLeaseExpiry(client *api.Client, leaseID string) (bool, error) {
	leaseInfo, err := client.Sys().Lookup(leaseID)
	if err != nil {
		return false, fmt.Errorf("查询租约信息失败: %v", err)
	}

	if leaseInfo == nil {
		return true, nil // 租约不存在视为需要更新
	}

	// 计算剩余时间
	expireTime, err := time.Parse(time.RFC3339, leaseInfo.Data["expire_time"].(string))
	if err != nil {
		return false, fmt.Errorf("解析过期时间失败: %v", err)
	}

	remaining := time.Until(expireTime)
	log.Printf("租约剩余时间: %v", remaining)

	// 如果剩余时间小于阈值或已过期
	return remaining < expiryThreshold, nil
}

// 获取vault数据库secret engine中信息
func getConnectionURL(data map[string]interface{}) string {
	return data["connection_details"].(map[string]interface{})["connection_url"].(string)
}

// end
