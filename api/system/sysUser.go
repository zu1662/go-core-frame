package system

import (
	"go-core-frame/global"
	"go-core-frame/models"
	"go-core-frame/pkg/app"
	"go-core-frame/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Login 用户登录
// @Tags user
// @Summary 用户登录
// @Produce  application/json
// @Param data body models.LoginForm true "用户登录接口"
// @Success 200 {string} string "{"code":200,"data":{"token": "token value", "expire": "expire time"},"msg":"ok"}"
// @Router /user/login [post]
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
		"msg":  "登陆成功",
		"data": token,
	})
}

// Logout 用户退出
// @Tags user
// @Summary 用户退出
// @Produce  application/json
// @Success 200 {string} string "{"code":200,"data":null,"msg":"ok"}"
// @Router /user/logout [post]
// @Security Authorization
func Logout(c *gin.Context) {

	username, _ := c.Get("username")

	err := global.Redis.Do("DEL", username).Err()
	if err != nil {
		utils.HasError(err, "退出失败", 0)
	}

	app.Custom(c, gin.H{
		"code": 200,
		"msg":  "退出成功",
	})

}

// GetSysUser 获取用户详情信息
// @Tags user
// @Summary 获取用户详情信息
// @Produce  application/json
// @Param userId int true "用户编码"
// @Success 200 {object} app.Response "{"code": 200, "data": [...], "msg": "ok"}"
// @Router /user/{userId} [get]
// @Security Authorization
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
