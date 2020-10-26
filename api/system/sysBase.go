package system

import (
	"errors"
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
	var captchaBody models.CaptchaBody
	var data models.LoginForm
	err := c.ShouldBindWith(&data, binding.JSON)
	utils.HasError(err, "抱歉未找到相关信息", 0)

	// 当环境不为dev时，开启captcha校验
	if config.ApplicationConfig.Mode != "dev" {
		if data.CaptchaID == "" {
			err = errors.New("captchaId 不能为空")
		}
		if data.CaptchaCode == "" {
			err = errors.New("captchaCode 不能为空")
		}
		utils.HasError(err, "", 0)
		captchaFlag := captchaBody.VerifyCaptcha(data.CaptchaID, data.CaptchaCode)
		if !captchaFlag {
			err = errors.New("验证码校验失败")
		}
		utils.HasError(err, "", 0)
	}

	// 设置 username 便于 logger 使用
	c.Set("username", data.Username)

	errValidate := utils.StructValidate(data)
	utils.HasError(errValidate, "", 0)

	// 用户密码 再次加密
	data.Password = utils.GetSHA256HashCode([]byte(data.Password))

	sysUser := models.SysUserWithPsw{}
	sysUser.UserName = data.Username
	sysUser.Password = data.Password
	loginUser, getErr := sysUser.GetUser()
	utils.HasError(getErr, "用户名或密码不正确", 0)

	j := models.NewJWT()
	token, tokenErr := j.CreateToken(&models.UserClaims{
		ID:       loginUser.ID,
		UUID:     loginUser.UUID,
		Usercode: loginUser.UserCode,
		Username: loginUser.UserName,
		RoleID:   loginUser.RoleID,
		PostID:   loginUser.PostID,
		DeptID:   loginUser.DeptID,
	})
	utils.HasError(tokenErr, "创建token失败", 0)
	app.Custom(c, gin.H{
		"code": 1,
		"msg":  "登陆成功",
		"data": token,
	})
}

// GetCaptcha 获取验证码
// @Tags base
// @Summary 获取验证码
// @param type query string false "验证码类型"
// @Produce  application/json
// @Success 200 {string} string "{"code":200,"data":{"id": "captcha id", "captcha": "base64Captcha String"},"msg":"ok"}"
// @Router /base/captcha [get]
func GetCaptcha(c *gin.Context) {
	captchaBody := models.CaptchaBody{}
	captchaBody.CaptchaType = c.Request.FormValue("type")
	id, b64s, err := captchaBody.GenerateCaptcha()
	utils.HasError(err, "创建验证码失败", 0)
	var mp = make(map[string]interface{})
	mp["id"] = id
	mp["captcha"] = b64s
	app.Custom(c, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": mp,
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
	userClaims := utils.GetUserClaims(c)

	userView := models.SysUserInfo{}
	userView.UserCode = userClaims.Username
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
