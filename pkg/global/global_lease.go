package global

import (
	"context"
	"encoding/json"
	"log"
	"pandax/pkg/cache"
	"strings"
	"time"

	"github.com/PandaXGO/PandaKit/starter"
	"github.com/google/uuid"
	"github.com/hashicorp/vault/api"
	"github.com/kakuilan/kgo"
)

const (
	RedisLeaseKey      = "vault:database:lease"  // Redis存储键名
	redisLeaseKey2     = "vault:database:lease2" // Redis存储键名
	redisPwdKey        = "vault:database:"       // Redis存储键名
	renewalThreshold   = 3 * 24 * time.Hour      // 3天有效期阈值
	shortLeaseDuration = 184 * time.Minute       // 短租约续期时长
	maxRetries         = 5                       //数据库最大连接数
	role               = "pAdmin"                //vault管理角色，同时管理以下两个数据库连接
	hrp                = "localhost_odoo18"      //HR连接
	iot                = "Pandax05"              //IOT连接
)

type RedisLease struct {
	LeaseID   string    `redis:"lease_id"`
	Uuid      string    `redis:"uuid"`
	CreatedAt time.Time `redis:"created_at"`
}

// 初始化连接,最多尝试五次
func DatabaseManage() {
	// for i := 0; i < maxRetries; i++ {
	// 	if err := InitializeConnections(); err != nil {
	// 		log.Printf("数据库初始化失败（第%d次）: %v", i+1, err)
	// 		// Log.Infof("数据库连接失败,初始化第 %d 次", i)
	// 		time.Sleep(time.Duration(i) * time.Minute)
	// 		continue
	// 	}
	// 	return
	// }

	err := InitializeConnections()
	if err != nil {
		log.Printf("初始化数据库 %s 失败: %v", hrp, err)
	}

	Log.Error("数据库连接初始化失败")
}

/*获取数据库连接以及状态处理*/
func InitializeConnections() error {
	// 1. 获取凭证
	cred, err := manageDatabaseCredential() //这里获取的唯有首次read创建动态凭证的返回值,无论如何都有账号
	if err != nil {
		return err
	}

	// 获取数据库连接信息
	// tt := time.Now()
	db, ee := vaultClient.Logical().Read("database/config/" + hrp)
	if ee != nil {
		// log.Fatalf("获取数据库连接: %v", ee)
		return ee
	}
	// 获取动态角色的租约信息
	username := cred.Data["username"].(string)
	password := cred.Data["password"].(string)
	conn_url := getConnectionURL(db.Data, username, password)
	maxOpen, _ := db.Data["connection_details"].(map[string]interface{})["max_open_connections"].(json.Number).Int64()
	maxIdel, _ := db.Data["connection_details"].(map[string]interface{})["max_idle_connections"].(json.Number).Int64()
	hrdbGorm := starter.DbGorm{Type: Conf.Server.DbType}
	hrdbGorm.Dsn = conn_url
	hrdbGorm.MaxIdleConns = int(maxIdel)
	hrdbGorm.MaxOpenConns = int(maxOpen)
	HrDb = hrdbGorm.GormInit()
	// log.Printf("解析odoo数据库连接HrDb= %+v\n", HrDb)
	// pandax数据库连接
	DB2, ee := vaultClient.Logical().Read("database/config/" + iot)
	if ee != nil {
		// log.Fatalf("获取数据库连接: %v", ee)
		return ee
	}
	psd := getConnectionURL(DB2.Data, username, password)
	maxOpen2, _ := DB2.Data["connection_details"].(map[string]interface{})["max_open_connections"].(json.Number).Int64()
	maxIdel2, _ := DB2.Data["connection_details"].(map[string]interface{})["max_idle_connections"].(json.Number).Int64()
	dbGorm := starter.DbGorm{Type: Conf.Server.DbType}
	dbGorm.Dsn = psd
	dbGorm.MaxIdleConns = int(maxIdel2)
	dbGorm.MaxOpenConns = int(maxOpen2)
	Db = dbGorm.GormInit()
	// log.Println("创建数据库连接————初始化耗时: ", time.Since(tt))
	Log.Info("数据库初始化成功.")
	return nil
}

/*
凭证生命周期管理

返回租约信息，内容可以是初始化数据即包含数据库账号的数据，也可以是租约查询的数据即仅有租约相关信息
租约在刷新时不会修改生成的数据库凭证信息只是对它的有效期做了延长，因此应该可以从cred/list中首先判断leaseID是否存在，然后查看
该leaseID的有效期，判断是否在有效期内，如果在不需要重建数据库连接，反之则销毁租约创建新租约和新的数据库连接。
系统重启————不管租约是否到期，都清空redis缓存以及注销租约，申请数据库凭证重建数据库连接。
系统运行时监控租约————监控租约有效期。1、续约，此时需要重新缓存租约id以及更新redis 租约TTL，数据库连接没有更改；2、注销，更新redis，重建数据库连接。

如果启动时redis中保存有租约有效值,则清空redis缓存以及撤销该租约(如果还存在)
*/
func manageDatabaseCredential() (*api.Secret, error) {
	// 1. 检查redis是否保存租约 如果存在原租约则销毁
	var cred *api.Secret
	var err error
	var leaseID string
	exist, _ := cache.RedisDb.Exists(context.Background(), RedisLeaseKey).Result() //1为true
	if exist == 1 {
		leaseID = GetStoredLeaseID(RedisLeaseKey) // 从存储中读取（如Redis/DB）
		// 开始监控？ 不对 在有效期内的话leaseID也是存在的，这样会把系统运行期间有效期内的租约也清除吧
		if leaseID != "" {
			// log.Printf("当从redis获取有效的租约值: %v", leaseID)
			list, r := vaultClient.Logical().List("sys/leases/lookup/database/creds/" + role)
			if r != nil {
				log.Fatalf("查看租约信息失败: %v", r)
				// Log.Info("当前租约已失效")
			}
			keys := kgo.KConv.ToStr(list.Data["keys"])
			parts := []string{}
			if keys == "" {
				cache.RedisDb.Del(context.Background(), RedisLeaseKey)
			} else {
				parts = strings.Split(keys, "/")
			}
			if keys != "" && strings.Contains(keys, parts[len(parts)-1]) {
				// 撤销旧租约
				err := vaultClient.Sys().Revoke(leaseID)
				if err != nil {
					log.Printf("租约撤销失败: %v", err)
					return nil, err
				}
				log.Println("租约已撤销（超过3天）", leaseID)
				cache.RedisDb.Del(context.Background(), RedisLeaseKey)
			}
		}
	}

	// 2. 创建新的
	cred, err = vaultClient.Logical().Read("database/creds/" + role)
	if err != nil {
		return nil, err
	}
	log.Println("初始化租约=", cred)
	SaveLeaseID(RedisLeaseKey, cred.LeaseID) // 存储新租约ID
	return cred, nil
}

/* 销毁租约
 */
func RevokeLease() {
	// log.Println("定时任务执行......................")
	leaseID := GetStoredLeaseID(RedisLeaseKey)
	if err := vaultClient.Sys().Revoke(leaseID); err != nil {
		// log.Printf("租约定时销毁失败: %v\n", err)
		return
	}
	cred, err := vaultClient.Logical().Read("database/creds/" + role)
	if err != nil {
		// log.Printf("定时生成新租约失败: %v", err)
		return
	}
	log.Println("定时更新租约=", cred.LeaseID)
	SaveLeaseID(RedisLeaseKey, cred.LeaseID)
	updateConnections(cred)
}

// MonitorLease协程，不能依赖主程序 监控租约并自动续期  传入系统状态和系统运行时使用的租约
// 当租约renew之后，不管之前的有效期还有多长时间，都只从此刻开始赋予租约renew中指定的时间，新的过期时间是time.now + increment（s）
func MonitorLease(ctx context.Context) {
	// 错误恢复机制​​  添加 defer 和 recover() 防止协程崩溃  运行失败 报错：runtime error: invalid memory address or nil pointer dereference
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		log.Printf("监控MonitorLease崩溃: %v", r)
	// 	}
	// }()
	log.Println("监控MonitorLease中.....................................")
	ticker := time.NewTicker(5 * time.Minute)
	var times int
	defer ticker.Stop()
	// 判断 最后五分钟 如果remaining<=0的话就说明已经完全过期了
	for {
		select {
		case <-ticker.C:
			leaseID := GetStoredLeaseID(RedisLeaseKey)
			// log.Printf("监控租约 %v 中.....................................", leaseID)
			leaseInfo, _ := vaultClient.Sys().Lookup(leaseID)
			// log.Println("leaseInfo= ", leaseInfo)
			ttl, _ := leaseInfo.Data["ttl"].(json.Number).Int64()
			// log.Println("TTL= ", ttl)
			if ttl <= 299 {
				log.Println("-------TTL 达到阈值，尝试续期...")
				// 续约
				_, r := vaultClient.Sys().Renew(leaseID, int(shortLeaseDuration.Seconds()))
				if r != nil {
					log.Println("续订租约失败")
				}
				times++
				Log.Info("触发续期次数： ", times)
				// Log.Info("角色续订租约： %+v\n", renew)
			}
			if times >= 8 {
				// 三天到期：销毁
				if err := vaultClient.Sys().Revoke(leaseID); err != nil {
					// log.Printf("租约销毁失败: %v\n", err)
					return
				}
				log.Println("已续约8次，销毁租约")
				// 1. 获取凭证
				cred, err := vaultClient.Logical().Read("database/creds/" + role)
				if err != nil {
					// log.Printf("监控生成新租约失败: %v", err)
					return
				}
				Log.Info("MonitorLease更新租约=", cred)
				SaveLeaseID(RedisLeaseKey, cred.LeaseID)
				times = 0
				updateConnections(cred)
			}
		// 主程序关闭
		case <-ctx.Done():
			return
		}
	}
}

func updateConnections(cred *api.Secret) error {
	// 获取数据库连接信息
	db, ee := vaultClient.Logical().Read("database/config/" + hrp)
	if ee != nil {
		// log.Fatalf("获取数据库连接: %v", ee)
		return ee
	}
	// 获取动态角色的租约信息
	username := cred.Data["username"].(string)
	password := cred.Data["password"].(string)
	conn_url := getConnectionURL(db.Data, username, password)
	maxOpen, _ := db.Data["connection_details"].(map[string]interface{})["max_open_connections"].(json.Number).Int64()
	maxIdel, _ := db.Data["connection_details"].(map[string]interface{})["max_idle_connections"].(json.Number).Int64()
	hrdbGorm := starter.DbGorm{Type: Conf.Server.DbType}
	hrdbGorm.Dsn = conn_url
	hrdbGorm.MaxIdleConns = int(maxIdel)
	hrdbGorm.MaxOpenConns = int(maxOpen)
	HrDb = hrdbGorm.GormInit()
	// log.Printf("MonitorLease更新解析odoo数据库连接HrDb= %+v\n", HrDb)
	// pandax数据库连接
	DB2, ee := vaultClient.Logical().Read("database/config/" + iot)
	if ee != nil {
		// log.Fatalf("获取数据库连接: %v", ee)
		return ee
	}
	psd := getConnectionURL(DB2.Data, username, password)
	maxOpen2, _ := DB2.Data["connection_details"].(map[string]interface{})["max_open_connections"].(json.Number).Int64()
	maxIdel2, _ := DB2.Data["connection_details"].(map[string]interface{})["max_idle_connections"].(json.Number).Int64()
	dbGorm := starter.DbGorm{Type: Conf.Server.DbType}
	dbGorm.Dsn = psd
	dbGorm.MaxIdleConns = int(maxIdel2)
	dbGorm.MaxOpenConns = int(maxOpen2)
	Db = dbGorm.GormInit()
	// log.Printf("MonitorLease更新解析Pandax数据库连接Db= %+v\n", Db)
	return nil
}

// 存储/读取租约ID（需根据实际存储实现）
func SaveLeaseID(key string, value string) {
	ctx := context.Background()
	redis_lease := new(RedisLease)
	redis_lease.CreatedAt = time.Now()
	redis_lease.LeaseID = value
	redis_lease.Uuid = uuid.New().String()
	_, err := cache.RedisDb.HMSet(ctx, key, map[string]interface{}{
		"uuid":       redis_lease.Uuid,
		"lease_id":   redis_lease.LeaseID,
		"created_at": redis_lease.CreatedAt,
	}).Result()

	if err == nil {
		// 设置 TTL，自动清理过期凭证
		cache.RedisDb.Expire(ctx, key, renewalThreshold)
	}
}

func GetStoredLeaseID(key string) string {
	value, err := cache.RedisDb.HGetAll(context.Background(), key).Result()
	if err != nil {
		log.Fatalf("读取Redis失败: %v", err)
	}
	return value["lease_id"]
}
