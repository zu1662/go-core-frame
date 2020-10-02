package models

import (
	"go-core-frame/global"
	"time"
)

// LoginLog 操作日志Model
type LoginLog struct {
	ID          int       `gorm:"" json:"id"`  // 数据库id
	UserName    string    `json:"userName"`    // 操作人
	IPAddress   string    `json:"ipAddress"`   // IP地址
	IPLocation  string    `json:"ipLocation"`  // IP归属地
	Browser     string    `json:"browser"`     // 浏览器
	OS          string    `json:"os"`          // 操作系统
	Result      string    `json:"reslut"`      // 操作结果
	Description string    `json:"description"` // 操作描述（user-agent）
	LoginTime   time.Time `json:"loginTime"`   // 操作时间
}

// tableName 获取当前表的名称
func (e *LoginLog) tableName() string {
	return "sys_login_log"
}

// Get LoginLog 信息获取
func (e *LoginLog) Get() (LoginLogs []LoginLog, err error) {
	table := global.DB.Table(e.tableName())

	if e.ID > 0 {
		table = table.Where("id = ?", e.ID)
	}

	if err = table.Find(&LoginLogs).Error; err != nil {
		return
	}
	return
}

// Create LoginLog 创建
func (e *LoginLog) Create() (LoginLog LoginLog, err error) {
	doc := LoginLog
	table := global.DB.Table(e.tableName())
	result := table.Create(&e)
	if result.Error != nil {
		err = result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}
