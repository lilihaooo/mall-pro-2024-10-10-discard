package config

// Config 配置文件
type Config struct {
	Server `yaml:"server"`
	Mysql  `yaml:"mysql"`
	Es     `yaml:"es"`
}
