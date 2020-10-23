package middleware

import (
	"errors"
	"go-core-frame/global"
	"go-core-frame/models"
	"go-core-frame/pkg/app"
	"go-core-frame/pkg/config"
	"go-core-frame/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// APIAuth api权限校验
func APIAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(config.JWTConfig.HeaderName)

		// 从redis获取用户信息
		userJSON, _ := global.Redis.Get(token).Result()
		var userClaims models.UserClaims
		userClaims.UnmarshalBinary([]byte(userJSON))

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqURI := utils.GetAPIPath(c.FullPath(), "")

		api := models.SysAPI{}
		api.Path = strings.TrimLeft(reqURI, "/v1")
		api.Method = reqMethod
		nowAPI, apiErr := api.GetAPI()

		if apiErr != nil {
			app.WithCode(c, 403, errors.New("接口地址不正确"), "")
			c.Abort()
			return
		}

		sysRoleAPI := models.SysRoleAPI{}
		sysRoleAPI.APIID = nowAPI.ID
		sysRoleAPI.RoleID = userClaims.RoleID
		nowRoleAPI, err := sysRoleAPI.GetRoleAPI()
		if err != nil {
			app.WithCode(c, 403, errors.New("接口不允许访问, 清联系管理员"), "")
			c.Abort()
			return
		}
		if len(nowRoleAPI) == 0 {
			app.WithCode(c, 403, errors.New("接口不允许访问, 请联系管理员"), "")
			c.Abort()
			return
		}

		c.Next()
	}
}
