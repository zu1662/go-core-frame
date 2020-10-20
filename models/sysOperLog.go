package models

import (
	"errors"
	"go-core-frame/global"
	"time"
)

// OperLog 操作日志Model
type OperLog struct {
	ID          int       `gorm:"" json:"id"`  // 数据库id
	OperName    string    `json:"operName"`    // 操作人
	OperTitle   string    `json:"operTitle"`   // 操作名称
	Method      string    `json:"method"`      // 操作类型
	Path        string    `json:"path"`        // 地址
	Params      string    `json:"params"`      // 参数
	LatencyTime string    `json:"latencyTime"` // 执行时间
	IPAddress   string    `json:"ipAddress"`   // IP地址
	IPLocation  string    `json:"ipLocation"`  // IP归属地
	Browser     string    `json:"browser"`     // 浏览器
	OS          string    `json:"os"`          // 操作系统
	Result      string    `json:"result"`      // 操作结果
	OperTime    time.Time `json:"operTime"`    // 操作时间
}

// tableName 获取当前表的名称
func (e *OperLog) tableName() string {
	return "sys_oper_log"
}

// GetDetail OperLog 信息详情
func (e *OperLog) GetDetail() (OperLogs OperLog, err error) {
	table := global.DB.Table(e.tableName())

	if e.ID > 0 {
		table = table.Where("id = ?", e.ID)
	}

	table = table.Where("is_deleted = ?", 0)

	if err = table.First(&OperLogs).Error; err != nil {
		err = errors.New("获取详情信息失败")
		return
	}
	return
}

// GetPage OperLog List列表信息
func (e *OperLog) GetPage(pageSize int, pageIndex int) ([]OperLog, int64, error) {
	var doc []OperLog

	table := global.DB.Table(e.tableName())
	if e.IPAddress != "" {
		table = table.Where("ip_address LIKE ?", "%"+e.IPAddress+"%")
	}
	if e.Method != "" {
		table = table.Where("method LIKE ?", "%"+e.Method+"%")
	}
	if e.OperName != "" {
		table = table.Where("oper_name LIKE ?", "%"+e.OperName+"%")
	}

	table = table.Where("is_deleted = ?", 0)

	var count int64
	err := table.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = table.Order("id desc").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error
	if err != nil {
		return nil, 0, err
	}
	return doc, count, nil
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

// Delete OperLog 逻辑删除
func (e *OperLog) Delete(ids []int) (err error) {
	table := global.DB.Table(e.tableName())
	err = table.Where("id in (?)", ids).Update("is_deleted", 1).Error
	return
}

// Clean OperLog 逻辑清空
func (e *OperLog) Clean() (err error) {
	table := global.DB.Table(e.tableName())
	err = table.Where("is_deleted = ?", 0).Update("is_deleted", 1).Error
	return
}
