package models

import (
	"errors"
	"go-core-frame/global"
)

// SysRole 部门信息结构
type SysRole struct {
	ID          int    `gorm:"" json:"id"`                //数据库id
	RoleCode    string `json:"roleCode" valid:"required"` // 角色code
	RoleName    string `json:"roleName" valid:"required"` // 角色名称
	Description string `json:"description"`               // 描述
	Sort        string `json:"sort"`                      //排序
	Status      string `json:"status"`                    // 状态
	BaseModel
}

// tableName 获取当前表的名称
func (e *SysRole) tableName() string {
	return "sys_role"
}

// GetRole 获取角色详情数据
func (e *SysRole) GetRole() (sysRole SysRole, err error) {
	table := global.DB.Table(e.tableName())
	if e.ID > 0 {
		table = table.Where("id = ?", e.ID)
	}

	if e.RoleCode != "" {
		table = table.Where("role_code = ?", e.RoleCode)
	}

	if e.RoleName != "" {
		table = table.Where("role_name = ?", e.RoleName)
	}

	if e.Status != "" {
		table = table.Where("status = ?", e.Status)
	}

	table = table.Where("is_deleted = ?", 0)

	if err = table.First(&sysRole).Error; err != nil {
		return
	}
	return
}

// GetPage List列表信息
func (e *SysRole) GetPage(pageSize int, pageIndex int) ([]SysRole, int64, error) {
	var doc []SysRole

	table := global.DB.Table(e.tableName())

	if e.RoleCode != "" {
		table = table.Where("role_code LIKE ?", "%"+e.RoleCode+"%")
	}

	if e.RoleName != "" {
		table = table.Where("role_name LIKE ?", "%"+e.RoleName+"%")
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

//UpdateRole 修改
func (e *SysRole) UpdateRole() (update SysRole, err error) {
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

// InsertRole 添加角色
func (e *SysRole) InsertRole() (id int, err error) {
	table := global.DB.Table(e.tableName())
	// check 用户名
	var count int64
	table.Where("role_code = ?", e.RoleCode).Count(&count)
	if count > 0 {
		err = errors.New("角色编码已存在！")
		return
	}

	table.Where("role_name = ?", e.RoleName).Count(&count)
	if count > 0 {
		err = errors.New("角色名称已存在！")
		return
	}

	//添加数据
	if err = table.Create(&e).Error; err != nil {
		return
	}
	id = e.ID
	return
}

// DeleteRole 逻辑删除
func (e *SysRole) DeleteRole() (err error) {
	table := global.DB.Table(e.tableName())
	err = table.Where("id = ?", e.ID).Update("is_deleted", 1).Error
	return
}
