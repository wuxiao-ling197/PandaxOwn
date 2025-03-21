package config

type Taos struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // 服务器地址:端口
	Username string `mapstructure:"username" json:"username" yaml:"username"` // 数据库用户名
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 数据库密码
	Database string `mapstructure:"database" json:"database" yaml:"database"`
	Config   string `mapstructure:"config" json:"config" yaml:"config"`
}
