package helpers

import (
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func SetupLogger() {
	log := logrus.New()
	logrus.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	Logger = log
}
