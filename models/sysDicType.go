package models

import (
	"errors"
	"go-core-frame/global"
)

// SysDictType 字典类型信息结构
type SysDictType struct {
	ID       int    `gorm:"" json:"id"`                //数据库id
	DictName string `json:"dictName" valid:"required"` // 字典名称
	DictType string `json:"dictType" valid:"required"` // 字典类型（编码）
	Sort     string `json:"sort"`                      // 状态
	Status   string `json:"status"`                    // 状态
	BaseModel
}

// tableName 获取当前表的名称
func (e *SysDictType) tableName() string {
	return "sys_dict_type"
}

// GetDictType 获取详情
func (e *SysDictType) GetDictType() (sysDictType SysDictType, err error) {
	table := global.DB.Table(e.tableName())

	if e.ID > 0 {
		table = table.Where("id = ?", e.ID)
	}

	if e.DictName != "" {
		table = table.Where("dict_name = ?", e.DictName)
	}

	if e.DictType != "" {
		table = table.Where("dict_type = ?", e.DictType)
	}

	table = table.Where("is_deleted = ?", 0)

	if err = table.First(&sysDictType).Error; err != nil {
		return
	}
	return
}

// GetDictTypePage List列表信息
func (e *SysDictType) GetDictTypePage(pageSize int, pageIndex int) ([]SysDictType, int64, error) {
	var doc []SysDictType

	table := global.DB.Table(e.tableName())

	if e.DictName != "" {
		table = table.Where("dict_name LIKE ?", "%"+e.DictName+"%")
	}

	if e.DictType != "" {
		table = table.Where("dict_type LIKE ?", "%"+e.DictType+"%")
	}

	if e.Status != "" {
		table = table.Where("status = ?", e.Status)
	}

	table = table.Where("is_deleted = ?", 0)

	var count int64

	err := table.Order("sort").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error
	err = table.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	return doc, count, nil
}

// UpdateDictType 修改
func (e *SysDictType) UpdateDictType() (update SysDictType, err error) {
	table := global.DB.Table(e.tableName())
	if err = table.First(&update, e.ID).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = table.Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

//InsertDictType 添加
func (e *SysDictType) InsertDictType() (id int, err error) {
	table := global.DB.Table(e.tableName())
	// check 用户名
	var count int64
	table.Where("dict_type = ?", e.DictType).Count(&count)
	if count > 0 {
		err = errors.New("字典值已存在！")
		return
	}

	//添加数据
	if err = table.Create(&e).Error; err != nil {
		return
	}
	id = e.ID
	return
}

// DeleteDictType 逻辑删除
func (e *SysDictType) DeleteDictType() (err error) {
	table := global.DB.Table(e.tableName())
	err = table.Where("id = ?", e.ID).Update("is_deleted", 1).Error
	return
}

// GetDictTypeAll 所有dictType信息
func (e *SysDictType) GetDictTypeAll() ([]SysDictType, error) {
	var doc []SysDictType

	table := global.DB.Table(e.tableName())

	if e.Status != "" {
		table = table.Where("status = ?", e.Status)
	}

	table = table.Where("is_deleted = ?", 0)

	err := table.Order("sort").Find(&doc).Error
	if err != nil {
		return nil, err
	}
	return doc, nil
}
