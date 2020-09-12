package config

import "github.com/spf13/viper"

// MysqlConfig 初始化
var MysqlConfig = new(Mysql)

// Mysql 结构体
type Mysql struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
	Config   string
}

// InitMysql mysql 配置结构体初始化
func InitMysql(cfg *viper.Viper) *Mysql {
	return &Mysql{
		Username: cfg.GetString("username"),
		Password: cfg.GetString("password"),
		Host:     cfg.GetString("host"),
		Port:     cfg.GetString("port"),
		DBName:   cfg.GetString("dbname"),
		Config:   cfg.GetString("config"),
	}
}
