package database

import (
	"go-core-frame/global"
	"go-core-frame/pkg/config"

	"github.com/go-redis/redis"
)

// RedisDB info
var RedisDB *redis.Client

// RedisSetup Gorm & Mysql 初始化
func RedisSetup() {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisConfig.Host + ":" + config.RedisConfig.Port,
		Password: config.RedisConfig.Password, // no password set
		DB:       config.RedisConfig.DBName,   // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		global.Logger.Fatal(" Redis connect error :", err)
	} else {
		global.Logger.Info(" Redis connect success !")
		RedisDB = client
	}

	global.Redis = RedisDB
}
