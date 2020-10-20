package models

import (
	"errors"
	"go-core-frame/global"
)

// SysDictData 字典类型信息结构
type SysDictData struct {
	ID          int    `gorm:"" json:"id"`                  //数据库id
	DictTypeID  string `json:"dictTypeId" valid:"required"` // 字典类型ID（编码）
	DictLabel   string `json:"dictLabel" valid:"required"`  // 字典值字段
	DictValue   string `json:"dictValue" valid:"required"`  // 字典值字段
	Description string `json:"discription"`                 // 描述
	Sort        string `json:"sort"`                        //排序
	Status      string `json:"status"`                      // 状态
	BaseModel
}

// tableName 获取当前表的名称
func (e *SysDictData) tableName() string {
	return "sys_dict_data"
}

// GetDictData 获取详情
func (e *SysDictData) GetDictData() (sysDictData SysDictData, err error) {
	table := global.DB.Table(e.tableName())

	if e.ID > 0 {
		table = table.Where("id = ?", e.ID)
	}

	if e.DictTypeID != "" {
		table = table.Where("dict_type_id = ?", e.DictTypeID)
	}

	if e.DictLabel != "" {
		table = table.Where("dict_label = ?", e.DictLabel)
	}

	if e.DictValue != "" {
		table = table.Where("dict_value = ?", e.DictValue)
	}

	table = table.Where("is_deleted = ?", 0)

	if err = table.First(&sysDictData).Error; err != nil {
		return
	}
	return
}

// GetDictDataPage List列表信息
func (e *SysDictData) GetDictDataPage(pageSize int, pageIndex int) ([]SysDictData, int64, error) {
	var doc []SysDictData

	table := global.DB.Table(e.tableName())

	if e.DictTypeID != "" {
		table = table.Where("dict_type_id LIKE ?", "%"+e.DictTypeID+"%")
	}

	if e.DictLabel != "" {
		table = table.Where("dict_label LIKE ?", "%"+e.DictLabel+"%")
	}

	if e.DictValue != "" {
		table = table.Where("dict_value LIKE ?", "%"+e.DictValue+"%")
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

// UpdateDictData 修改
func (e *SysDictData) UpdateDictData() (update SysDictData, err error) {
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

//InsertDictData 添加
func (e *SysDictData) InsertDictData() (id int, err error) {
	table := global.DB.Table(e.tableName())
	// check 用户名
	var count int64
	table = table.Where("is_deleted = ?", 0)
	table.Where("dict_label = ?", e.DictLabel).Count(&count)
	if count > 0 {
		err = errors.New("字典数据值已存在！")
		return
	}

	//添加数据
	if err = table.Create(&e).Error; err != nil {
		return
	}
	id = e.ID
	return
}

// DeleteDictData 逻辑删除
func (e *SysDictData) DeleteDictData() (err error) {
	table := global.DB.Table(e.tableName())
	err = table.Where("id = ?", e.ID).Update("is_deleted", 1).Error
	return
}

// GetDictDataAll List列表信息
func (e *SysDictData) GetDictDataAll() ([]SysDictData, error) {
	var doc []SysDictData

	table := global.DB.Table(e.tableName())

	if e.DictTypeID != "" {
		table = table.Where("dict_type_id = ?", e.DictTypeID)
	}

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
