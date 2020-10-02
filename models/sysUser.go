package models

import (
	"fmt"
	"go-core-frame/global"
)

// SysUser 用户信息结构
type SysUser struct {
	ID       int    `gorm:"" json:"id"` //数据库id
	UserCode string `json:"userCode"`   // 唯一编码
	UserName string `json:"userName"`   // 名称
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

func GetUserInfo(userId int) (SysUser SysUser, err error) {
	global.DB.Table("sys_user").Where("id = ?", "1").First(&SysUser)
	fmt.Println(SysUser)
	return
}
