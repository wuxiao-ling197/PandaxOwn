package config

type Oss struct {
	Endpoint   string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	AccessKey  string `mapstructure:"accessKey" json:"accessKey" yaml:"accessKey"`
	SecretKey  string `mapstructure:"secretKey" json:"secretKey" yaml:"secretKey"`
	BucketName string `mapstructure:"bucketName" json:"bucketName" yaml:"bucketName"`
	UseSSL     bool   `mapstructure:"useSSL" json:"useSSL" yaml:"useSSL"`
}
