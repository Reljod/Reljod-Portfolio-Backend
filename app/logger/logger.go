package logger

import (
	"os"
	"strings"

	"github.com/Reljod/Reljod-Portfolio-Backend/app/environ"
	log "github.com/sirupsen/logrus"
)

var Logger *log.Logger
var env environ.Environ

func init() {
	env.Setup()

	Logger = log.New()
	Logger.SetOutput(os.Stdout)
	Logger.SetFormatter(&log.TextFormatter{})

	logLevel := env.LogLevel
	if logLevel == "" {
		logLevel = "INFO"
	}

	setLogLevel(Logger, logLevel)
}

func setLogLevel(l *log.Logger, logLevel string) {

	upperStrLevel := strings.ToUpper(logLevel)
	if upperStrLevel == "INFO" {
		l.SetLevel(log.InfoLevel)
	} else if upperStrLevel == "DEBUG" {
		l.SetLevel(log.DebugLevel)
	} else {
		l.SetLevel(log.WarnLevel)
	}
}
