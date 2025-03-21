package config

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"reflect"
	"strconv"
	"sync"

	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/utils"
	"github.com/mitchellh/mapstructure"
)

func InitConfig(configFilePath string) *Config {
	// 获取启动参数中，配置文件的绝对路径
	path, _ := filepath.Abs(configFilePath)
	startConfigParam = &CmdConfigParam{ConfigFilePath: path}
	// 读取配置文件信息
	yc := &Config{}
	if err := utils.LoadYml(startConfigParam.ConfigFilePath, yc); err != nil {
		panic(any(fmt.Sprintf("读取配置文件[%s]失败: %s", startConfigParam.ConfigFilePath, err.Error())))
	}
	// 校验配置文件内容信息
	yc.Valid()

	return yc

}

// 2025-1-13 定义一个全局变量和一个互斥锁 解决vault变量保存结构体时不完整的问题
var (
	vc = &Config{}
	mu sync.Mutex // 用于保护访问全局变量的互斥锁
)

// add 2025-1-13 从Vault读取保存配置 如果结构中存在错误项、遗漏项 请在结构体中添加定义/描述？mapstructure:"port" json:"port"
func InitVaultConfig(config map[string]interface{}) *Config {
	mu.Lock()         //加锁，保证线程安全
	defer mu.Unlock() // 解锁，确保在函数返回时解锁
	// 使用 mapstructure 提供的 Decode 的 DecoderConfig 配置，配合自定义的转换函数，以支持将字符串转换为相应的目标类型
	decoderConfig := mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   vc,
		TagName:  "mapstructure",
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			// 将字符串转换为 int
			func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
				if f.Kind() == reflect.String {
					if t.Kind() == reflect.Int {
						if strValue, ok := data.(string); ok {
							return strconv.Atoi(strValue) // 将字符串转换为 int
						}
					} else if t.Kind() == reflect.Int64 {
						if strValue, ok := data.(string); ok {
							return strconv.ParseInt(strValue, 10, 64) // 将字符串转换为 int
						}
					} else if t.Kind() == reflect.Bool {
						// 字符串转 bool
						if strValue, ok := data.(string); ok {
							switch strValue {
							case "true":
								return true, nil
							case "false":
								return false, nil
							default:
								return nil, fmt.Errorf("invalid value for bool: %s", strValue)
							}
						}
					}
				}
				return data, nil
			},
		),
	}

	// 使用自定义配置解码
	decoder, err := mapstructure.NewDecoder(&decoderConfig)
	if err != nil {
		log.Fatalf("Error creating decoder: %v", err)
	}

	err = decoder.Decode(config)
	if err != nil {
		log.Fatalf("Error decoding: %v", err)
	}

	// log.Print(vc.App)
	// log.Print(vc.Server)
	// log.Print(vc.Casbin)
	// log.Print(vc.Postgresql)
	// log.Print(vc.Redis)
	// log.Print(vc.Taos)
	// log.Print(vc.Mqtt)
	// log.Print(vc.Oss)
	// log.Print(vc.Log)
	// log.Print(vc.Jwt)
	vc.Valid()
	return vc
}

//end

// 启动配置参数
type CmdConfigParam struct {
	ConfigFilePath string // -e  配置文件路径
}

// 启动可执行文件时的参数
var startConfigParam *CmdConfigParam

// yaml配置文件映射对象 Hrdb
type Config struct {
	App        *App        `yaml:"app" json:"app"`
	Server     *Server     `yaml:"server" json:"server"`
	Queue      *Queue      `yaml:"queue" json:"queue"`
	Jwt        *Jwt        `yaml:"jwt" json:"jwt"`
	Redis      *Redis      `yaml:"redis" json:"redis"`
	Mysql      *Mysql      `yaml:"mysql" json:"mysql"`
	Postgresql *Postgresql `yaml:"postgresql" json:"postgresql"`
	Hrdb       *Postgresql `yaml:"hrdb" json:"hrdb"`
	Oss        *Oss        `yaml:"oss" json:"oss"`
	Taos       *Taos       `yaml:"taos" json:"taos"`
	Mqtt       *Mqtt       `yaml:"mqtt" json:"mqtt"`
	Casbin     *Casbin     `yaml:"casbin" json:"casbin"`
	Gen        *Gen        `yaml:"gen" json:"gen"`
	Ys         *Ys         `yaml:"ys" json:"ys"`
	Log        *Log        `yaml:"log" json:"log"`
}

// 配置文件内容校验
func (c *Config) Valid() {
	biz.IsTrue(c.Jwt != nil, "项目配置的[jwt]信息不能为空")
	c.Jwt.Valid()
}

// 获取执行可执行文件时，指定的启动参数
func getStartConfig() *CmdConfigParam {
	configFilePath := flag.String("e", "./config.yml", "配置文件路径，默认为可执行文件目录")
	flag.Parse()
	// 获取配置文件绝对路径
	path, _ := filepath.Abs(*configFilePath)
	sc := &CmdConfigParam{ConfigFilePath: path}
	return sc
}
