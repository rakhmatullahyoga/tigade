package tool

import (
	"go.uber.org/zap"
)

type LogAdapter struct {
	*zap.Logger
}

func NewLogger(env string) *LogAdapter {
	var l *zap.Logger
	switch env {
	case "development":
		l, _ = zap.NewDevelopment()
	default:
		l, _ = zap.NewProduction()
	}
	return &LogAdapter{l}
}

func (logger *LogAdapter) LogInfo(trackId string, tags []string, message string) {
	logger.Info(message, zap.String("track_id", trackId),
		zap.Strings("tags", tags))
}

func (logger *LogAdapter) LogError(trackId string, tags []string, message string, err error, code int) {
	logger.Error(err.Error(), zap.String("track_id", trackId),
		zap.Strings("tags", tags),
		zap.Int("code", code),
		zap.String("message", message))
}

func (logger *LogAdapter) LogPanic(trackId string, tags []string, message string, err error, code int) {
	logger.Panic(err.Error(), zap.String("track_id", trackId),
		zap.Strings("tags", tags),
		zap.Int("code", code),
		zap.String("message", message))
}
