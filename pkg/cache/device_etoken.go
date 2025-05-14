package cache

import (
	"context"
	"time"

	"github.com/PandaXGO/PandaKit/rediscli"
)

var RedisDb *rediscli.RedisDB

// 存储数据库凭证
func SetDatabaseCredential(key string, value any, duration time.Duration) error {
	return RedisDb.Set(key, value, duration)
}

// 获取数据库凭证
// func GetDatabaseCredential(key string, value any) error {
// 	return RedisDb.Get(key, value)
// }

// SetDeviceEtoken key 是设备的时候为token， 是子设备的时候为设备编码
func SetDeviceEtoken(key string, value any, duration time.Duration) error {
	return RedisDb.Set(key, value, duration)
}

// GetDeviceEtoken value 是参数指针
func GetDeviceEtoken(key string, value interface{}) error {
	return RedisDb.Get(key, value)
}

// DelDeviceEtoken 删除指定的key
func DelDeviceEtoken(key string) error {
	return RedisDb.Del(context.Background(), key).Err()
}

func ExistsDeviceEtoken(key string) bool {
	//原代码为： exists, _ := RedisDb.Exists(RedisDb.Context(), key).Result()
	exists, _ := RedisDb.Exists(context.Background(), key).Result()
	return exists == 1
}
