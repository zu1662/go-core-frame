package models

import (
	"errors"
	"go-core-frame/global"
)

// LoginForm 登录结构
type LoginForm struct {
	Username    string `json:"username" valid:"required"`
	Password    string `json:"password" valid:"required"`
	CaptchaID   string `json:"captchaId"`
	CaptchaCode string `json:"captchaCode"`
}

// SysUser 用户信息结构
type SysUser struct {
	ID       int    `gorm:"" json:"id"`                //数据库id
	UUID     string `json:"uuid"`                      // UUID
	UserCode string `json:"userCode" valid:"required"` // 编码
	UserName string `json:"userName" valid:"required"` // 名称
	Mobile   string `json:"mobile" valid:"required"`   // 手机号
	Avatar   string `json:"avatar"`                    //头像
	Gender   string `json:"gender"`                    //性别
	Email    string `json:"email"`                     //邮箱
	RoleID   int    `json:"roleId"`                    // 角色编码
	DeptID   int    `json:"deptId"`                    //部门编码
	PostID   int    `json:"postId"`                    //职位编码
	Status   string `json:"status"`                    // 状态
	BaseModel
}

// SysUserWithPsw 包含密码的用户信息
type SysUserWithPsw struct {
	SysUser
	Password string `json:"password"` // 密码
}

// SysUserView 用户冗余信息体
type SysUserView struct {
	SysUser
	RoleName string `json:"roleName"` //角色名称
	DeptName string `json:"deptName"` // 部门名称
	PostName string `json:"postName"` // 岗位名称
}

// tableName 获取当前表的名称
func (e *SysUser) tableName() string {
	return "sys_user"
}

// GetUser 获取用户数据
func (e *SysUserWithPsw) GetUser() (sysUser SysUserWithPsw, err error) {
	table := global.DB.Table(e.tableName())

	if e.ID > 0 {
		table = table.Where("id = ?", e.ID)
	}

	if e.UserName != "" {
		table = table.Where("user_code = ?", e.UserName)
	}

	if e.Password != "" {
		table = table.Where("password = ?", e.Password)
	}

	table = table.Where("is_deleted = ?", 0)

	if err = table.First(&sysUser).Error; err != nil {
		return
	}
	return
}

// GetUser 获取用户不带密码详情
func (e *SysUser) GetUser() (sysUser SysUser, err error) {
	table := global.DB.Table(e.tableName())

	if e.ID > 0 {
		table = table.Where("id = ?", e.ID)
	}

	if e.UserName != "" {
		table = table.Where("user_code = ?", e.UserName)
	}

	table = table.Where("is_deleted = ?", 0)

	if err = table.First(&sysUser).Error; err != nil {
		return
	}
	return
}

// GetPage SysUser List列表信息
func (e *SysUserView) GetPage(pageSize int, pageIndex int) ([]SysUserView, int64, error) {
	var doc []SysUserView

	table := global.DB.Table(e.tableName())
	table.Select([]string{"sys_user.*", "sys_role.role_name", "sys_dept.dept_name", "sys_post.post_name"})
	table.Joins("left outer join sys_role on sys_user.role_id=sys_role.id")
	table.Joins("left outer join sys_post on sys_user.post_id=sys_post.id")
	table.Joins("left outer join sys_dept on sys_user.dept_id=sys_dept.id")

	if e.UserName != "" {
		table = table.Where("sys_user.user_name LIKE ?", "%"+e.UserName+"%")
	}

	if e.UserCode != "" {
		table = table.Where("sys_user.user_code LIKE ?", "%"+e.UserCode+"%")
	}

	if e.Mobile != "" {
		table = table.Where("sys_user.mobile LIKE ?", "%"+e.Mobile+"%")
	}

	if e.Status != "" {
		table = table.Where("sys_user.status = ?", e.Status)
	}

	if e.DeptID > 0 {
		table = table.Where("sys_user.dept_id = ?", e.DeptID)
	}

	table = table.Where("sys_user.is_deleted = ?", 0)

	var count int64

	err := table.Order("sys_user.id").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error
	err = table.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	return doc, count, nil
}

//UpdateUser SysUser 修改
func (e *SysUser) UpdateUser() (update SysUser, err error) {
	table := global.DB.Table(e.tableName())
	if err = table.First(&update, e.ID).Error; err != nil {
		return
	}
	if e.RoleID == 0 {
		e.RoleID = update.RoleID
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = table.Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

//InsertUser 添加用户
func (e *SysUser) InsertUser() (id int, err error) {
	table := global.DB.Table(e.tableName())
	// check 用户名
	var count int64
	table.Where("is_deleted = ?", 0)
	table.Where("user_name = ?", e.UserName).Count(&count)
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

// DeleteUser SysUser 逻辑删除
func (e *SysUser) DeleteUser(userIds []int) (err error) {
	table := global.DB.Table(e.tableName())
	err = table.Where("id in (?)", userIds).Update("is_deleted", 1).Error
	return
}

// SysUserInfo 用户信息详情，登录返回给前端
type SysUserInfo struct {
	UUID     string `json:"uuid"`     // UUID
	UserCode string `json:"userCode"` // 编码
	UserName string `json:"userName"` // 名称
	Mobile   string `json:"mobile"`   // 手机号
	Email    string `json:"email"`    //邮箱
	Avatar   string `json:"avatar"`   //头像
	Gender   string `json:"gender"`   //性别
	RoleName string `json:"roleName"` //角色
	DeptName string `json:"deptName"` //部门
	PostName string `json:"postName"` //岗位
	BaseModel
}

// tableName 获取当前表的名称
func (e *SysUserInfo) tableName() string {
	return "sys_user"
}

// GetUserInfo 获取用户的具体信息
func (e *SysUserInfo) GetUserInfo() (SysUserInfo SysUserInfo, err error) {
	table := global.DB.Table(e.tableName())
	table.Select([]string{"sys_user.uuid", "sys_user.user_name", "sys_user.user_code", "sys_user.avatar", "sys_user.mobile", "sys_user.gender", "sys_role.role_name", "sys_dept.dept_name", "sys_post.post_name"})
	table.Joins("left outer join sys_role on sys_user.role_id=sys_role.id")
	table.Joins("left outer join sys_post on sys_user.post_id=sys_post.id")
	table.Joins("left outer join sys_dept on sys_user.dept_id=sys_dept.id")

	if e.UserCode != "" {
		table.Where("user_code = ?", e.UserCode)
	}

	if e.UserName != "" {
		table.Where("user_name = ?", e.UserName)
	}

	if e.Mobile != "" {
		table.Where("mobile = ?", e.Mobile)
	}

	if err = table.First(&SysUserInfo).Error; err != nil {
		return
	}
	return
}
