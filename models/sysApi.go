package models

import (
	"errors"
	"go-core-frame/global"
)

// SysAPI API信息结构
type SysAPI struct {
	ID     int    `gorm:"" json:"id"`            //数据库id
	Pid    int    `json:"pid" valid:"required"`  // 父级id
	Name   string `json:"name" valid:"required"` // 名称
	Path   string `json:"path" valid:"required"` // 路径
	Method string `json:"method"`                // 请求类型
	Type   string `json:"type" valid:"required"` // 类型 0目录1接口
	Sort   int    `json:"sort"`                  //排序
	Status string `json:"status"`                // 状态
	BaseModel
}

// SysAPIView API树结构
type SysAPIView struct {
	SysAPI
	Children []SysAPIView `json:"children" gorm:"-"` // 子级
}

// tableName 获取当前表的名称
func (e *SysAPI) tableName() string {
	return "sys_api"
}

// GetAPI 获取详情数据
func (e *SysAPI) GetAPI() (sysAPI SysAPI, err error) {
	table := global.DB.Table(e.tableName())
	if e.ID > 0 {
		table = table.Where("id = ?", e.ID)
	}
	if e.Path != "" {
		table = table.Where("path LIKE ?", "%"+e.Path+"%")
	}
	if e.Method != "" {
		table = table.Where("method = ?", e.Method)
	}

	table = table.Where("is_deleted = ?", 0)

	if err = table.First(&sysAPI).Error; err != nil {
		return
	}
	return
}

//UpdateAPI 修改
func (e *SysAPI) UpdateAPI() (update SysAPI, err error) {
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

// InsertAPI 添加
func (e *SysAPI) InsertAPI() (id int, err error) {
	table := global.DB.Table(e.tableName())
	// check 用户名
	var count int64
	table = table.Where("is_deleted = ?", 0)
	table.Where("name = ?", e.Name).Count(&count)
	if count > 0 {
		err = errors.New("API名称已存在")
		return
	}
	table.Where("path = ?", e.Path).Count(&count)
	if count > 0 {
		err = errors.New("API路径已存在")
		return
	}

	//添加数据
	if err = table.Create(&e).Error; err != nil {
		return
	}
	id = e.ID
	return
}

// DeleteAPI 逻辑删除
func (e *SysAPI) DeleteAPI() (err error) {
	table := global.DB.Table(e.tableName())
	err = table.Where("id = ?", e.ID).Update("is_deleted", 1).Error
	return
}

// GetAPITree  部门树结构信息
func (e *SysAPIView) GetAPITree() ([]SysAPIView, error) {
	var doc []SysAPIView
	var docAll []SysAPIView
	var docView []SysAPIView
	table := global.DB.Table(e.SysAPI.tableName())

	// 全部信息
	table = table.Where("is_deleted = ?", 0)
	err := table.Order("sort").Find(&docAll).Error
	if err != nil {
		return nil, err
	}

	// 搜索条件信息
	if e.ID > 0 {
		table = table.Where("id = ?", e.ID)
	}
	table = table.Where("is_deleted = ?", 0)
	err = table.Order("sort").Find(&doc).Error
	if err != nil {
		return nil, err
	}

	for _, nowAPI := range doc {
		if e.ID == 0 && nowAPI.Pid != 0 {
			continue
		}
		newAPI := recursionAPI(&docAll, nowAPI)
		docView = append(docView, newAPI)
	}
	return docView, nil
}

// recursion 递归树结构
func recursionAPI(apiList *[]SysAPIView, nowAPI SysAPIView) SysAPIView {
	for _, api := range *apiList {
		if api.Pid == nowAPI.ID {
			newAPI := recursionAPI(apiList, api)
			nowAPI.Children = append(nowAPI.Children, newAPI)
		} else {
			continue
		}
	}
	return nowAPI
}
