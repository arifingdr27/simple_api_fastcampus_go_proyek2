package helpers

import (
	"github.com/sirupsen/logrus"
)

func SetupLogger() *logrus.Logger {
	logger := logrus.New()
	// logger.SetFormatter(&log.JSONFormatter{})
	logrus.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	logger.Info("arifin")
	return logger

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	// log.SetOutput(os.Stdout)

	// // Only log the warning severity or above.
	// log.SetLevel(log.WarnLevel)
	// logger.WithFields(logrus.Fields{
	// 	"animal": "walrus",
	// 	"size":   10,
	// }).Info("A group of walrus emerges from the ocean")
}
