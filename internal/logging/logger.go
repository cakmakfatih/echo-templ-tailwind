package logging

import (
	"log/slog"
	"os"
)

type Logger interface{}

type stdLogger struct {
	l *slog.Logger
}

func NewLogger() Logger {
	l := slog.New(slog.NewTextHandler(os.Stdout, nil))

	return &stdLogger{
		l: l,
	}
}
