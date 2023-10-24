package config

import (
	logger "agent-server/log"

	"github.com/spf13/viper"
)

var log = logger.New()

var (
	LogCount int
	LogTime  int
	LogSize  int
	LogPath  string
	CertPath string
	Ip       string
	Port     int
)

func init() {

}

// Init config values and overite values from config file
func InitConfig(cert string, logPath string, config string) {
	viper.SetConfigFile("setting/agent-server.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Error("Failed to read config file,", err)
		return
	}
	err = viper.Unmarshal()
	if err != nil {
		log.Error("Failed to set config object, ", err)
	}
}
