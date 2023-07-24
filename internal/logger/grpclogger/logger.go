package grpclogger

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/sirupsen/logrus"
)

func InterceptorLogger(log logrus.FieldLogger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		f := make(map[string]any, len(fields)/2)
		i := logging.Fields(fields).Iterator()

		for i.Next() {
			k, v := i.At()
			f[k] = v
		}

		log := log.WithFields(f)

		switch lvl {
		case logging.LevelDebug:
			log.Debug(msg)

		case logging.LevelInfo:
			log.Info(msg)

		case logging.LevelWarn:
			log.Warn(msg)

		case logging.LevelError:
			log.Error(msg)

		default:
			log.Fatalf("unknown level %v", lvl)
		}

	})
}
