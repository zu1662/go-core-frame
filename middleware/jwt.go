package middleware

import (
	"go-core-frame/models"
	"go-core-frame/pkg/app"
	"go-core-frame/pkg/config"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// JWTAuth JWT登录校验
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(config.JWTConfig.HeaderName)
		if token == "" {
			app.Error(c, 401, models.ErrTokenEmpty, "")
			c.Abort()
			return
		}
		// if service.IsBlacklist(token) {
		// 	response.Result(response.ERROR, gin.H{
		// 		"reload": true,
		// 	}, "您的帐户异地登陆或令牌失效", c)
		// 	c.Abort()
		// 	return
		// }
		j := models.NewJWT()
		claims, err := j.ParseToken(token)

		// 设置 username 便于 logger 使用
		c.Set("username", claims.UserClaims.Username)

		if err != nil {
			if err == models.ErrTokenExpired {
				// 当Token已过期，但是 符合 BufferTime 的时间内，生成新的 Token
				if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
					claims.StandardClaims.ExpiresAt = time.Now().Unix() + j.Timeout
					newToken, _ := j.CreateToken(&claims.UserClaims)
					c.Header("New-Token", newToken.Token)
					c.Header("New-ExpiresAt", strconv.FormatInt(newToken.Expire, 10))
					c.Next()
					return
				}
				app.Error(c, 401, models.ErrTokenExpired, "Token已失效，请重新登录")
				c.Abort()
				return
			}
			app.Error(c, 401, err, "")
			c.Abort()
			return
		}
		c.Next()
	}
}
