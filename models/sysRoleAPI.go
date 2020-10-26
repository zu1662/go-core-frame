package models

import (
	"errors"
	"go-core-frame/global"
)

// SysRoleAPI 角色接口关联结构
type SysRoleAPI struct {
	ID     int `gorm:"" json:"id"`              //数据库id
	RoleID int `json:"roleId" valid:"required"` // 角色ID
	APIID  int `json:"apiId" valid:"required"`  // 接口ID
	BaseModel
}

// SysRoleAPIView 角色接口列表关联结构
type SysRoleAPIView struct {
	RoleID  int   `json:"roleId" valid:"required"`  // 角色ID
	APIList []int `json:"apiList" valid:"required"` // 接口ID
	BaseModel
}

// tableName 获取当前表的名称
func (e *SysRoleAPI) tableName() string {
	return "sys_role_api"
}

// tableName 获取当前表的名称
func (e *SysRoleAPIView) tableName() string {
	return "sys_role_api"
}

// GetRoleAPI Role 关联的 api List列表信息
func (e *SysRoleAPI) GetRoleAPI() ([]SysRoleAPI, error) {
	var doc []SysRoleAPI

	table := global.DB.Table(e.tableName())

	if e.RoleID == 0 {
		return nil, errors.New("roleId不能为空")
	}
	table = table.Where("role_id = ?", e.RoleID)
	if e.APIID > 0 {
		table = table.Where("api_id = ?", e.APIID)
	}

	table = table.Where("is_deleted = ?", 0)

	err := table.Find(&doc).Error
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// UpdateRoleAPI 修改
func (e *SysRoleAPIView) UpdateRoleAPI() (err error) {
	table := global.DB.Table(e.tableName())
	// 先删除之前的roleId关联的数据
	if e.RoleID > 0 {
		if err = table.Where("role_id = ?", e.RoleID).Delete(SysRoleAPI{}).Error; err != nil {
			return
		}
	}
	// 添加新的menu关联数据
	for _, apiID := range e.APIList {
		var sysRoleAPI = SysRoleAPI{
			RoleID:    e.RoleID,
			APIID:     apiID,
			BaseModel: e.BaseModel,
		}
		//添加数据
		if err = table.Create(&sysRoleAPI).Error; err != nil {
			return
		}
	}
	return
}
