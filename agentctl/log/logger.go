package logger

import (
	"agentctl/constant"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	rotate "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// init logs
func init() {
	path := constant.DEFAULT_LOG_PATH

	writer, err := rotate.New(
		filepath.Join(path, fmt.Sprintf("agentctl-%s.log", "%Y%m%d")),
		rotate.WithLinkName(filepath.Join(path, "agentctl.log")),
		rotate.WithRotationCount(uint(constant.DEFAULT_LOG_COUNT)),
		rotate.WithRotationTime(time.Hour*time.Duration(constant.DEFAULT_LOG_TIME)),
	)
	if err == nil {
		log.SetOutput(io.MultiWriter(writer, os.Stdout))
	} else {
		log.Fatal("Config logger failure, ", err)
	}

	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:           "2006-01-02 15:04:05",
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
		DisableLevelTruncation:    true,
	})
	log.SetLevel(logrus.InfoLevel)
}

func New() *logrus.Logger {
	return log
}
