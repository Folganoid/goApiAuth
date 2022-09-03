package api

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func NewLogger(logFile *os.File, logLevel string) (*log.Logger, error) {

	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{})
	logger.SetOutput(logFile)

	if logLevel == "debug" {
		logger.SetLevel(log.DebugLevel)
	}
	if logLevel == "info" {
		logger.SetLevel(log.InfoLevel)
	}

	//....

	return logger, nil
}
