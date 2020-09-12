package config

import "github.com/spf13/viper"

// LoggerConfig LoggerConfig
var LoggerConfig = new(Logger)

// Logger Logger struct
type Logger struct {
	Path       string
	Level      string
	Stdout     bool
	EnabledBus bool
	EnabledReq bool
	EnabledJob bool
}

// InitLogger Logger 结构初始化
func InitLogger(cfg *viper.Viper) *Logger {
	return &Logger{
		Path:       cfg.GetString("path"),
		Level:      cfg.GetString("level"),
		Stdout:     cfg.GetBool("stdout"),
		EnabledBus: cfg.GetBool("enabledbus"),
		EnabledReq: cfg.GetBool("enabledreq"),
		EnabledJob: cfg.GetBool("enabledjob"),
	}
}
