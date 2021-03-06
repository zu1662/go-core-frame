package system

import (
	"errors"
	"go-core-frame/models"
	"go-core-frame/pkg/app"
	"go-core-frame/utils"

	"github.com/gin-gonic/gin"
)

// GetRoleAPI 角色接口列表
// @Summary 角色接口列表
// @Tags roleapi
// @Param roleId query string false "角色ID"
// @Success 200 {object} app.Response "{"code": 1, "msg": "ok", "data": [...]}"
// @Router /roleapi/list [get]
// @Security Authrization
func GetRoleAPI(c *gin.Context) {
	var data models.SysRoleAPI

	roleID := c.Request.FormValue("roleId")
	if roleID == "" {
		err := errors.New("roleId 不能为空")
		utils.HasError(err, "", 0)
	}
	data.RoleID, _ = utils.StringToInt(roleID)

	result, err := data.GetRoleAPI()
	utils.HasError(err, "", 0)

	var arr []int
	for _, sysRoleAPI := range result {
		arr = append(arr, sysRoleAPI.APIID)
	}
	var mp = make(map[string]interface{}, 3)
	mp["roleId"] = data.RoleID
	mp["apiList"] = arr

	app.Custom(c, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": mp,
	})
}

// UpdateRoleAPI 更新角色接口菜单
// @Summary 更新角色接口菜单
// @Tags roleapi
// @Param data body models.SysRoleAPIView true "body"
// @Success 200 {string} string	"{"code": 1, "msg": "修改成功"}"
// @Success 200 {string} string	"{"code": 0, "msg": "修改失败"}"
// @Router /roleapi/update/{roleId} [put]
func UpdateRoleAPI(c *gin.Context) {
	var data models.SysRoleAPIView
	err := c.Bind(&data)
	utils.HasError(err, "", 0)

	if data.RoleID == 0 {
		err = errors.New("roleId 不能为空")
		utils.HasError(err, "", 0)
	}

	errValidate := utils.StructValidate(data)
	utils.HasError(errValidate, "", 0)

// 获取用户信息
	userClaims := utils.GetUserClaims(c)
	data.CreateBy = userClaims.Username
	data.CreateTime = utils.GetCurrentTime()
	data.UpdateBy = userClaims.Username
	data.UpdateTime = utils.GetCurrentTime()

	err = data.UpdateRoleAPI()

	utils.HasError(err, "修改失败", 0)

	app.OK(c, "修改成功", nil)
}
