package utils

import (
	"go-core-frame/global"
	"go-core-frame/models"
	"go-core-frame/pkg/config"

	"github.com/gin-gonic/gin"
)

// GetUserClaims 获取token内的用户信息
func GetUserClaims(c *gin.Context) models.UserClaims {
	// 获取Token
	token := c.Request.Header.Get(config.JWTConfig.HeaderName)
	// 从redis获取用户信息
	userJSON, _ := global.Redis.Get(token).Result()
	var userClaims models.UserClaims
	userClaims.UnmarshalBinary([]byte(userJSON))
	return userClaims
}
