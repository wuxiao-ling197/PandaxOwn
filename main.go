package main

// 总是被最后一个初始化，因为要先加载其他依赖文件

import (
	"context"
	"log"
	"os"
	"os/signal"
	"pandax/iothub"
	"pandax/pkg/global"
	"pandax/pkg/initialize"
	"pandax/pkg/middleware"
	"syscall"
	"time"

	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/robfig/cron/v3"

	vault "github.com/hashicorp/vault/api"
	"github.com/spf13/cobra"
)

var (
	configFile string
	configData map[string]interface{} //映射为json格式的object对象
)

var rootCmd = &cobra.Command{
	Use:   "pandax is the main component in the panda.",
	Short: `pandax is go go-restful frame`,
	PreRun: func(cmd *cobra.Command, args []string) {
		// cc := getVault()
		// if cc != nil {
		// 	// 读取配置文件
		// 	// global.Conf = config.InitConfig(configFile)
		// 	global.Conf = config.InitVaultConfig(cc["server"].(map[string]interface{}))
		// 	global.Conf = config.InitVaultConfig(cc)

		// 	// 代码迁移至global位置了
		// 	global.Log = logger.InitLog(global.Conf.Log.File.GetFilename(), global.Conf.Log.Level)
		// 	dbGorm := starter.DbGorm{Type: global.Conf.Server.DbType}
		// 	if global.Conf.Server.DbType == "mysql" {
		// 		dbGorm.Dsn = global.Conf.Mysql.Dsn()
		// 		dbGorm.MaxIdleConns = global.Conf.Mysql.MaxIdleConns
		// 		dbGorm.MaxOpenConns = global.Conf.Mysql.MaxOpenConns
		// 	} else {
		// 		dbGorm.Dsn = global.Conf.Postgresql.PgDsn()
		// 		dbGorm.MaxIdleConns = global.Conf.Postgresql.MaxIdleConns
		// 		dbGorm.MaxOpenConns = global.Conf.Postgresql.MaxOpenConns
		// 	}
		// 	// global.Log.Infof("dbGorm=  %v", global.Conf.Postgresql)
		// 	// global.Log.Infof("dbGorm type=  %v", dbGorm.Type)
		// 	global.Db = dbGorm.GormInit()
		// 	hrdbGorm := starter.DbGorm{Type: global.Conf.Server.DbType}
		// 	hrdbGorm.Dsn = global.Conf.Hrdb.PgDsn()
		// 	hrdbGorm.MaxIdleConns = global.Conf.Hrdb.MaxIdleConns
		// 	hrdbGorm.MaxOpenConns = global.Conf.Hrdb.MaxOpenConns
		// 	global.HrDb = hrdbGorm.GormInit()
		// 	starter.Db = global.Db
		// 	global.Log.Infof("global.Db=  %v  连接成功", global.Db)

		// 	client, err := rediscli.NewRedisClient(global.Conf.Redis.Host, global.Conf.Redis.Password, global.Conf.Redis.Port, global.Conf.Redis.Db)
		// 	if err != nil {
		// 		global.Log.Panic("Redis连接错误")
		// 	} else {
		// 		global.Log.Info("Redis连接成功")
		// 	}
		// 	cache.RedisDb = client
		// 	tDengine, err := tdengine.InitTd(global.Conf.Taos.Username, global.Conf.Taos.Password, global.Conf.Taos.Host, global.Conf.Taos.Database)
		// 	if err != nil {
		// 		global.Log.Panic("Tdengine连接错误")
		// 	} else {
		// 		global.Log.Info("Tdengine连接成功")
		// 	}
		// 	global.TdDb = tDengine
		// 	initialize.InitTable()
		// 	// 初始化事件监听
		// 	go initialize.InitEvents()
		// } else {
		// 	global.Log.Panic("请配置config")
		// }
	},
	Run: func(cmd *cobra.Command, args []string) {
		// 前置 函数
		restfulx.UseBeforeHandlerInterceptor(middleware.PermissionHandler)
		// 后置 函数
		restfulx.UseAfterHandlerInterceptor(middleware.LogHandler)
		restfulx.UseAfterHandlerInterceptor(middleware.OperationHandler)

		app := initialize.InitRouter()
		global.Log.Info("路由初始化完成")
		app.Start(context.TODO())
		// 2025-4-14 开启租约监控
		ctx, cancel := context.WithCancel(context.Background())
		// go global.MonitorLease(ctx)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					// 添加定时任务（Cron 表达式）
					c := cron.New(cron.WithLocation(time.Local)) // 设置本地时区
					_, err := c.AddFunc("35 11 * * 2", func() {  // 每周五 23:59
						global.RevokeLease()
					})
					if err != nil {
						panic("定时任务配置错误: " + err.Error())
					}
					c.Start()
					defer c.Stop()

					global.MonitorLease(ctx) // 监控循环可自动重启
				}
			}
		}()
		//开启IOTHUB
		go iothub.InitIothub()
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGTERM, os.Interrupt)
		<-stop
		if err := app.Stop(context.TODO()); err != nil {
			log.Fatalf("fatal app stop: %s", err)
			os.Exit(-3)
		}
		// 优雅关闭
		cancel()
	},
}

// 初始化
func init() {
	// getVault()
	rootCmd.Flags().StringVar(&configFile, "config", getEnvStr("PANDA_CONFIG", "./config.yml"), "panda config file path.")
}

// 返回配置文件信息
func getEnvStr(env string, defaultValue string) string {
	v := os.Getenv(env)
	if v == "" {
		return defaultValue
	}
	return v
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		rootCmd.PrintErrf("panda root cmd execute: %s", err)
		os.Exit(1)
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
	log.Print("远程配置端连接成功")

	// Authenticate 这里以超级管理员进行token认证
	// WARNING: This quickstart uses the root token for our Vault dev server.
	// Don't do this in production!
	client.SetToken("hvs.rvdWltAtDl9DFUt5zaoPSYmr")
	log.Print("Authenticate success!")
	secret, err := client.KVv2("pandax05eric").Get(ctx, "pandax")
	if err != nil {
		log.Fatalf("unable to read secret: %v", err)
	}
	// server, _ := secret.Data["server"].(map[string]interface{})
	// taos, _ := server["taos"].(map[string]interface{})
	// log.Print(taos, taos["database"].(string), taos["host"].(string), taos["password"].(string))
	configData = secret.Data
	return configData
}
