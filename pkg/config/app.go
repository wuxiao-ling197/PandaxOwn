package config

import "fmt"

type App struct {
	Name    string `mapstructure:"name" json:"name" yaml:"name"`
	Version string `mapstructure:"version" json:"version" yaml:"version"`
}

func (a *App) GetAppInfo() string {
	return fmt.Sprintf("[%s:%s]", a.Name, a.Version)
}
