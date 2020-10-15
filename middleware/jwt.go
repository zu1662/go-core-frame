package middleware

import (
	"go-core-frame/global"
	"go-core-frame/models"
	"go-core-frame/pkg/app"
	"go-core-frame/pkg/config"
	"math"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// JWTAuth JWT登录校验
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(config.JWTConfig.HeaderName)
		if token == "" {
			app.WithCode(c, 401, models.ErrTokenEmpty, "Token为空，请重新登录")
			c.Abort()
			return
		}

		// 获取 redis 内的 token
		IsBlacklist, _ := global.Redis.SIsMember("tokenBlock", token).Result()
		if IsBlacklist {
			app.WithCode(c, 401, models.ErrTokenEmpty, "您的帐户异地登陆或令牌失效")
			c.Abort()
			return
		}

		// 从redis获取用户信息
		userJSON, _ := global.Redis.Get(token).Result()
		var userClaims models.UserClaims
		userClaims.UnmarshalBinary([]byte(userJSON))

		j := models.NewJWT()
		claims, err := j.ParseToken(token)

		if err != nil {
			if err == models.ErrTokenExpired {
				// 当Token已过期，但是 符合 BufferTime 的时间内，生成新的 Token
				if math.Abs(float64(claims.ExpiresAt-time.Now().Unix())) < float64(claims.BufferTime) {
					claims.StandardClaims.ExpiresAt = time.Now().Unix() + j.Timeout
					newToken, _ := j.CreateToken(&claims.UserClaims)
					c.Header("New-Token", newToken.Token)
					c.Header("New-Expires", strconv.FormatInt(newToken.Expire, 10))
					c.Next()
					return
				}
				app.WithCode(c, 401, models.ErrTokenExpired, "Token已失效，请重新登录")
				c.Abort()
				return
			}
			app.WithCode(c, 401, err, "无法解析Token，请重新登录")
			c.Abort()
			return
		}

		// 设置 username 便于 logger 使用
		c.Set("username", claims.UserClaims.Usercode)

		c.Next()
	}
}
