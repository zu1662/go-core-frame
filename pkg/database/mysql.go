package database

import (
	"go-core-frame/global"
	"go-core-frame/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// MysqlDB info
var MysqlDB *gorm.DB

// MysqlSetup Gorm & Mysql 初始化
func MysqlSetup() {
	var connectDsn = config.MysqlConfig.Username + ":" +
		config.MysqlConfig.Password + "@tcp(" +
		config.MysqlConfig.Host + ":" +
		config.MysqlConfig.Port + ")/" +
		config.MysqlConfig.DBName + "?" +
		config.MysqlConfig.Config
	//连接数据库
	gormLogLevel := logger.Silent
	if config.ApplicationConfig.Mode == "dev" {
		gormLogLevel = logger.Info
	}
	MysqlDB, err := gorm.Open(mysql.Open(connectDsn), &gorm.Config{
		Logger: logger.Default.LogMode(gormLogLevel),
	})
	if err != nil {
		global.Logger.Fatal(" Mysql connect error :", err)
	} else {
		global.Logger.Info(" Mysql connect success !")
	}

	global.DB = MysqlDB
}
