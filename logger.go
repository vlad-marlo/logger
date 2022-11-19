package logger

import (
	"github.com/sirupsen/logrus"
)

// WithOpts ...
func WithOpts(opts ...OptFunc) *logrus.Logger {
	l := logrus.New()
	for _, o := range opts {
		o(l)
	}
	return l
}

// AddOpts ...
func AddOpts(l *logrus.Logger, opts ...OptFunc) {
	for _, o := range opts {
		o(l)
	}
}
