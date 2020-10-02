package models

import (
	"go-core-frame/global"
	"time"
)

// OperLog 操作日志Model
type OperLog struct {
	ID          int       `gorm:"" json:"id"`  // 数据库id
	OperName    string    `json:"operName"`    // 操作人
	OperTitle   string    `json:"operTitle"`   // 操作名称
	Method      string    `json:"method"`      // 操作类型
	Path        string    `json:"Path"`        // 地址
	Params      string    `json:"params"`      // 参数
	LatencyTime string    `json:"latencyTime"` // 执行时间
	IPAddress   string    `json:"ipAddress"`   // IP地址
	IPLocation  string    `json:"ipLocation"`  // IP归属地
	Browser     string    `json:"browser"`     // 浏览器
	OS          string    `json:"os"`          // 操作系统
	Result      string    `json:"reslut"`      // 操作结果
	OperTime    time.Time `json:"operTime"`    // 操作时间
}

// tableName 获取当前表的名称
func (e *OperLog) tableName() string {
	return "sys_oper_log"
}

// Get OperLog 信息获取
func (e *OperLog) Get() (OperLogs []OperLog, err error) {
	table := global.DB.Table(e.tableName())

	if e.ID > 0 {
		table = table.Where("id = ?", e.ID)
	}

	if err = table.Find(&OperLogs).Error; err != nil {
		return
	}
	return
}

// Create OperLog 创建
func (e *OperLog) Create() (OperLog OperLog, err error) {
	doc := OperLog
	table := global.DB.Table(e.tableName())
	result := table.Create(&e)

	if result.Error != nil {
		err = result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}
