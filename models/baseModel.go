package models

import "time"

type BaseModel struct {
	CreateBy   string    `json:"createBy"`   // 创建人
	CreateTime time.Time `json:"createTime"` // 创建时间
	UpdateBy   string    `json:"updateBy"`   // 修改人
	UpdateTime time.Time `json:"updateTime"` // 修改时间
}
