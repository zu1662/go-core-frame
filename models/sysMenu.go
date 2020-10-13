package models

import (
	"errors"
	"go-core-frame/global"
)

// SysMenu 菜单信息结构
type SysMenu struct {
	ID         int    `gorm:"" json:"id"`                  //数据库id
	Pid        int    `json:"pid"`                         // 父级id
	Name       string `json:"name" valid:"required"`       // 名称
	Title      string `json:"title" valid:"required"`      // 标题
	Icon       string `json:"icon"`                        // 图标
	Path       string `json:"path"`                        // 路径
	Component  string `json:"component"`                   // 组件地址
	Permission string `json:"permission" valid:"required"` // 权限
	Visible    string `json:"visible"`                     // 可见
	Cache      string `json:"cache"`                       // 缓存
	Type       string `json:"type"`                        // 类型 0目录1菜单2按钮
	Sort       string `json:"sort"`                        //排序
	BaseModel
}

// SysMenuView 菜单树结构
type SysMenuView struct {
	SysMenu
	Children []SysMenuView `json:"children" gorm:"-"` // 子级
}

// tableName 获取当前表的名称
func (e *SysMenu) tableName() string {
	return "sys_menu"
}

// GetMenu 获取菜单详情数据
func (e *SysMenu) GetMenu() (sysMenu SysMenu, err error) {
	table := global.DB.Table(e.tableName())
	if e.ID > 0 {
		table = table.Where("id = ?", e.ID)
	}

	table = table.Where("is_deleted = ?", 0)

	if err = table.First(&sysMenu).Error; err != nil {
		return
	}
	return
}

//UpdateMenu menu 修改
func (e *SysMenu) UpdateMenu() (update SysMenu, err error) {
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

// InsertMenu 添加菜单
func (e *SysMenu) InsertMenu() (id int, err error) {
	table := global.DB.Table(e.tableName())
	// check 用户名
	var count int64
	table.Where("name = ?", e.Name).Count(&count)
	if count > 0 {
		err = errors.New("菜单名称已存在")
		return
	}
	table.Where("title = ?", e.Title).Count(&count)
	if count > 0 {
		err = errors.New("菜单标题已存在")
		return
	}
	table.Where("permission = ?", e.Permission).Count(&count)
	if count > 0 {
		err = errors.New("菜单权限已存在")
		return
	}

	//添加数据
	if err = table.Create(&e).Error; err != nil {
		return
	}
	id = e.ID
	return
}

// DeleteMenu 逻辑删除
func (e *SysMenu) DeleteMenu() (err error) {
	table := global.DB.Table(e.tableName())
	err = table.Where("id = ?", e.ID).Update("is_deleted", 1).Error
	return
}

// GetMenuTree  部门树结构信息
func (e *SysMenuView) GetMenuTree() ([]SysMenuView, error) {
	var doc []SysMenuView
	var docView []SysMenuView
	table := global.DB.Table(e.SysMenu.tableName())

	if e.Title != "" {
		table = table.Where("title LIKE ?", "%"+e.Title+"%")
	}

	table = table.Where("is_deleted = ?", 0)

	err := table.Order("sort").Find(&doc).Error
	if err != nil {
		return nil, err
	}

	for _, nowMenu := range doc {
		if nowMenu.Pid != 0 {
			continue
		}
		newMenu := recursionMenu(&doc, nowMenu)
		docView = append(docView, newMenu)
	}
	return docView, nil
}

// recursion 递归树结构
func recursionMenu(deptList *[]SysMenuView, nowMenu SysMenuView) SysMenuView {
	for _, dept := range *deptList {
		if dept.Pid == nowMenu.ID {
			newMenu := recursionMenu(deptList, dept)
			nowMenu.Children = append(nowMenu.Children, newMenu)
		} else {
			continue
		}
	}
	return nowMenu
}
