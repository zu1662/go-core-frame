package system

import (
	"errors"
	"go-core-frame/models"
	"go-core-frame/pkg/app"
	"go-core-frame/utils"

	"github.com/gin-gonic/gin"
)

// GetDeptDetail 获取部门详情信息
// @Tags dept
// @Summary 获取部门详情信息
// @Produce  application/json
// @Param deptId query int true "部门编码"
// @Success 200 {object} app.Response "{"code": 1, "data": {...}, "msg": "ok"}"
// @Router /dept/info/{deptId} [get]
// @Security Authorization
func GetDeptDetail(c *gin.Context) {
	deptID, _ := utils.StringToInt(c.Param("deptId"))
	sysDept := models.SysDept{}
	sysDept.ID = deptID
	nowDept, err := sysDept.GetDept()
	utils.HasError(err, "抱歉未找到相关信息", 0)

	sysUser := models.SysUser{}
	sysUser.ID = nowDept.LeaderID
	nowUser, _ := sysUser.GetUser()

	var mp = make(map[string]interface{})
	mp["id"] = nowDept.ID
	mp["pid"] = nowDept.Pid
	mp["deptName"] = nowDept.DeptName
	mp["sort"] = nowDept.Sort
	mp["status"] = nowDept.Status
	mp["leaderId"] = nowDept.LeaderID
	mp["leaderName"] = nowUser.UserName
	mp["leaderEmail"] = nowUser.Email
	mp["leaderMobile"] = nowUser.Mobile
	app.Custom(c, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": mp,
	})
}

// UpdateDept 更新部门
// @Summary 更新部门
// @Tags dept
// @Param data body models.SysDept true "body"
// @Success 200 {string} string	"{"code": 1, "msg": "修改成功"}"
// @Success 200 {string} string	"{"code": 0, "msg": "修改失败"}"
// @Router /dept/update/{deptId} [put]
func UpdateDept(c *gin.Context) {
	var data models.SysDept
	err := c.Bind(&data)
	utils.HasError(err, "", 0)

	if data.ID <= 0 {
		err = errors.New("id 不能为空")
		utils.HasError(err, "", 0)
	}

	errValidate := utils.StructValidate(data)
	utils.HasError(errValidate, "", 0)

	// 获取用户信息
	userClaims := utils.GetUserClaims(c)
	data.UpdateBy = userClaims.Username
	data.UpdateTime = utils.GetCurrentTime()

	result, err := data.UpdateDept()

	utils.HasError(err, "修改失败", 0)

	app.OK(c, "修改成功", result)
}

// InsertDept 添加部门
// @Summary 添加部门
// @Tags dept
// @Param data body models.SysDept true "部门数据"
// @Success 200 {string} string	"{"code": 1, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 0, "message": "添加失败"}"
// @Router /api/v1/dept/add [post]
func InsertDept(c *gin.Context) {
	var data models.SysDept
	err := c.Bind(&data)
	utils.HasError(err, "非法数据格式", 0)

	errValidate := utils.StructValidate(data)
	utils.HasError(errValidate, "", 0)

	// 获取用户信息
	userClaims := utils.GetUserClaims(c)
	data.CreateBy = userClaims.Username
	data.CreateTime = utils.GetCurrentTime()
	data.UpdateBy = userClaims.Username
	data.UpdateTime = utils.GetCurrentTime()

	id, err := data.InsertDept()
	utils.HasError(err, "添加失败", 0)
	app.OK(c, "添加成功", id)
}

// DeleteDept 删除部门
// @Summary 删除部门
// @Tags dept
// @Param deptId query string true "部门id"
// @Success 200 {string} string	"{"code": 1, "msg": "删除成功"}"
// @Router /dept/delete/{deptId} [delete]
func DeleteDept(c *gin.Context) {
	ID, _ := utils.StringToInt(c.Param("deptId"))

	var data models.SysDept
	data.ID = ID
	// 获取用户信息
	userClaims := utils.GetUserClaims(c)
	data.UpdateBy = userClaims.Username

	err := data.DeleteDept()
	utils.HasError(err, "删除失败", 0)
	app.OK(c, "删除成功", "")
}

// GetDeptTree 部门tree
// @Summary 部门tree
// @Tags dept
// @Param deptId query string false "deptId"
// @Success 200 {object} app.Response "{"code": 1, "msg": "ok", "data": [...]}"
// @Router /dept/tree [get]
// @Security Authrization
func GetDeptTree(c *gin.Context) {
	var data models.SysDeptView
	var err error

	ID, _ := utils.StringToInt(c.Request.FormValue("deptId"))
	data.ID = ID
	result, err := data.GetDeptTree()
	utils.HasError(err, "", 0)

	app.Custom(c, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": result,
	})
}
