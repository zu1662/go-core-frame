package config

import "github.com/spf13/viper"

// RedisConfig 初始化
var RedisConfig = new(Redis)

// Redis 结构体
type Redis struct {
	Password string
	Host     string
	Port     string
	DBName   int
}

// InitRedis mysql 配置结构体初始化
func InitRedis(cfg *viper.Viper) *Redis {
	return &Redis{
		Password: cfg.GetString("password"),
		Host:     cfg.GetString("host"),
		Port:     cfg.GetString("port"),
		DBName:   cfg.GetInt("db"),
	}
}
