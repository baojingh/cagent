package config

import (
	"agent-server/constant"

	"github.com/spf13/viper"
)

var (
	LogCount   int
	LogTime    int
	LogSize    int
	LogPath    string
	CertPath   string
	ConfigPath string
	Ip         string
	Port       int
	TLSEnabled bool
)

// Init config values and overite values from config file
func init() {
	v := viper.New()
	// The config file location cannot be changed.
	v.SetConfigFile(constant.DEFAULT_CONF_PATH)
	err := v.ReadInConfig()
	if err != nil {
		return
	}
	Ip = v.GetString("service.ip")
	Port = v.GetInt("service.port")
	LogPath = v.GetString("log.path")
	LogCount = v.GetInt("logrotate.count")
	LogSize = v.GetInt("logrotate.fileSize")
	LogTime = v.GetInt("logrotate.time")
	CertPath = v.GetString("cert.path")
	TLSEnabled = v.GetBool("tls.enabled")
}
