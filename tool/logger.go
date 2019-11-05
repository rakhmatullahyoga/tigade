package tool

import (
	"go.uber.org/zap"
)

type Logger interface {
	LogInfo(trackId string, tags []string, message string)
	LogError(trackId string, tags []string, message string, error error, code int)
	LogPanic(trackId string, tags []string, message string, error error, code int)
}

type LogAdapter struct {
	*zap.Logger
}

func NewLogger(env string) *LogAdapter {
	var l *zap.Logger
	switch env {
	case "production":
		l, _ = zap.NewProduction()
	default:
		l, _ = zap.NewDevelopment()
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
