package config

type Ys struct {
	AppKey string `mapstructure:"appKey" json:"appKey" yaml:"appKey"`
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"`
}
