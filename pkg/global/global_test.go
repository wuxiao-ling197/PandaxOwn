package global

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	vault "github.com/hashicorp/vault/api"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestLease(t *testing.T) {
	// saveLeaseID(redisLeaseKey, "beoqq1ptKm4E0l23TyzqJitg")
	// v := getStoredLeaseID(redisLeaseKey)
	// log.Printf("获取= %+v\n", v)
	// cache.RedisDb.Del(context.Background(), redisPwdKey)
	DatabaseManage()
}

func TestAPI(t *testing.T) {
	config := vault.DefaultConfig()
	config.Address = "http://192.168.0.30:8200"
	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to initialize Vault client: %v", err)
	}
	client.SetToken("hvs.rvdWltAtDl9DFUt5zaoPSYmr")

	// 查看角色当前可用租约列表
	// list, r := client.Logical().List("sys/leases/lookup/database/creds/pAdmin")
	// if r != nil {
	// 	log.Fatalf("查看租约信息失败: %v", r)
	// }
	// fmt.Printf("logic读取可用租约列表： %+v,%T\n", list, list.Data["keys"])
	// keys := kgo.KConv.ToStr(list.Data["keys"])
	// fmt.Printf("99= %T\n", keys)
	// fmt.Println("00  ", keys)
	// fmt.Println("存在=", strings.Contains(keys, "aq1IgmTvsklS2o1UG93EXYJ9"))

	// for _, v := range keys {
	// 	val := reflect.ValueOf(v)
	// 	fmt.Printf("99= %+v  %T\n", val, val)
	// 	if val.Interface() == "aq1IgmTvsklS2o1UG93EXYJ6" {
	// 		fmt.Println("存在@！")
	// 	}
	// }
	/*&{RequestID:249e3cf9-7547-3d51-8990-ed6e8dd53445 LeaseID: LeaseDuration:0 Renewable:false
	Data:map[keys:[aq1IgmTvsklS2o1UG93EXYJ6]] Warnings:[] Auth:<nil> WrapInfo:<nil> MountType:system}*/

	// 生成数据库凭据，有新的租约和数据库账户信息生成 等同于CLI命令：vault read database/creds/pAdmin
	// secret, err := client.Logical().Read("database/creds/pAdmin")
	// if err != nil {
	// 	log.Fatalf("获取凭证失败: %v", err)
	// }
	// fmt.Printf("返回凭证信息secret： %+v\n", secret)
	/*&{RequestID:61778435-0231-4d18-791e-d89055f6cb67 LeaseID:database/creds/pAdmin/fY20y0kTdEdqgwNTAtDGARzW
	LeaseDuration:3600 Renewable:true Data:map[password:lvd-GhGus2r3vVrZEape
	username:v-root-pAdmin-24Js0mFty7VkFx670gci-1743658141]
	Warnings:[] Auth:<nil> WrapInfo:<nil> MountType:database}*/

	// 读取动态数据库凭证
	// params := make(map[string][]string)  //ReadRawWithData
	// params["lease_id"] = append(params["lease_id"], "aq1IgmTvsklS2o1UG93EXYJ6")
	// params := map[string]interface{}{
	// 	"lease_id": "aq1IgmTvsklS2o1UG93EXYJ6",
	// } //Write
	// fmt.Printf("lookup查看租约信息结果：%+v\n", params)
	// passport, err := client.Sys().Lookup("database/creds/pAdmin/aq1IgmTvsklS2o1UG93EXYJ6") //.Read("sys/leases/lookup")
	// if err != nil {
	// 	log.Fatalf("查看租约信息失败: %v", err)
	// }
	// fmt.Printf("lookup查看租约信息结果：%+v\n", passport)
	/*&{RequestID:d9dc202b-4ff7-3e6c-099a-f448cc4e1889 LeaseID: LeaseDuration:0 Renewable:false
	Data:map[
	expire_time:2025-04-11T17:13:20.0276206+08:00
	id:database/creds/pAdmin/aq1IgmTvsklS2o1UG93EXYJ6
	issue_time:2025-04-08T17:13:20.0271053+08:00
	last_renewal:<nil>
	renewable:true
	ttl:98154
	] Warnings:[] Auth:<nil> WrapInfo:<nil> MountType:system}
	续订后租约信息：
	&{RequestID:80a7ae24-93d6-8977-f6d9-db02c991d394 LeaseID: LeaseDuration:0 Renewable:false
	Data:map[
	expire_time:2025-04-10T14:18:23.8315719+08:00
	id:database/creds/pAdmin/aq1IgmTvsklS2o1UG93EXYJ6
	issue_time:2025-04-08T17:13:20.0271053+08:00
	last_renewal:2025-04-10T14:01:43.8315719+08:00
	renewable:true
	ttl:964
	] Warnings:[] Auth:<nil> WrapInfo:<nil> MountType:system}
	*/

	// 续订租约
	renew, r := client.Sys().Renew("database/creds/pAdmin/GOWBZyApScv85lxgMiLZe576", 1000)
	if r != nil {
		log.Fatalf("续订租约失败: %v", r)
	}
	log.Printf("续订租约结果： %+v\n", renew)
	/*&{RequestID:ebcb8bf7-3efb-c245-ae29-6cc82e35639b LeaseID:database/creds/pAdmin/aq1IgmTvsklS2o1UG93EXYJ6 LeaseDuration:1000
	Renewable:true Data:map[] Warnings:[] Auth:<nil> WrapInfo:<nil> MountType:system}*/

	// 注销租约
	// erro := client.Sys().Revoke("database/creds/pAdmin/KIot2YtQgnlXvY9NlMOFtdIZ")
	// if err != nil {
	// 	log.Printf("租约撤销失败: %v", erro)
	// }
	// log.Println("租约已撤销")
}

func TestConnectStr(t *testing.T) {
	dsn := "postgresql://v-root-pAdmin-sea1IrDM8wKe8OiI4wbl-1744093361:Bu-DoE7ZV06HUvHO7wKp@192.168.0.30:5432/odoo18?sslmode=disable"

	str, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("sql失败: %v", err)
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: str,
	}), &gorm.Config{})
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("实例化失败: %v", err)
	}

	// 6. 验证连接有效性
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	log.Printf("连接成功:%+v\n", sqlDB)

}

func TestMonitor(t *testing.T) {
	// timer 需要人为调用Reset方法 简单的定时服务 cg
	// timer := time.NewTimer(5 * time.Second)
	// select {
	// case <-timer.C:
	// 	fmt.Println("触发监控 at ", time.Now())
	// 	//运行：2025/04/14 13:42:25.... 输出：触发监控  2025-04-14 13:42:30.591095785 +0800 CST m=+5.026884804
	// }
	// timer.Stop()

	// 每个时间间隔均会执行 sb
	// timer := time.NewTimer(5 * time.Second)
	// for {
	// 	timer.Reset(5 * time.Second)
	// 	select {
	// 	case <-timer.C:
	// 		fmt.Println("触发监控 at ", time.Now())
	// 		/* 输出：
	// 		2025/04/14 13:45:44 远程配置端连接成功！
	// 		2025/04/14 13:45:44 Authenticate success！
	// 		=== RUN   TestMonitor
	// 		触发监控 at  2025-04-14 13:45:49.513489685 +0800 CST m=+5.023577354
	// 		触发监控 at  2025-04-14 13:45:54.517494709 +0800 CST m=+10.027582375
	// 		触发监控 at  2025-04-14 13:45:59.521566274 +0800 CST m=+15.031653942
	// 		触发监控 at  2025-04-14 13:46:04.525487732 +0800 CST m=+20.035575401
	// 		触发监控 at  2025-04-14 13:46:09.529556438 +0800 CST m=+25.039644106
	// 		panic: test timed out after 30s */
	// 	}
	// }

	// ticker 会自动续期
	// ticker := time.NewTicker(5 * time.Second)
	// for range ticker.C {
	// 	fmt.Println("触发监控 at ", time.Now())
	// 	/*输出：
	// 	2025/04/14 13:50:00 远程配置端连接成功！
	// 	2025/04/14 13:50:00 Authenticate success！
	// 	=== RUN   TestMonitor
	// 	触发监控 at  2025-04-14 13:50:05.841495547 +0800 CST m=+5.022802742
	// 	触发监控 at  2025-04-14 13:50:10.845487941 +0800 CST m=+10.026795132
	// 	触发监控 at  2025-04-14 13:50:15.84550434 +0800 CST m=+15.026811537
	// 	触发监控 at  2025-04-14 13:50:20.840761438 +0800 CST m=+20.022068628
	// 	触发监控 at  2025-04-14 13:50:25.841022814 +0800 CST m=+25.022330009
	// 	触发监控 at  2025-04-14 13:50:30.840712343 +0800 CST m=+30.022019538
	// 	panic: test timed out after 30s
	// 	        running tests:
	// 	                TestMonitor (30s)
	// 	*/
	// }
	// ticker.Stop()

	// ticker 定时监控并指定监控次数
	ticker := time.NewTicker(5 * time.Second)
	fmt.Println("触发监控 at ", time.Now())
	go func(t *time.Ticker) {
		times := 0 //监控次数
		for {
			<-t.C
			fmt.Println("监控中... at ", time.Now())
			times++
			if times >= 3 {
				fmt.Println("触发监控次数： ", times)
				ticker.Stop()
			}
		}
	}(ticker)
	time.Sleep(30 * time.Second)
	fmt.Println("监控结束 at ", time.Now())
	/*输出：
		=== RUN   TestMonitor
	触发监控 at  2025-04-14 14:02:05.299767165 +0800 CST m=+0.025731854
	监控中... at  2025-04-14 14:02:10.301794456 +0800 CST m=+5.027759145
	监控中... at  2025-04-14 14:02:15.303826635 +0800 CST m=+10.029791324
	监控中... at  2025-04-14 14:02:20.301149442 +0800 CST m=+15.027114131
	触发监控次数：  3
	监控结束 at  2025-04-14 14:02:35.301501603 +0800 CST m=+30.027466294
	--- PASS: TestMonitor (30.00s)
	PASS
	ok      pandax/pkg/global       30.030s
	*/
}

func TestDuration(t *testing.T) {
	// m := 1 * time.Hour
	// log.Println(int(m.Seconds()))

	// now := time.Now()
	// fmt.Println("Current time:", now)
	// fmt.Println("Location:", now.Location()) //Location: Local

	// tz := os.Getenv("TZ")
	// if tz != "" {
	// 	fmt.Printf("环境变量TZ设置为: %s\n", tz)
	// } else {
	// 	fmt.Println("环境变量TZ未设置，使用系统默认时区")
	// }

	// loc, err := time.LoadLocation("Asia/Shanghai")
	// if err != nil {
	// 	fmt.Println("加载时区失败:", err)
	// 	return
	// }
	// shanghaiTime := time.Now().In(loc)
	// fmt.Printf("上海时间: %s\n", shanghaiTime)

	now := time.Now()
	// 输出当前时间及其时区信息
	fmt.Printf("当前时间: %s\n", now)
	fmt.Printf("时区名称: %s\n", now.Location())
	fmt.Printf("时区偏移: %s\n", now.Format("-07:00 MST"))
}
