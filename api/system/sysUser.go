package system

import (
	"errors"
	"go-core-frame/models"
	"go-core-frame/pkg/app"
	"go-core-frame/utils"

	"github.com/gin-gonic/gin"
)

// GetUserDetail 获取用户详情信息
// @Tags user
// @Summary 获取用户详情信息
// @Produce  application/json
// @Param userId query int true "用户编码"
// @Success 200 {object} app.Response "{"code": 1, "data": [...], "msg": "ok"}"
// @Router /user/info/{userId} [get]
// @Security Authorization
func GetUserDetail(c *gin.Context) {
	userID, _ := utils.StringToInt(c.Param("userId"))
	sysUser := models.SysUser{}
	sysUser.ID = userID
	nowUser, err := sysUser.GetUser()
	utils.HasError(err, "抱歉未找到相关信息", 0)

	app.Custom(c, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": nowUser,
	})
}

// GetUserList 用户列表
// @Summary 用户列表
// @Tags user
// @Param userName query string false "用户名称"
// @Param mobile query string false "手机"
// @Param status query string false "状态""
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} app.Response "{"code": 1, "msg": "ok", "data": [...]}"
// @Router /user/list [get]
// @Security Authrization
func GetUserList(c *gin.Context) {
	var data models.SysUserView
	var err error
	var pageSize = 10
	var pageIndex = 1

	size := c.Request.FormValue("pageSize")
	if size != "" {
		pageSize, _ = utils.StringToInt(size)
	}

	index := c.Request.FormValue("pageIndex")
	if index != "" {
		pageIndex, _ = utils.StringToInt(index)
	}

	data.UserName = c.Request.FormValue("userName")
	data.Mobile = c.Request.FormValue("mobile")
	data.Status = c.Request.FormValue("status")
	data.DeptID, _ = utils.StringToInt(c.Request.FormValue("deptId"))
	result, count, err := data.GetPage(pageSize, pageIndex)
	utils.HasError(err, "", 0)

	var mp = make(map[string]interface{}, 3)
	mp["list"] = result
	mp["total"] = count
	mp["pageIndex"] = pageIndex
	mp["pageSize"] = pageSize

	app.Custom(c, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": mp,
	})
}

// UpdateUser 更新用户
// @Summary 更新用户
// @Tags user
// @Param data body models.SysUser true "body"
// @Success 200 {string} string	"{"code": 1, "msg": "修改成功"}"
// @Success 200 {string} string	"{"code": 0, "msg": "修改失败"}"
// @Router /user/update/{userId} [put]
// @Security Authrization
func UpdateUser(c *gin.Context) {
	var data models.SysUser
	err := c.Bind(&data)
	utils.HasError(err, "", 0)

	if data.ID <= 0 {
		err = errors.New("id 不能为空")
		utils.HasError(err, "", 0)
	}

	username, ok := c.Get("username")
	if !ok {
		username = "-"
	}
	data.UpdateBy = username.(string)
	data.UpdateTime = utils.GetCurrentTime()

	result, err := data.UpdateUser()

	utils.HasError(err, "修改失败", 0)

	app.OK(c, "修改成功", result)
}

// ResetUserPsw 更新用户密码
// @Summary 更新用户密码
// @Tags user
// @Param data body models.SysUser true "body"
// @Success 200 {string} string	"{"code": 1, "msg": "修改成功"}"
// @Success 200 {string} string	"{"code": 0, "msg": "修改失败"}"
// @Router /user/resetpsw [put]
// @Security Authrization
func ResetUserPsw(c *gin.Context) {
	var data models.SysUserWithPsw
	err := c.Bind(&data)
	utils.HasError(err, "", 0)

	if data.ID <= 0 {
		err = errors.New("id 不能为空")
		utils.HasError(err, "", 0)
	}

	// 用户密码 再次加密
	data.Password = utils.GetSHA256HashCode([]byte(data.Password))

	username, ok := c.Get("username")
	if !ok {
		username = "-"
	}
	data.UpdateBy = username.(string)
	data.UpdateTime = utils.GetCurrentTime()

	result, err := data.UpdateUser()

	utils.HasError(err, "修改失败", 0)

	app.OK(c, "修改成功", result)
}

// InsertUser 添加用户
// @Summary 添加用户
// @Tags user
// @Param data body models.SysUser true "用户数据"
// @Success 200 {string} string	"{"code": 1, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 0, "message": "添加失败"}"
// @Router /api/v1/user/add [post]
// @Security Authrization
func InsertUser(c *gin.Context) {
	var sysuser models.SysUserWithPsw
	err := c.Bind(&sysuser)
	utils.HasError(err, "非法数据格式", 0)

	errValidate := utils.StructValidate(sysuser)
	utils.HasError(errValidate, "", 0)

	username, ok := c.Get("username")
	if !ok {
		username = "-"
	}
	sysuser.CreateBy = username.(string)
	sysuser.CreateTime = utils.GetCurrentTime()
	sysuser.UpdateBy = username.(string)
	sysuser.UpdateTime = utils.GetCurrentTime()

	// 生成用户 uuid
	uuid := utils.GetUUID()
	sysuser.UUID = uuid

	// 新建账号默认密码为123456
	md5Str := utils.GetMD5HashCode([]byte("123456"))
	// 密码 SHA256 加密
	sysuser.Password = utils.GetSHA256HashCode([]byte(md5Str))

	id, err := sysuser.InsertUser()
	utils.HasError(err, "添加失败", 0)
	app.OK(c, "添加成功", id)
}

// DeleteUser 删除用户
// @Summary 删除用户
// @Tags user
// @Param userId query string true "用户id"
// @Success 200 {string} string	"{"code": 1, "msg": "删除成功"}"
// @Router /user/delete/{userId} [delete]
// @Security Authrization
func DeleteUser(c *gin.Context) {
	idsStr := c.Param("userId")

	var data models.SysUser
	userIds := utils.IdsStrToIdsIntGroup(idsStr)

	username, ok := c.Get("username")
	if !ok {
		username = "-"
	}
	data.UpdateBy = username.(string)

	err := data.DeleteUser(userIds)
	utils.HasError(err, "删除失败", 0)
	app.OK(c, "删除成功", nil)
}
