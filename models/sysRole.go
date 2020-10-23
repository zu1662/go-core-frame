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
	Sort        int    `json:"sort"`                      //排序
	Status      string `json:"status"`                    // 状态
	BaseModel
}

// SysRoleView 角色返回结构
type SysRoleView struct {
	SysRole
	MenuList []int `json:"menuList"` // 关联菜单信息
	APIList  []int `json:"apiList"`  // 关联api信息
}

// tableName 获取当前表的名称
func (e *SysRole) tableName() string {
	return "sys_role"
}

// GetRole 获取角色详情数据
func (e *SysRoleView) GetRole() (sysRoleView SysRoleView, err error) {
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

	if err = table.First(&sysRoleView).Error; err != nil {
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

// GetList 所有role信息
func (e *SysRole) GetList() ([]SysRole, error) {
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

	err := table.Order("sort").Find(&doc).Error
	if err != nil {
		return nil, err
	}
	return doc, nil
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
	table = table.Where("is_deleted = ?", 0)
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
func (e *SysRole) DeleteRole(ids []int) (err error) {
	table := global.DB.Table(e.tableName())
	err = table.Where("id in (?)", ids).Update("is_deleted", 1).Error
	return
}
