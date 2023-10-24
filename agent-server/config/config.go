package config

import (
	logger "agent-server/log"
	"agent-server/utils"

	"github.com/spf13/viper"
)

var log = logger.New()

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
	v.SetConfigFile("agent-server/setting/agent-server.yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Error("Failed to read config file,", err)
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
	log.Infof("Default config value is loaded.")
}

// Setup values customized by users
func SetupConfig(cert string, logPath string, config string, tlsEnabled bool) {
	if !utils.IsEmpty(cert) {
		CertPath = cert
		log.Infof("Setup cert path %s from cmd line.", CertPath)
	}
	if !utils.IsEmpty(logPath) {
		LogPath = logPath
		log.Infof("Setup log path %s from cmd line.", LogPath)
	}
	if !utils.IsEmpty(config) {
		ConfigPath = config
		log.Infof("Setup config path %s from cmd line.", ConfigPath)
	}
	TLSEnabled = tlsEnabled
	log.Infof("LogPath: %s, LogCount:%d", LogPath, LogCount)

	log.Infof(`LogPath: %s, LogCount:%d, 
		LogSize: %d, LogTime: %d, CertPath :%s, TLSEnabled: %v`,
		LogPath, LogCount, LogSize, LogTime,
		CertPath, TLSEnabled)
	log.Infof(`LogPath: %s, LogCount:%d, 
	LogSize: %d, LogTime: %d, CertPath :%s, TLSEnabled: %v`,
		LogPath, LogCount, LogSize, LogTime,
		CertPath, TLSEnabled)
}
