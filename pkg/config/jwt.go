package config

import "github.com/spf13/viper"

// JWTConfig init
var JWTConfig = new(JWT)

// JWT struct
type JWT struct {
	Secret     string
	Timeout    int64
	BufferTime int64
	HeaderName string
}

// InitJWT get config
func InitJWT(cfg *viper.Viper) *JWT {
	return &JWT{
		Secret:     cfg.GetString("secret"),
		Timeout:    cfg.GetInt64("timeout"),
		BufferTime: cfg.GetInt64("bufferTime"),
		HeaderName: cfg.GetString("headerName"),
	}
}
