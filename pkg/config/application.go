package config

import "github.com/spf13/viper"

// ApplicationConfig init
var ApplicationConfig = new(Application)

// Application struct
type Application struct {
	ReadTimeout   int
	WriterTimeout int
	Host          string
	Port          string
	APIVersion    string
	Name          string
	Mode          string
	EnableDP      bool
}

// InitApplication get config
func InitApplication(cfg *viper.Viper) *Application {
	return &Application{
		ReadTimeout:   cfg.GetInt("readTimeout"),
		WriterTimeout: cfg.GetInt("writerTimeout"),
		Host:          cfg.GetString("host"),
		Port:          cfg.GetString("port"),
		APIVersion:    cfg.GetString("apiversion"),
		Name:          cfg.GetString("name"),
		Mode:          cfg.GetString("mode"),
		EnableDP:      cfg.GetBool("enabledp"),
	}
}
