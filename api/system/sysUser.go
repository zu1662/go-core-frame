package system

import (
	"go-core-frame/models"
	"go-core-frame/pkg/app"
	"go-core-frame/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Login 用户登录
func Login(c *gin.Context) {
	var data models.LoginForm
	err := c.ShouldBindWith(&data, binding.JSON)
	utils.HasError(err, "抱歉未找到相关信息", 0)

	// 设置 username 便于 logger 使用
	c.Set("username", data.Username)

	errValidate := utils.StructValidate(data)
	utils.HasError(errValidate, "", 0)

	sysUser := models.SysUser{}
	sysUser.UserName = data.Username
	sysUser.Password = data.Password
	loginUser, getErr := sysUser.Get()
	utils.HasError(getErr, "用户名或密码不正确", 0)

	j := models.NewJWT()
	token, tokenErr := j.CreateToken(&models.UserClaims{
		UUID:     loginUser.UUID,
		Usercode: loginUser.UserCode,
		Username: loginUser.UserName,
	})
	utils.HasError(tokenErr, "获取Token失败", 0)
	app.Custom(c, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": token,
	})
}

// GetSysUser 获取用户详情信息
func GetSysUser(c *gin.Context) {
	userID, _ := utils.StringToInt(c.Param("userId"))
	sysUser := models.SysUser{}
	sysUser.ID = userID
	nowUser, err := sysUser.Get()
	utils.HasError(err, "抱歉未找到相关信息", 0)

	app.Custom(c, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": nowUser,
	})
}
