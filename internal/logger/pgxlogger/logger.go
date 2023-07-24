package pgxlogger

import (
	"context"

	"github.com/jackc/pgx/v5/tracelog"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	l logrus.FieldLogger
}

func NewLogger(l logrus.FieldLogger) *Logger {
	return &Logger{l: l}
}

func (l *Logger) Log(_ context.Context, level tracelog.LogLevel, msg string, data map[string]interface{}) {
	var logger logrus.FieldLogger
	if data != nil {
		logger = l.l.WithFields(data)
	} else {
		logger = l.l
	}

	logger = logger.WithField("protocol", "postgres")

	switch level {
	case tracelog.LogLevelTrace:
		logger.WithField("PGX_LOG_LEVEL", level).Trace(msg)

	case tracelog.LogLevelDebug:
		logger.Debug(msg)

	case tracelog.LogLevelInfo:
		logger.Info(msg)

	case tracelog.LogLevelWarn:
		logger.Warn(msg)

	case tracelog.LogLevelError:
		logger.Error(msg)

	default:
		logger.WithField("INVALID_PGX_LOG_LEVEL", level).Error(msg)
	}
}
