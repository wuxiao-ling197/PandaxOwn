package config

import "fmt"

type Server struct {
	Port        int            `mapstructure:"port" json:"port" yaml:"port"`
	GrpcPort    int            `mapstructure:"grpc-port" json:"grpc-port" yaml:"grpc-port"`
	TcpPort     int            `mapstructure:"tcp-port" json:"tcp-port" yaml:"tcp-port"`
	HttpPort    int            `mapstructure:"http-port" json:"http-port" yaml:"http-port"`
	Model       string         `mapstructure:"model" json:"model" yaml:"model"`
	Cors        bool           `mapstructure:"cors" json:"cors" yaml:"cors"`
	Rate        *Rate          `mapstructure:"rate" json:"rate" yaml:"rate"`
	IsInitTable bool           `mapstructure:"isInitTable" json:"isInitTable" yaml:"isInitTable"`
	DbType      string         `mapstructure:"db-type" json:"db-type" yaml:"db-type"`
	ExcelDir    string         `mapstructure:"excel-dir" json:"excel-dir" yaml:"excel-dir"`
	Tls         *Tls           `mapstructure:"tls" json:"tls" yaml:"tls"`
	Static      *[]*Static     `mapstructure:"static" json:"static" yaml:"static"`
	StaticFile  *[]*StaticFile `mapstructure:"static-file" json:"static-file" yaml:"static-file"`
}

func (s *Server) GetPort() string {
	return fmt.Sprintf(":%d", s.Port)
}

type Static struct {
	RelativePath string `yaml:"relative-path"`
	Root         string `yaml:"root"`
}

type StaticFile struct {
	RelativePath string `yaml:"relative-path"`
	Filepath     string `yaml:"filepath"`
}

type Tls struct {
	Enable   bool   `yaml:"enable"`    // 是否启用tls
	KeyFile  string `yaml:"key-file"`  // 私钥文件路径
	CertFile string `yaml:"cert-file"` // 证书文件路径
}

type Rate struct {
	Enable  bool    `yaml:"enable"`   // 是否限流
	RateNum float64 `yaml:"rate-num"` // 限流数量
}
