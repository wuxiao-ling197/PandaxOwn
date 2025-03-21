package config

import "path"

type Log struct {
	Level string   `mapstructure:"level" json:"level" yaml:"level"`
	File  *LogFile `mapstructure:"file" json:"file" yaml:"file"`
}

type LogFile struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

// 获取完整路径文件名
func (l *LogFile) GetFilename() string {
	var filepath, filename string
	if l == nil {
		return ""
	}
	if fp := l.Path; fp == "" {
		filepath = "./"
	} else {
		filepath = fp
	}
	if fn := l.Name; fn == "" {
		filename = "default.log"
	} else {
		filename = fn
	}

	return path.Join(filepath, filename)
}
