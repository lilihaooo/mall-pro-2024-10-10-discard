package config

import "fmt"

type Es struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
}

func (e Es) Url() string {
	return fmt.Sprintf("%s:%d", e.Host, e.Port)
}
