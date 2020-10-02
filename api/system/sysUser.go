package system

import (
	"go-core-frame/models"
	"go-core-frame/utils"

	"github.com/gin-gonic/gin"
)

// GetSysUser 获取用户详情信息
func GetSysUser(c *gin.Context) {
	userID, _ := utils.StringToInt(c.Param("userId"))
	result, _ := models.GetUserInfo(userID)
	c.JSON(200, gin.H{
		"message": "ok",
		"result":  result,
	})
}
