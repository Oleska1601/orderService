package logger

import (
	"log/slog"
	"os"
	"strings"
)

type LoggerInterface interface {
	Debug(string, ...any)
	Error(string, ...any)
	Info(string, ...any)
	Warn(string, ...any)
}

type Logger struct {
	logger *slog.Logger
}

var _ LoggerInterface = (*Logger)(nil)

func NewLogger(level string) *Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	var loggerLevel slog.Level
	switch strings.ToLower(level) {
	case "debug":
		loggerLevel = slog.LevelDebug
	case "info":
		loggerLevel = slog.LevelInfo
	case "warn":
		loggerLevel = slog.LevelWarn
	case "error":
		loggerLevel = slog.LevelError
	default:
		loggerLevel = slog.LevelInfo
	}

	slog.SetLogLoggerLevel(loggerLevel)
	slog.SetDefault(logger)
	return &Logger{
		logger: logger,
	}

}

func (l *Logger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}
