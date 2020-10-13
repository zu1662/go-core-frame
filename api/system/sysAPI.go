package system

import (
	"errors"
	"go-core-frame/models"
	"go-core-frame/pkg/app"
	"go-core-frame/utils"

	"github.com/gin-gonic/gin"
)

// GetAPIDetail 获取API详情信息
// @Tags api
// @Summary 获取API详情信息
// @Produce  application/json
// @Param apiId query int true "api编码"
// @Success 200 {object} app.Response "{"code": 1, "data": {...}, "msg": "ok"}"
// @Router /api/info/{apiId} [get]
// @Security Authorization
func GetAPIDetail(c *gin.Context) {
	apiID, _ := utils.StringToInt(c.Param("apiId"))
	sysAPI := models.SysAPI{}
	sysAPI.ID = apiID
	nowAPI, err := sysAPI.GetAPI()
	utils.HasError(err, "抱歉未找到相关信息", 0)

	app.Custom(c, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": nowAPI,
	})
}

// UpdateAPI 更新api
// @Summary 更新api
// @Tags api
// @Param data body models.SysAPI true "body"
// @Success 200 {string} string	"{"code": 1, "msg": "修改成功"}"
// @Success 200 {string} string	"{"code": 0, "msg": "修改失败"}"
// @Router /api/update/{apiId} [put]
func UpdateAPI(c *gin.Context) {
	var data models.SysAPI
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

	result, err := data.UpdateAPI()

	utils.HasError(err, "修改失败", 0)

	app.OK(c, "修改成功", result)
}

// InsertAPI 添加api
// @Summary 添加api
// @Tags menu
// @Param data body models.SysAPI true "部门数据"
// @Success 200 {string} string	"{"code": 1, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 0, "message": "添加失败"}"
// @Router /api/add [post]
func InsertAPI(c *gin.Context) {
	var data models.SysAPI
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

	id, err := data.InsertAPI()
	utils.HasError(err, "添加失败", 0)
	app.OK(c, "添加成功", id)
}

// DeleteAPI 删除api
// @Summary 删除api
// @Tags api
// @Param apiId query string true "api id"
// @Success 200 {string} string	"{"code": 1, "msg": "删除成功"}"
// @Router /api/delete/{apiId} [delete]
func DeleteAPI(c *gin.Context) {
	ID, _ := utils.StringToInt(c.Param("apiId"))

	var data models.SysAPI
	data.ID = ID
	username, ok := c.Get("username")
	if !ok {
		username = "-"
	}
	data.UpdateBy = username.(string)

	err := data.DeleteAPI()
	utils.HasError(err, "删除失败", 0)
	app.OK(c, "删除成功", "")
}

// GetAPITree api tree
// @Summary api tree
// @Tags api
// @Param apiId query string false "apiId"
// @Success 200 {object} app.Response "{"code": 1, "msg": "ok", "data": [...]}"
// @Router /api/tree [get]
// @Security Authrization
func GetAPITree(c *gin.Context) {
	var data models.SysAPIView
	var err error

	ID, _ := utils.StringToInt(c.Request.FormValue("apiId"))
	data.ID = ID
	result, err := data.GetAPITree()
	utils.HasError(err, "", 0)

	app.Custom(c, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": result,
	})
}
