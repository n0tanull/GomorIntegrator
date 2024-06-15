package melog

import (
	"github.com/gookit/slog"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
}
type logger struct {
	slog slog.Logger
}

var _ Logger = (*logger)(nil)

func New() Logger {
	return &logger{
		slog: slog.Logger{},
	}
}
func (l *logger) Debug(msg string) {
	l.slog.Debug(msg)
}
func (l *logger) Info(msg string) {
	l.slog.Info(msg)
}
func (l *logger) Warn(msg string) {
	l.slog.Warn(msg)
}
func (l *logger) Error(msg string) {
	l.slog.Error(msg)
}
func (l *logger) Fatal(msg string) {
	l.slog.Fatal(msg)
}
