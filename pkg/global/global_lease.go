package global

import (
	"log"
	"time"

	"github.com/hashicorp/vault/api"
)

const (
	renewalThreshold   = 3 * 24 * time.Hour // 3天有效期阈值
	shortLeaseDuration = 1 * time.Hour      // 短租约续期时长
)

var (
	vaultClient *api.Client
	dbRole      = "role1" // Vault数据库动态角色名称
)

func main() {
	initVaultClient()

	// 主循环（实际应用中可改为后台goroutine）
	for {
		cred, err := manageDatabaseCredential()
		if err != nil {
			log.Printf("凭证管理失败: %v", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		// 使用凭证连接数据库（示例）
		useDatabaseCredential(cred)

		// 检查间隔（根据实际需求调整）
		time.Sleep(5 * time.Minute)
	}
}

// 初始化Vault客户端
func initVaultClient() {
	config := api.DefaultConfig()
	config.Address = "http://192.168.0.30:8200" // Vault地址

	var err error
	vaultClient, err = api.NewClient(config)
	if err != nil {
		log.Fatalf("创建Vault客户端失败: %v", err)
	}

	vaultClient.SetToken("hvs.xxx") // 替换为实际Token
}

// 凭证生命周期管理
func manageDatabaseCredential() (*api.Secret, error) {
	// 1. 检查现有租约
	leaseID := getStoredLeaseID() // 从存储中读取（如Redis/DB）
	var cred *api.Secret
	var err error

	if leaseID != "" {
		cred, err = renewOrRotateLease(leaseID)
		if err != nil {
			log.Printf("租约管理失败: %v", err)
		}
	}

	// 2. 如果没有有效租约，创建新的
	if cred == nil {
		cred, err = vaultClient.Logical().Read("database/creds/" + dbRole)
		if err != nil {
			return nil, err
		}
		saveLeaseID(cred.LeaseID) // 存储新租约ID
		log.Println("创建新租约")
	}

	return cred, nil
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

// 存储/读取租约ID（需根据实际存储实现）
func saveLeaseID(leaseID string) {
	// 示例：存储到Redis或文件
}

func getStoredLeaseID() string {
	// 示例：从存储中读取
	return ""
}

// 使用数据库凭证（示例）
func useDatabaseCredential(cred *api.Secret) {
	username := cred.Data["username"].(string)
	password := cred.Data["password"].(string)

	// 实际连接数据库逻辑
	log.Printf("使用凭证连接数据库 - 用户名: %s", username)
}
