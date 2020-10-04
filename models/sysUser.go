package models

import "go-core-frame/global"

// LoginForm 登录结构
type LoginForm struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}

// SysUser 用户信息结构
type SysUser struct {
	ID       int    `gorm:"" json:"id"` //数据库id
	UUID     string `json:"uuid"`       // UUID
	UserCode string `json:"userCode"`   // 编码
	UserName string `json:"userName"`   // 名称
	Password string `json:"password"`   // 密码
	Mobile   string `json:"mobile"`     // 手机号
	Avatar   string `json:"avatar"`     //头像
	Gender   string `json:"gender"`     //性别
	Email    string `json:"email"`      //邮箱
	RoleID   int    `json:"roleId"`     // 角色编码
	DeptID   int    `json:"deptId"`     //部门编码
	PostID   int    `json:"postId"`     //职位编码
	Status   string `json:"status"`     // 状态
	BaseModel
}

// tableName 获取当前表的名称
func (e *SysUser) tableName() string {
	return "sys_user"
}

// Get 获取用户数据
func (e *SysUser) Get() (sysUser SysUser, err error) {
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

	if err = table.First(&sysUser).Error; err != nil {
		return
	}
	return
}
