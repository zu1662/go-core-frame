package system

import (
	"go-core-frame/global"
	"go-core-frame/models"
	"go-core-frame/pkg/app"
	"go-core-frame/pkg/config"
	"go-core-frame/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Login 用户登录
// @Tags base
// @Summary 用户登录
// @Produce  application/json
// @Param data body models.LoginForm true "用户登录接口"
// @Success 200 {string} string "{"code":200,"data":{"token": "token value", "expire": "expire time"},"msg":"ok"}"
// @Router /base/login [post]
func Login(c *gin.Context) {
	var data models.LoginForm
	err := c.ShouldBindWith(&data, binding.JSON)
	utils.HasError(err, "抱歉未找到相关信息", 0)

	// 设置 username 便于 logger 使用
	c.Set("username", data.Username)

	errValidate := utils.StructValidate(data)
	utils.HasError(errValidate, "", 0)

	// 用户密码 再次加密
	data.Password = utils.GetSHA256HashCode([]byte(data.Password))

	sysUser := models.SysUser{}
	sysUser.UserName = data.Username
	sysUser.Password = data.Password
	loginUser, getErr := sysUser.GetUser()
	utils.HasError(getErr, "用户名或密码不正确", 0)

	j := models.NewJWT()
	token, tokenErr := j.CreateToken(&models.UserClaims{
		UUID:     loginUser.UUID,
		Usercode: loginUser.UserCode,
		Username: loginUser.UserName,
	})
	utils.HasError(tokenErr, "创建token失败", 0)
	app.Custom(c, gin.H{
		"code": 1,
		"msg":  "登陆成功",
		"data": token,
	})
}

// Logout 用户退出
// @Tags base
// @Summary 用户退出
// @Produce  application/json
// @Success 200 {string} string "{"code":200,"data":null,"msg":"ok"}"
// @Router /base/logout [post]
// @Security Authorization
func Logout(c *gin.Context) {

	token := c.Request.Header.Get(config.JWTConfig.HeaderName)

	err := global.Redis.SAdd("tokenBlock", token).Err()
	if err != nil {
		utils.HasError(err, "退出失败", 0)
	}

	app.Custom(c, gin.H{
		"code": 1,
		"msg":  "退出成功",
	})
}

// GetUserInfo 获取当前登录用户详情信息
// @Tags base
// @Summary 获取当前登录用户详情信息
// @Produce  application/json
// @Success 200 {object} app.Response "{"code": 200, "data": {...}, "msg": "ok"}"
// @Router /base/getuserinfo [get]
// @Security Authorization
func GetUserInfo(c *gin.Context) {

	username, _ := c.Get("username")

	userView := models.SysUserInfo{}
	userView.UserCode = username.(string)
	userInfo, err := userView.GetUserInfo()
	if err != nil {
		utils.HasError(err, "", 0)
	}

	userInfo.Mobile = utils.MobileSecurity(userInfo.Mobile)

	app.Custom(c, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": userInfo,
	})
}
