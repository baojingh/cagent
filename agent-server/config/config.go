package config

import (
	"agent-server/config"
	"agent-server/utils"

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

// Setup values customized by users
func SetupConfig(cert string, logPath string, tlsEnabled bool,
	configPath string) {
	if utils.IsEmpty(configPath) {
		ConfigPath = config.DEFAULT_CONF_PATH
	}
	v := viper.New()
	// The config file location cannot be changed.
	v.SetConfigFile(configPath)
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

	if !utils.IsEmpty(cert) {
		CertPath = cert
	}
	if !utils.IsEmpty(logPath) {
		LogPath = logPath
	}

	TLSEnabled = tlsEnabled
}
