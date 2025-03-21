package config

import (
	"github.com/PandaXGO/PandaKit/biz"
)

type Jwt struct {
	Key        string `mapstructure:"key" json:"key" yaml:"key"`
	ExpireTime int64  `mapstructure:"expire-time" json:"expire-time" yaml:"expire-time"` // 过期时间，单位分钟
}

func (j *Jwt) Valid() {
	biz.IsTrue(j.Key != "", "项目配置之 [jwt.key] 不能为空")
	biz.IsTrue(j.ExpireTime != 0, "项目配置之 [jwt.expire-time] 不能为空")
}
