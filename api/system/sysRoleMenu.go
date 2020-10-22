package system

import (
	"errors"
	"go-core-frame/models"
	"go-core-frame/pkg/app"
	"go-core-frame/utils"

	"github.com/gin-gonic/gin"
)

// GetRoleMenu 角色列表
// @Summary 角色列表
// @Tags rolemenu
// @Param roleId query string false "角色ID"
// @Success 200 {object} app.Response "{"code": 1, "msg": "ok", "data": [...]}"
// @Router /rolemenu/list [get]
// @Security Authrization
func GetRoleMenu(c *gin.Context) {
	var data models.SysRoleMenu

	roleID, _ := utils.StringToInt(c.Request.FormValue("roleId"))
	if roleID <= 0 {
		err := errors.New("roleId 不能为空")
		utils.HasError(err, "", 0)
	}
	data.RoleID = roleID

	arr, err := data.GetRoleMenu()
	utils.HasError(err, "", 0)

	var mp = make(map[string]interface{}, 3)
	mp["roleId"] = data.RoleID
	mp["menuList"] = arr

	app.Custom(c, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": mp,
	})
}

// UpdateRoleMenu 更新角色菜单
// @Summary 更新角色菜单
// @Tags rolemenu
// @Param data body models.SysRoleMenuView true "body"
// @Success 200 {string} string	"{"code": 1, "msg": "修改成功"}"
// @Success 200 {string} string	"{"code": 0, "msg": "修改失败"}"
// @Router /rolemenu/update/{roleId} [put]
func UpdateRoleMenu(c *gin.Context) {
	var data models.SysRoleMenuView
	err := c.Bind(&data)
	utils.HasError(err, "", 0)

	if data.RoleID <= 0 {
		err = errors.New("roleId 不能为空")
		utils.HasError(err, "", 0)
	}

	errValidate := utils.StructValidate(data)
	utils.HasError(errValidate, "", 0)

	username, ok := c.Get("username")
	if !ok {
		username = "-"
	}
	data.CreateBy = username.(string)
	data.CreateTime = utils.GetCurrentTime()
	data.UpdateBy = username.(string)
	data.UpdateTime = utils.GetCurrentTime()

	err = data.UpdateRoleMenu()

	utils.HasError(err, "修改失败", 0)

	app.OK(c, "修改成功", nil)
}
