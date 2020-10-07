package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// 应用配置项
var cfgApplication *viper.Viper

// 日志配置项
var cfgLogger *viper.Viper

// jwt 配置项
var cfgJWT *viper.Viper

// mysql 配置项
var cfgMysql *viper.Viper

// redis 配置项
var cfgRedis *viper.Viper

// Setup 配置文件初始化
func Setup(path string) {
	// 设置配置文件信息
	viper.SetConfigType("yml")
	viper.SetConfigFile(path)

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}

	// 应用配置初始化
	cfgApplication = viper.Sub("settings.application")
	if cfgApplication == nil {
		panic("No found settings.application in the configuration")
	}
	ApplicationConfig = InitApplication(cfgApplication)

	// 日志配置初始化
	cfgLogger = viper.Sub("settings.logger")
	if cfgLogger == nil {
		panic("No found settings.logger in the configuration")
	}
	LoggerConfig = InitLogger(cfgLogger)

	// JWT配置初始化
	cfgJWT = viper.Sub("settings.jwt")
	if cfgJWT == nil {
		panic("No found settings.jwt in the configuration")
	}
	JWTConfig = InitJWT(cfgJWT)

	// 数据库配置初始化
	cfgMysql = viper.Sub("settings.mysql")
	if cfgMysql == nil {
		panic("No found settings.mysql in the configuration")
	}
	MysqlConfig = InitMysql(cfgMysql)

	// redis 配置初始化
	cfgRedis = viper.Sub("settings.redis")
	if cfgMysql == nil {
		panic("No found settings.redis in the configuration")
	}
	RedisConfig = InitRedis(cfgRedis)
}
