package models

import (
	"errors"
	"go-core-frame/global"
)

// SysPost 岗位信息结构
type SysPost struct {
	ID          int    `gorm:"" json:"id"`                //数据库id
	PostCode    string `json:"postCode" valid:"required"` // 编码
	PostName    string `json:"postName" valid:"required"` // 名称
	Description string `json:"description"`               // 描述
	Sort        int    `json:"sort"`                      //排序
	Status      string `json:"status"`                    // 状态
	BaseModel
}

// tableName 获取当前表的名称
func (e *SysPost) tableName() string {
	return "sys_post"
}

// GetPost 获取岗位数据
func (e *SysPost) GetPost() (sysPost SysPost, err error) {
	table := global.DB.Table(e.tableName())

	if e.ID > 0 {
		table = table.Where("id = ?", e.ID)
	}

	if e.PostCode != "" {
		table = table.Where("post_code = ?", e.PostCode)
	}

	if e.PostName != "" {
		table = table.Where("post_name = ?", e.PostName)
	}

	table = table.Where("is_deleted = ?", 0)

	if err = table.First(&sysPost).Error; err != nil {
		return
	}
	return
}

// GetPage List列表信息
func (e *SysPost) GetPage(pageSize int, pageIndex int) ([]SysPost, int64, error) {
	var doc []SysPost

	table := global.DB.Table(e.tableName())

	if e.PostCode != "" {
		table = table.Where("post_code LIKE ?", "%"+e.PostCode+"%")
	}

	if e.PostName != "" {
		table = table.Where("post_name LIKE ?", "%"+e.PostName+"%")
	}

	if e.Status != "" {
		table = table.Where("status = ?", e.Status)
	}

	table = table.Where("is_deleted = ?", 0)

	var count int64
	err := table.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = table.Order("sort").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error
	if err != nil {
		return nil, 0, err
	}
	return doc, count, nil
}

// GetAll List列表信息
func (e *SysPost) GetAll() ([]SysPost, error) {
	var doc []SysPost

	table := global.DB.Table(e.tableName())

	if e.PostCode != "" {
		table = table.Where("post_code LIKE ?", "%"+e.PostCode+"%")
	}

	if e.PostName != "" {
		table = table.Where("post_name LIKE ?", "%"+e.PostName+"%")
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

// UpdatePost 岗位修改
func (e *SysPost) UpdatePost() (update SysPost, err error) {
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

//InsertPost 添加岗位
func (e *SysPost) InsertPost() (id int, err error) {
	table := global.DB.Table(e.tableName())
	// check 用户名
	var count int64
	table = table.Where("is_deleted = ?", 0)
	table.Where("post_code = ?", e.PostCode).Count(&count)
	if count > 0 {
		err = errors.New("账户已存在！")
		return
	}

	//添加数据
	if err = table.Create(&e).Error; err != nil {
		return
	}
	id = e.ID
	return
}

// DeletePost 逻辑删除
func (e *SysPost) DeletePost(ids []int) (err error) {
	table := global.DB.Table(e.tableName())
	err = table.Where("id in (?)", ids).Update("is_deleted", 1).Error
	return
}
