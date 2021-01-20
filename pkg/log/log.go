package log

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {

	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05.000"
	customFormatter.FullTimestamp = true
	customFormatter.DisableLevelTruncation = true

	log.SetFormatter(customFormatter)

	logLevelStr, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		logLevelStr = "debug"
	}

	logLevel, err := log.ParseLevel(logLevelStr)
	if err != nil {
		logLevel = log.DebugLevel
	}

	log.SetLevel(logLevel)
}
