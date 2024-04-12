package logging

import (
	"log/slog"
	"os"
)

type Logger interface {
	Info(msg string, args ...any)
	Debug(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

type stdLogger struct {
	l *slog.Logger
}

func NewLogger() Logger {
	l := slog.New(slog.NewTextHandler(os.Stdout, nil))

	return &stdLogger{
		l: l,
	}
}

func (s *stdLogger) Info(msg string, args ...any) {
	s.l.Info(msg, args...)
}

func (s *stdLogger) Debug(msg string, args ...any) {
	s.l.Debug(msg, args...)
}

func (s *stdLogger) Warn(msg string, args ...any) {
	s.l.Warn(msg, args...)
}

func (s *stdLogger) Error(msg string, args ...any) {
	s.l.Error(msg, args...)
}
