package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func New(debug bool) *Logger {
	log := logrus.New()
	log.SetReportCaller(debug)

	// log.SetFormatter(&logrus.JSONFormatter{})
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	log.SetOutput(os.Stdout)

	if debug {
		log.SetLevel(logrus.TraceLevel)
	} else {
		log.SetLevel(logrus.InfoLevel)
	}

	return &Logger{log}
}
