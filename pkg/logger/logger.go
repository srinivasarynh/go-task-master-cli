package logger

import (
	"os"
	"taskmaster/internal/config"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func Init() {
	Logger = logrus.New()

	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	Logger.SetLevel(logrus.InfoLevel)
}

func ConfigureLogger() error {
	if config.AppConfig == nil {
		return nil
	}

	level, err := logrus.ParseLevel(config.AppConfig.Logging.Level)
	if err != nil {
		level = logrus.InfoLevel
	}

	Logger.SetLevel(level)

	if config.AppConfig.Logging.File != "" {
		file, err := os.OpenFile(config.AppConfig.Logging.File, os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		Logger.SetOutput(file)
	}
	return nil
}

func GetLogger() *logrus.Logger {
	return Logger
}
