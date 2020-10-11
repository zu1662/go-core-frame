package system

import (
	"errors"
	"go-core-frame/models"
	"go-core-frame/pkg/app"
	"go-core-frame/utils"

	"github.com/gin-gonic/gin"
)

// GetPostDetail 获取岗位详情信息
// @Tags post
// @Summary 获取岗位详情信息
// @Produce  application/json
// @Param postId int true "部门编码"
// @Success 200 {object} app.Response "{"code": 1, "data": {...}, "msg": "ok"}"
// @Router /post/info/{postId} [get]
// @Security Authorization
func GetPostDetail(c *gin.Context) {
	postID, _ := utils.StringToInt(c.Param("postId"))
	sysPost := models.SysPost{}
	sysPost.ID = postID
	nowDept, err := sysPost.GetPost()
	utils.HasError(err, "抱歉未找到相关信息", 0)

	app.Custom(c, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": nowDept,
	})
}

// GetPostList 岗位列表
// @Summary 岗位列表
// @Tags post
// @Param postCode query string false "岗位编码"
// @Param postName query string false "岗位名称"
// @Param status query string false "状态""
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} app.Response "{"code": 1, "msg": "ok", "data": [...]}"
// @Router /post/list [get]
// @Security Authrization
func GetPostList(c *gin.Context) {
	var data models.SysPost
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

	data.PostCode = c.Request.FormValue("postCode")
	data.PostName = c.Request.FormValue("postName")
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

// UpdatePost 更新岗位
// @Summary 更新岗位
// @Tags post
// @Param data body models.SysPost true "body"
// @Success 200 {string} string	"{"code": 1, "msg": "修改成功"}"
// @Success 200 {string} string	"{"code": 0, "msg": "修改失败"}"
// @Router /post/update/{postId} [put]
func UpdatePost(c *gin.Context) {
	var data models.SysPost
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

	result, err := data.UpdatePost()

	utils.HasError(err, "修改失败", 0)

	app.OK(c, "修改成功", result)
}

// InsertPost 添加岗位
// @Summary 添加岗位
// @Tags post
// @Param data body models.SysPost true "岗位数据"
// @Success 200 {string} string	"{"code": 1, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 0, "message": "添加失败"}"
// @Router /api/v1/post/add [post]
func InsertPost(c *gin.Context) {
	var data models.SysPost
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

	id, err := data.InsertPost()
	utils.HasError(err, "添加失败", 0)
	app.OK(c, "添加成功", id)
}

// DeletePost 删除岗位
// @Summary 删除岗位
// @Tags dept
// @Param postId query string true "岗位id"
// @Success 200 {string} string	"{"code": 1, "msg": "删除成功"}"
// @Router /post/delete/{postId} [delete]
func DeletePost(c *gin.Context) {
	ID, _ := utils.StringToInt(c.Param("postId"))

	var data models.SysPost
	data.ID = ID
	username, ok := c.Get("username")
	if !ok {
		username = "-"
	}
	data.UpdateBy = username.(string)

	err := data.DeletePost()
	utils.HasError(err, "删除失败", 0)
	app.OK(c, "删除成功", nil)
}
