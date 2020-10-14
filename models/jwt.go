package models

import (
	"encoding/json"
	"errors"
	"go-core-frame/global"
	"go-core-frame/pkg/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	// ErrTokenExpired Token 过期
	ErrTokenExpired = errors.New("Token is expired")
	// ErrTokenMalformed 令牌格式不正确
	ErrTokenMalformed = errors.New("That's not even a token")
	// ErrTokenInvalid Token 校验失败
	ErrTokenInvalid = errors.New("Couldn't handle this token")
	// ErrTokenEmpty Token 为空
	ErrTokenEmpty = errors.New("Token is empty")
	// ErrTokenBlocked 用户被禁止或者Token被禁止
	ErrTokenBlocked = errors.New("Token is blocked")
)

// JWT 结构体
type JWT struct {
	SigningKey []byte //签名
	Issuer     string // 签名发行者
	Timeout    int64  // 过期时间
	BufferTime int64  // 缓冲时间
}

// Token 返回token结构
type Token struct {
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}

// UserClaims  用户信息载体结构
type UserClaims struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Usercode string `json:"usercode"`
}

// MarshalBinary json转化
func (s *UserClaims) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

// UnmarshalBinary json转化
func (s *UserClaims) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}

// JWTClaims JWT 参数
type JWTClaims struct {
	UserClaims
	BufferTime int64
	jwt.StandardClaims
}

// NewJWT 生成JWT实例（参数初始化）
func NewJWT() *JWT {
	return &JWT{
		[]byte(config.JWTConfig.Secret),
		config.ApplicationConfig.Name,
		config.JWTConfig.Timeout,
		config.JWTConfig.BufferTime,
	}
}

// CreateToken 生成Token
func (j *JWT) CreateToken(userClaims *UserClaims) (*Token, error) {
	claims := &JWTClaims{
		UserClaims: UserClaims{
			UUID:     userClaims.UUID,
			Username: userClaims.Username,
			Usercode: userClaims.Usercode,
		},
		BufferTime: j.BufferTime,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: time.Now().Unix() + j.Timeout,
			// 指定token发行人
			Issuer: j.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(j.SigningKey)
	nowToken := &Token{
		Token:  token,
		Expire: claims.StandardClaims.ExpiresAt,
	}

	// 把 User信息 存储在redis内
	err = global.Redis.Set(token, userClaims, time.Duration(j.Timeout)*time.Second).Err()

	return nowToken, err
}

// ParseToken 解析 token
func (j *JWT) ParseToken(tokenString string) (*JWTClaims, error) {
	var jwtClaims = new(JWTClaims)
	token, err := jwt.ParseWithClaims(tokenString, jwtClaims, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				if claims, ok := token.Claims.(*JWTClaims); ok {
					return claims, ErrTokenExpired
				}
				return nil, ErrTokenExpired
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, ErrTokenInvalid
	}
	return nil, ErrTokenInvalid
}
