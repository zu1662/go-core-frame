package database

import (
	"go-core-frame/global"
	"go-core-frame/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MysqlDB info
var MysqlDB *gorm.DB

// Setup Gorm & Mysql 初始化
func Setup() {
	var connectDsn = config.MysqlConfig.Username + ":" +
		config.MysqlConfig.Password + "@tcp(" +
		config.MysqlConfig.Host + ":" +
		config.MysqlConfig.Port + ")/" +
		config.MysqlConfig.DBName + "?" +
		config.MysqlConfig.Config
	//连接数据库
	MysqlDB, err := gorm.Open(mysql.Open(connectDsn), &gorm.Config{})
	if err != nil {
		global.Logger.Fatal(" Mysql connect error :", err)
	} else {
		global.Logger.Info(" Mysql connect success !")
	}

	global.DB = MysqlDB
}
