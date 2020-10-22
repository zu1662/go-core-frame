package system

import (
	"errors"
	"go-core-frame/models"
	"go-core-frame/pkg/app"
	"go-core-frame/utils"

	"github.com/gin-gonic/gin"
)

// GetRoleDetail 获取角色详情信息
// @Tags role
// @Summary 获取角色详情信息
// @Produce  application/json
// @Param roleId query int true "角色编码"
// @Success 200 {object} app.Response "{"code": 1, "data": {...}, "msg": "ok"}"
// @Router /role/info/{roleId} [get]
// @Security Authorization
func GetRoleDetail(c *gin.Context) {
	roleID, _ := utils.StringToInt(c.Param("roleId"))
	if roleID <= 0 {
		err := errors.New("roleId 不能为空")
		utils.HasError(err, "", 0)
	}

	sysRoleView := models.SysRoleView{}
	sysRoleView.ID = roleID
	nowRoleView, err := sysRoleView.GetRole()
	utils.HasError(err, "抱歉未找到相关信息", 0)

	// 获取 role menu 关联信息
	var menuData models.SysRoleMenu

	menuData.RoleID = roleID

	arr, err := menuData.GetRoleMenu()
	utils.HasError(err, "", 0)

	nowRoleView.MenuList = arr

	app.Custom(c, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": nowRoleView,
	})
}

// GetRoleList 角色列表
// @Summary 角色列表
// @Tags role
// @Param roleCode query string false "角色编码"
// @Param roleName query string false "角色名称"
// @Param status query string false "状态""
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} app.Response "{"code": 1, "msg": "ok", "data": [...]}"
// @Router /role/list [get]
// @Security Authrization
func GetRoleList(c *gin.Context) {
	var data models.SysRole
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

	data.RoleCode = c.Request.FormValue("roleCode")
	data.RoleName = c.Request.FormValue("roleName")
	data.Status = c.Request.FormValue("status")
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

// GetRoleAll 角色列表
// @Summary 角色列表
// @Tags role
// @Param roleCode query string false "角色编码"
// @Param roleName query string false "角色名称"
// @Param status query string false "状态""
// @Success 200 {object} app.Response "{"code": 1, "msg": "ok", "data": [...]}"
// @Router /role/listall [get]
// @Security Authrization
func GetRoleAll(c *gin.Context) {
	var data models.SysRole
	var err error

	data.RoleCode = c.Request.FormValue("roleCode")
	data.RoleName = c.Request.FormValue("roleName")
	data.Status = c.Request.FormValue("status")
	result, err := data.GetList()
	utils.HasError(err, "", 0)

	app.Custom(c, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": result,
	})
}

// UpdateRole 更新角色
// @Summary 更新角色
// @Tags role
// @Param data body models.SysRole true "body"
// @Success 200 {string} string	"{"code": 1, "msg": "修改成功"}"
// @Success 200 {string} string	"{"code": 0, "msg": "修改失败"}"
// @Router /role/update/{roleId} [put]
func UpdateRole(c *gin.Context) {
	var data models.SysRoleView
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

	result, err := data.UpdateRole()

	utils.HasError(err, "修改失败", 0)

	// 修改 Role Menu 关联
	var roleMenuData models.SysRoleMenuView
	roleMenuData.RoleID = data.ID
	roleMenuData.MenuList = data.MenuList
	roleMenuData.BaseModel = data.BaseModel
	err = roleMenuData.UpdateRoleMenu()
	utils.HasError(err, "修改失败", 0)

	app.OK(c, "修改成功", result)
}

// InsertRole 添加角色
// @Summary 添加角色
// @Tags role
// @Param data body models.SysRole true "角色数据"
// @Success 200 {string} string	"{"code": 1, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 0, "message": "添加失败"}"
// @Router /api/v1/role/add [post]
func InsertRole(c *gin.Context) {
	var data models.SysRoleView
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

	id, err := data.InsertRole()
	utils.HasError(err, "添加失败", 0)

	// 修改 Role Menu 关联
	var roleMenuData models.SysRoleMenuView
	roleMenuData.RoleID = id
	roleMenuData.MenuList = data.MenuList
	roleMenuData.BaseModel = data.BaseModel
	err = roleMenuData.UpdateRoleMenu()
	utils.HasError(err, "修改失败", 0)

	app.OK(c, "添加成功", id)
}

// DeleteRole 删除角色
// @Summary 删除角色
// @Tags role
// @Param roleId query string true "角色id"
// @Success 200 {string} string	"{"code": 1, "msg": "删除成功"}"
// @Router /role/delete/{roleId} [delete]
func DeleteRole(c *gin.Context) {
	idsStr := c.Param("roleId")
	if idsStr == "" {
		err := errors.New("要删除的ID不能为空")
		utils.HasError(err, "", 0)
	}
	roleIds := utils.IdsStrToIdsIntGroup(idsStr)

	var data models.SysRole
	username, ok := c.Get("username")
	if !ok {
		username = "-"
	}
	data.UpdateBy = username.(string)

	err := data.DeleteRole(roleIds)
	utils.HasError(err, "删除失败", 0)
	app.OK(c, "删除成功", nil)
}
