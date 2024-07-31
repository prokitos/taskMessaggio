package models

type Config struct {
	Env      string         `yaml:"env" env-default:"local"`
	Database DatabaseConfig `yaml:"postgres"`
	Kafka    string         `yaml:"kafka"`
	Server   ServerConfig   `yaml:"server"`
}

type DatabaseConfig struct {
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
	Host   string `yaml:"host"`
	Name   string `yaml:"name"`
	Reload bool   `yaml:"reload"`
	Port   string `yaml:"port"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}
