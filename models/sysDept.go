package models

import (
	"errors"
	"go-core-frame/global"
)

// SysDept 部门信息结构
type SysDept struct {
	ID       int    `gorm:"" json:"id"`                //数据库id
	DeptName string `json:"deptName" valid:"required"` // 部门名称
	Pid      int    `json:"pid"`                       // 父级id
	LeaderID int    `json:"leaderId" valid:"required"` // 负责人
	Sort     string `json:"sort"`                      //排序
	Status   string `json:"status"`                    // 状态
	BaseModel
}

// SysDeptView 部门树状结构体
type SysDeptView struct {
	SysDept
	LeaderName   string        `json:"leaderName"`        // 负责人名称
	LeaderMobile string        `json:"leaderMobile"`      // 负责人手机
	LeaderEmail  string        `json:"leaderEmail"`       // 负责人邮箱
	Children     []SysDeptView `json:"children" gorm:"-"` // 子集组织
}

// tableName 获取当前表的名称
func (e *SysDept) tableName() string {
	return "sys_dept"
}

// GetDept 获取部门详情数据
func (e *SysDept) GetDept() (sysDept SysDept, err error) {
	table := global.DB.Table(e.tableName())
	if e.ID > 0 {
		table = table.Where("id = ?", e.ID)
	}

	if e.DeptName != "" {
		table = table.Where("dept_name = ?", e.DeptName)
	}

	if e.Status != "" {
		table = table.Where("status = ?", e.Status)
	}

	table = table.Where("is_deleted = ?", 0)

	if err = table.First(&sysDept).Error; err != nil {
		return
	}
	return
}

//UpdateDept dept 修改
func (e *SysDept) UpdateDept() (update SysDept, err error) {
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

// InsertDept 添加部门
func (e *SysDept) InsertDept() (id int, err error) {
	table := global.DB.Table(e.tableName())
	// check 用户名
	var count int64
	table.Where("dept_name = ?", e.DeptName).Count(&count)
	if count > 0 {
		err = errors.New("部门已存在！")
		return
	}

	//添加数据
	if err = table.Create(&e).Error; err != nil {
		return
	}
	id = e.ID
	return
}

// DeleteDept 逻辑删除
func (e *SysDept) DeleteDept() (err error) {
	table := global.DB.Table(e.tableName())
	err = table.Where("id = ?", e.ID).Update("is_deleted", 1).Error
	return
}

// GetDeptTree  部门树结构信息
func (e *SysDeptView) GetDeptTree() ([]SysDeptView, error) {
	var doc []SysDeptView
	var docAll []SysDeptView
	var docView []SysDeptView
	table := global.DB.Table(e.SysDept.tableName())

	// 预设join信息
	table.Select([]string{"sys_dept.*", "sys_user.user_name as leader_name", "sys_user.mobile as leader_mobile", "sys_user.email as leader_email"})
	table.Joins("left outer join sys_user on sys_dept.leader_id=sys_user.id")

	// 全部的信息
	table = table.Where("sys_dept.is_deleted = ?", 0)
	err := table.Order("sys_dept.sort").Find(&docAll).Error
	if err != nil {
		return nil, err
	}

	//搜索条件的信息
	if e.ID > 0 {
		table = table.Where("sys_dept.id = ?", e.ID)
	}
	table = table.Where("sys_dept.is_deleted = ?", 0)
	err = table.Order("sys_dept.sort").Find(&doc).Error
	if err != nil {
		return nil, err
	}

	for _, nowDept := range doc {
		if e.ID == 0 && nowDept.Pid != 0 {
			continue
		}
		newDept := recursionDept(&docAll, nowDept)
		docView = append(docView, newDept)
	}
	return docView, nil
}

// recursion 递归树结构
func recursionDept(deptList *[]SysDeptView, nowDept SysDeptView) SysDeptView {
	for _, dept := range *deptList {
		if dept.Pid == nowDept.ID {
			newDept := recursionDept(deptList, dept)
			nowDept.Children = append(nowDept.Children, newDept)
		} else {
			continue
		}
	}
	return nowDept
}
