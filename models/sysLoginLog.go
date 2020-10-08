package models

import (
	"errors"
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

// GetDetail LoginLog 信息详情
func (e *LoginLog) GetDetail() (LoginLogs LoginLog, err error) {
	table := global.DB.Table(e.tableName())

	if e.ID > 0 {
		table = table.Where("id = ?", e.ID)
	}

	table = table.Where("is_deleted = ?", 0)

	if err = table.First(&LoginLogs).Error; err != nil {
		err = errors.New("获取详情信息失败")
		return
	}
	return
}

// GetPage LoginLog List列表信息
func (e *LoginLog) GetPage(pageSize int, pageIndex int) ([]LoginLog, int64, error) {
	var doc []LoginLog

	table := global.DB.Table(e.tableName())
	if e.IPAddress != "" {
		table = table.Where("ip_address = ?", e.IPAddress)
	}
	if e.UserName != "" {
		table = table.Where("user_name = ?", e.UserName)
	}

	table = table.Where("is_deleted = ?", 0)

	var count int64

	err := table.Order("id desc").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error
	err = table.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	return doc, count, nil
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

// Delete LoginLog 逻辑删除
func (e *LoginLog) Delete() (err error) {
	table := global.DB.Table(e.tableName())
	err = table.Where("id = ?", e.ID).Update("is_deleted", 1).Error
	return
}
