package system

import (
	"errors"
	"go-core-frame/models"
	"go-core-frame/pkg/app"
	"go-core-frame/utils"

	"github.com/gin-gonic/gin"
)

// GetDictDataDetail 获取 dictdata 详情信息
// @Tags dict
// @Summary 获取 dictdata 详情信息
// @Produce  application/json
// @Param dictDataId query int true "dictData编码"
// @Success 200 {object} app.Response "{"code": 1, "data": {...}, "msg": "ok"}"
// @Router /dict/{dictDataId} [get]
// @Security Authorization
func GetDictDataDetail(c *gin.Context) {
	dictDataID, _ := utils.StringToInt(c.Param("dictDataId"))
	sysDictData := models.SysDictData{}
	sysDictData.ID = dictDataID
	data, err := sysDictData.GetDictData()
	utils.HasError(err, "抱歉未找到相关信息", 0)

	app.Custom(c, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": data,
	})
}

// GetDictDataList DictData列表
// @Summary DictData列表
// @Tags dict
// @Param dictLabel query string false "字典lebel""
// @Param dictValue query string false "字典值"
// @Param dictTypeId query string false "字典typeId"
// @Param status query string false "状态""
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} app.Response "{"code": 1, "msg": "ok", "data": [...]}"
// @Router /dict/dictdatalist [get]
// @Security Authrization
func GetDictDataList(c *gin.Context) {
	var data models.SysDictData
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

	dictTypeID := c.Request.FormValue("dictTypeId")
	if dictTypeID == "" {
		err = errors.New("dictTypeId 不能为空")
		utils.HasError(err, "", 0)
	}
	data.DictTypeID = dictTypeID

	data.DictLabel = c.Request.FormValue("dictLabel")
	data.DictValue = c.Request.FormValue("dictValue")
	data.Status = c.Request.FormValue("status")
	result, count, err := data.GetDictDataPage(pageSize, pageIndex)
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

// UpdateDictData 更新DictData
// @Summary 更新DictData
// @Tags dict
// @Param data body models.SysDictData true "body"
// @Success 200 {string} string	"{"code": 1, "msg": "修改成功"}"
// @Success 200 {string} string	"{"code": 0, "msg": "修改失败"}"
// @Router /dict/dictdataupdate/{dictDataId} [put]
func UpdateDictData(c *gin.Context) {
	var data models.SysDictData
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

	result, err := data.UpdateDictData()

	utils.HasError(err, "修改失败", 0)

	app.OK(c, "修改成功", result)
}

// InsertDictData 添加DicData
// @Summary 添加DicData
// @Tags dict
// @Param data body models.SysDictData true "dictData数据"
// @Success 200 {string} string	"{"code": 1, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 0, "message": "添加失败"}"
// @Router /dict/dictdataadd [post]
func InsertDictData(c *gin.Context) {
	var data models.SysDictData
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

	id, err := data.InsertDictData()
	utils.HasError(err, "添加失败", 0)
	app.OK(c, "添加成功", id)
}

// DeleteDictData 删除dictData
// @Summary 删除dictData
// @Tags dict
// @Param dictDataId query string true "dictDataId"
// @Success 200 {string} string	"{"code": 1, "msg": "删除成功"}"
// @Router /dict/dictdatadelete/{dictDataId} [delete]
func DeleteDictData(c *gin.Context) {
	idsStr := c.Param("dictDataId")
	if idsStr == "" {
		err := errors.New("要删除的ID不能为空")
		utils.HasError(err, "", 0)
	}

	dictDataIds := utils.IdsStrToIdsIntGroup(idsStr)
	var data models.SysDictData
	username, ok := c.Get("username")
	if !ok {
		username = "-"
	}
	data.UpdateBy = username.(string)

	err := data.DeleteDictData(dictDataIds)
	utils.HasError(err, "删除失败", 0)
	app.OK(c, "删除成功", nil)
}
