package models

import "go-core-frame/global"

// API API Model
type API struct {
	ID     int    `gorm:"" json:"id"` // 数据库id
	PID    int    `json:"pid"`        // 父级id
	Name   string `json:"name"`       // 接口名称
	Path   string `json:"path"`       // 接口地址
	Method string `json:"method"`     // 接口方式
	Status int    `json:"status"`     // 接口状态
	BaseModel
}

// tableName 获取当前表的名称
func (e *API) tableName() string {
	return "sys_api"
}

// Get API 信息获取
func (e *API) Get() (Apis []API, err error) {
	table := global.DB.Table(e.tableName())

	if e.ID > 0 {
		table = table.Where("id = ?", e.ID)
	}

	if e.Path != "" {
		table = table.Where("path = ?", e.Path)
	}

	if e.Method != "" {
		table = table.Where("method = ?", e.Method)
	}

	if err = table.Order("sort").Find(&Apis).Error; err != nil {
		return
	}
	return
}
