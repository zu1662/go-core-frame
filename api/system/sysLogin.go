package system

import (
	"errors"
	"go-core-frame/models"
	"go-core-frame/pkg/app"
	"go-core-frame/utils"

	"github.com/gin-gonic/gin"
)

// GetLoginLogInfo 登录日志详情
// @Summary 登录日志详情
// @Tags log
// @Param id query string false "id"
// @Success 200 {object} app.Response "{"code": 200, "msg": "ok", "data": {...}}"
// @Router /log/operloginfo [get]
// @Security Bearer
func GetLoginLogInfo(c *gin.Context) {
	var data models.LoginLog
	var err error

	id := c.Request.FormValue("id")

	if id == "" {
		err = errors.New("id为空")
		utils.HasError(err, "", 0)
		return
	}
	data.ID, _ = utils.StringToInt(id)

	// result, err := data.GetDetail()
	err = data.Delete()
	utils.HasError(err, "", 0)

	app.Custom(c, gin.H{
		"code": 1,
		"msg":  "ok",
		// "data": result,
	})
}

// GetLoginLogList 登录日志列表
// @Summary 登录日志列表
// @Tags log
// @Param operName query string false "登录名称"
// @Param operIp query string false "操作人ip地址""
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} app.Response "{"code": 200, "msg": "ok", "data": [...]}"
// @Router /log/operloglist [get]
// @Security Bearer
func GetLoginLogList(c *gin.Context) {
	var data models.LoginLog
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
	data.IPAddress = c.Request.FormValue("loginIp")
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
