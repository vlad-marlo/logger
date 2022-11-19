package hook

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
)

type OptsFunc func(h *hookWriter)

type hookWriter struct {
	lvl []logrus.Level
	out []io.Writer
}

func (h *hookWriter) Levels() []logrus.Level {
	return h.lvl
}

func (h *hookWriter) Fire(e *logrus.Entry) error {
	s, err := e.String()
	if err != nil {
		return fmt.Errorf("e.String(): %w", err)
	}
	b := []byte(s)
	for _, w := range h.out {
		_, _ = w.Write(b)
	}
	return nil
}

func New(lvl []logrus.Level, out []io.Writer, opts ...OptsFunc) logrus.Hook {
	h := &hookWriter{lvl, out}
	for _, f := range opts {
		f(h)
	}
	return h
}
