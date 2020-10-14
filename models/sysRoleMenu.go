package models

import (
	"go-core-frame/global"
)

// SysRoleMenu 部门菜单关联结构
type SysRoleMenu struct {
	ID     int `gorm:"" json:"id"`              //数据库id
	RoleID int `json:"roleId" valid:"required"` // 角色ID
	MenuID int `json:"menuId" valid:"required"` // 菜单ID
	BaseModel
}

// SysRoleMenuView 部门菜单列表关联结构
type SysRoleMenuView struct {
	RoleID   int   `json:"roleId" valid:"required"`   // 角色ID
	MenuList []int `json:"menuList" valid:"required"` // 菜单ID
	BaseModel
}

// tableName 获取当前表的名称
func (e *SysRoleMenu) tableName() string {
	return "sys_role_menu"
}

// tableName 获取当前表的名称
func (e *SysRoleMenuView) tableName() string {
	return "sys_role_menu"
}

// GetRoleMenu Role 关联的 Menu List列表信息
func (e *SysRoleMenu) GetRoleMenu() ([]SysRoleMenu, error) {
	var doc []SysRoleMenu

	table := global.DB.Table(e.tableName())

	if e.RoleID > 0 {
		table = table.Where("role_id = ?", e.RoleID)
	}

	if e.MenuID > 0 {
		table = table.Where("menu_id = ?", e.MenuID)
	}

	table = table.Where("is_deleted = ?", 0)

	err := table.Find(&doc).Error
	if err != nil {
		return nil, err
	}
	return doc, nil
}

//UpdateRoleMenu 修改
func (e *SysRoleMenuView) UpdateRoleMenu() (err error) {
	table := global.DB.Table(e.tableName())
	// 先删除之前的roleId关联的数据
	if e.RoleID > 0 {
		if err = table.Where("role_id = ?", e.RoleID).Delete(SysRoleMenu{}).Error; err != nil {
			return
		}
	}
	// 添加新的menu关联数据
	for _, menuID := range e.MenuList {
		var sysRoleMenu = SysRoleMenu{
			RoleID:    e.RoleID,
			MenuID:    menuID,
			BaseModel: e.BaseModel,
		}
		//添加数据
		if err = table.Create(&sysRoleMenu).Error; err != nil {
			return
		}
	}
	return
}
