package config

import "fmt"

type Kafka struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
}

func (e Kafka) Url() string {
	return fmt.Sprintf("%s:%d", e.Host, e.Port)
}
