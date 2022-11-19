package logger

import (
	"github.com/sirupsen/logrus"
	"io"
)

// OptFunc ...
type OptFunc func(l *logrus.Logger)

// WithOutput ...
func WithOutput(out io.Writer) OptFunc {
	return func(l *logrus.Logger) {
		l.SetOutput(out)
	}
}

// WithHook ...
func WithHook(h logrus.Hook) OptFunc {
	return func(l *logrus.Logger) {
		l.AddHook(h)
	}
}

// WithHooks ...
func WithHooks(hooks ...logrus.Hook) OptFunc {
	return func(l *logrus.Logger) {
		for _, h := range hooks {
			l.AddHook(h)
		}
	}
}

// WithFormatter ...
func WithFormatter(f logrus.Formatter) OptFunc {
	return func(l *logrus.Logger) {
		l.SetFormatter(f)
	}
}

// WithLevel ...
func WithLevel(lvl logrus.Level) OptFunc {
	return func(l *logrus.Logger) {
		l.SetLevel(lvl)
	}
}

// WithReportCaller ...
func WithReportCaller(report bool) OptFunc {
	return func(l *logrus.Logger) {
		l.SetReportCaller(report)
	}
}

// WithBufferPool ...
func WithBufferPool(poll logrus.BufferPool) OptFunc {
	return func(l *logrus.Logger) {
		l.SetBufferPool(poll)
	}
}

// WithDefaultFormatter ...
func WithDefaultFormatter(formatter string) OptFunc {
	return func(l *logrus.Logger) {
		switch formatter {
		case JSONFormatter, TextFormatter:
			l.SetFormatter(formatters[formatter])
		default:
			l.SetFormatter(formatters[TextFormatter])
		}
	}
}
