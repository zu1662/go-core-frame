package system

import (
	"errors"
	"go-core-frame/models"
	"go-core-frame/pkg/app"
	"go-core-frame/utils"

	"github.com/gin-gonic/gin"
)

// GetOperLogInfo 操作日志详情
// @Summary 操作日志详情
// @Tags log
// @Param id query string false "id"
// @Success 200 {object} app.Response "{"code": 200, "msg": "ok", "data": {...}}"
// @Router /log/operloginfo [get]
// @Security Bearer
func GetOperLogInfo(c *gin.Context) {
	var data models.OperLog
	var err error

	id := c.Request.FormValue("id")

	if id == "" {
		err = errors.New("id为空")
		utils.HasError(err, "", 0)
		return
	}
	data.ID, _ = utils.StringToInt(id)

	result, err := data.GetDetail()
	utils.HasError(err, "", 0)

	app.Custom(c, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": result,
	})
}

// GetOperLogList 操作日志列表
// @Summary 操作日志列表
// @Tags log
// @Param operName query string false "操作人"
// @Param method query string false "请求方式"
// @Param operIp query string false "操作人ip地址""
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} app.Response "{"code": 200, "msg": "ok", "data": [...]}"
// @Router /log/operloglist [get]
// @Security Bearer
func GetOperLogList(c *gin.Context) {
	var data models.OperLog
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

	data.OperName = c.Request.FormValue("operName")
	data.Method = c.Request.FormValue("method")
	data.IPAddress = c.Request.FormValue("operIp")
	result, count, err := data.GetPage(pageSize, pageIndex)
	utils.HasError(err, "", 0)

	var mp = make(map[string]interface{})
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

// DeleteOperlog 删除操作日志
// @Summary 删除操作日志
// @Tags log
// @Param logIds query string true "岗位ids"
// @Success 200 {string} string	"{"code": 1, "msg": "删除成功"}"
// @Router /log/deleteoperlog/{logIds} [delete]
func DeleteOperlog(c *gin.Context) {
	idsStr := c.Param("logIds")
	if idsStr == "" {
		err := errors.New("要删除的ID不能为空")
		utils.HasError(err, "", 0)
	}

	logIds := utils.IdsStrToIdsIntGroup(idsStr)

	var data models.OperLog

	err := data.Delete(logIds)
	utils.HasError(err, "删除失败", 0)
	app.OK(c, "删除成功", nil)
}

// CleanOperlog 清空操作日志
// @Summary 清空操作日志
// @Tags log
// @Success 200 {string} string	"{"code": 1, "msg": "删除成功"}"
// @Router /log/cleanoperlog/ [delete]
func CleanOperlog(c *gin.Context) {
	var data models.OperLog

	err := data.Clean()
	utils.HasError(err, "", 0)
	app.OK(c, "清空成功", nil)
}
