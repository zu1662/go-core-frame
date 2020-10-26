package utils

import (
	"go-core-frame/models"

	"github.com/gin-gonic/gin"
)

// GetUserClaims 获取token内的用户信息
func GetUserClaims(c *gin.Context) models.UserClaims {
	// 获取当前用户信息
	nowClaims, ok := c.Get("claims")
	userClaims := models.UserClaims{}
	userClaims.Username = "_"
	if ok {
		userClaims = nowClaims.(models.UserClaims)
	}
	return userClaims
}
