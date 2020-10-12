package system

import (
	"errors"
	"go-core-frame/models"
	"go-core-frame/pkg/app"
	"go-core-frame/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetDictTypeDetail 获取 dictype 详情信息
// @Tags dic
// @Summary 获取 dicttype 详情信息
// @Produce  application/json
// @Param dictTypeId query int true "dictType编码"
// @Success 200 {object} app.Response "{"code": 1, "data": {...}, "msg": "ok"}"
// @Router /post/dict/{dictTypeId} [get]
// @Security Authorization
func GetDictTypeDetail(c *gin.Context) {
	dictTypeID, _ := utils.StringToInt(c.Param("dictTypeId"))
	sysDictType := models.SysDictType{}
	sysDictType.ID = dictTypeID
	data, err := sysDictType.GetDictType()
	utils.HasError(err, "抱歉未找到相关信息", 0)

	app.Custom(c, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": data,
	})
}

// GetDictTypeList DictType列表
// @Summary DictType列表
// @Tags dict
// @Param dictName query string false "字典名称""
// @Param dictType query string false "字典类型"
// @Param status query string false "状态""
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} app.Response "{"code": 1, "msg": "ok", "data": [...]}"
// @Router /dict/dicttypelist [get]
// @Security Authrization
func GetDictTypeList(c *gin.Context) {
	var data models.SysDictType
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

	data.DictName = c.Request.FormValue("dictName")
	data.DictType = c.Request.FormValue("dictType")
	data.Status = c.Request.FormValue("status")
	result, count, err := data.GetDictTypePage(pageSize, pageIndex)
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

// UpdateDictType 更新DictType
// @Summary 更新DictType
// @Tags dict
// @Param data body models.SysDictType true "body"
// @Success 200 {string} string	"{"code": 1, "msg": "修改成功"}"
// @Success 200 {string} string	"{"code": 0, "msg": "修改失败"}"
// @Router /dict/dicttypeupdate/{dictTypeId} [put]
func UpdateDictType(c *gin.Context) {
	var data models.SysDictType
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

	result, err := data.UpdateDictType()

	utils.HasError(err, "修改失败", 0)

	app.OK(c, "修改成功", result)
}

// InsertDictType 添加DictType
// @Summary 添加DictType
// @Tags dict
// @Param data body models.SysDictType true "dictType数据"
// @Success 200 {string} string	"{"code": 1, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 0, "message": "添加失败"}"
// @Router /api/v1/dict/dicttypeadd [post]
func InsertDictType(c *gin.Context) {
	var data models.SysDictType
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

	id, err := data.InsertDictType()
	utils.HasError(err, "添加失败", 0)
	app.OK(c, "添加成功", id)
}

// DeleteDictType 删除dicttype
// @Summary 删除dicttype
// @Tags dict
// @Param dictTypeId query string true "dictTypeid"
// @Success 200 {string} string	"{"code": 1, "msg": "删除成功"}"
// @Router /dict/dicttypedelete/{dictTypeId} [delete]
func DeleteDictType(c *gin.Context) {
	ID, _ := utils.StringToInt(c.Param("dictTypeId"))

	var data models.SysDictType
	data.ID = ID
	username, ok := c.Get("username")
	if !ok {
		username = "-"
	}
	data.UpdateBy = username.(string)

	err := data.DeleteDictType()
	utils.HasError(err, "删除失败", 0)
	app.OK(c, "删除成功", nil)
}

// GetDictMap dict字典数据
// @Summary dict字典数据
// @Tags dict
// @Success 200 {string} string	"{"code": 1, "msg": "ok", "data": {...}}"
// @Router /dict/dicAll [get]
func GetDictMap(c *gin.Context) {
	var dictType models.SysDictType
	dictTypeList, err := dictType.GetDictTypeAll()
	utils.HasError(err, "获取信息失败", 0)

	var mp = make(map[string]interface{})

	for _, dictType := range dictTypeList {

		var dictData = models.SysDictData{}
		dictData.DictTypeID = strconv.Itoa(dictType.ID)
		dictDataList, _ := dictData.GetDictDataAll()
		mp[dictType.DictType] = dictDataList
	}

	app.Custom(c, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": mp,
	})

}
