package system

import (
	"errors"
	"go-core-frame/models"
	"go-core-frame/pkg/app"
	"go-core-frame/utils"

	"github.com/gin-gonic/gin"
)

// GetMenuDetail 获取菜单详情信息
// @Tags menu
// @Summary 获取菜单详情信息
// @Produce  application/json
// @Param menuId query int true "菜单编码"
// @Success 200 {object} app.Response "{"code": 1, "data": {...}, "msg": "ok"}"
// @Router /menu/info/{menuId} [get]
// @Security Authorization
func GetMenuDetail(c *gin.Context) {
	menuID, _ := utils.StringToInt(c.Param("menuId"))
	sysMenu := models.SysMenu{}
	sysMenu.ID = menuID
	nowDept, err := sysMenu.GetMenu()
	utils.HasError(err, "抱歉未找到相关信息", 0)

	app.Custom(c, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": nowDept,
	})
}

// UpdateMenu 更新菜单
// @Summary 更新菜单
// @Tags menu
// @Param data body models.SysMenu true "body"
// @Success 200 {string} string	"{"code": 1, "msg": "修改成功"}"
// @Success 200 {string} string	"{"code": 0, "msg": "修改失败"}"
// @Router /menu/update/{menuId} [put]
func UpdateMenu(c *gin.Context) {
	var data models.SysMenu
	err := c.Bind(&data)
	utils.HasError(err, "", 0)

	if data.ID <= 0 {
		err = errors.New("id 不能为空")
		utils.HasError(err, "", 0)
	}

	errValidate := utils.StructValidate(data)
	utils.HasError(errValidate, "", 0)

	username, ok := c.Get("username")
	if !ok {
		username = "-"
	}
	data.UpdateBy = username.(string)
	data.UpdateTime = utils.GetCurrentTime()

	result, err := data.UpdateMenu()

	utils.HasError(err, "修改失败", 0)

	app.OK(c, "修改成功", result)
}

// InsertMenu 添加菜单
// @Summary 添加菜单
// @Tags menu
// @Param data body models.SysMenu true "部门数据"
// @Success 200 {string} string	"{"code": 1, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 0, "message": "添加失败"}"
// @Router /api/v1/menu/add [post]
func InsertMenu(c *gin.Context) {
	var data models.SysMenu
	err := c.Bind(&data)
	utils.HasError(err, "非法数据格式", 0)

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

	id, err := data.InsertMenu()
	utils.HasError(err, "添加失败", 0)
	app.OK(c, "添加成功", id)
}

// DeleteMenu 删除菜单
// @Summary 删除菜单
// @Tags menu
// @Param menuId query string true "菜单id"
// @Success 200 {string} string	"{"code": 1, "msg": "删除成功"}"
// @Router /menu/delete/{menuId} [delete]
func DeleteMenu(c *gin.Context) {
	ID, _ := utils.StringToInt(c.Param("menuId"))

	var data models.SysMenu
	data.ID = ID
	username, ok := c.Get("username")
	if !ok {
		username = "-"
	}
	data.UpdateBy = username.(string)

	err := data.DeleteMenu()
	utils.HasError(err, "删除失败", 0)
	app.OK(c, "删除成功", "")
}

// GetMenuTree 菜单tree
// @Summary 菜单tree
// @Tags menu
// @Param menuId query string false "menuID"
// @Success 200 {object} app.Response "{"code": 1, "msg": "ok", "data": [...]}"
// @Router /menu/tree [get]
// @Security Authrization
func GetMenuTree(c *gin.Context) {
	// 获取用户信息
	userClaims := utils.GetUserClaims(c)
	var data models.SysMenuView
	var err error

	ID, _ := utils.StringToInt(c.Request.FormValue("menuId"))
	data.ID = ID

	result, err := data.GetMenuTree(userClaims.RoleID)
	utils.HasError(err, "", 0)

	app.Custom(c, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": result,
	})
}
