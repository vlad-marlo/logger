package hook

import "github.com/sirupsen/logrus"

func WithFileOutput(dir, entry, name string) OptsFunc {
	return func(h *hookWriter) {

		f, err := createOrGetFile(dir, entry, name)
		if err != nil {
			return
		}
		h.out = append(h.out, f)
	}
}

func WithLevels(levels []string) OptsFunc {
	return func(h *hookWriter) {
		var lvl []logrus.Level
		for _, l := range levels {
			if level, err := logrus.ParseLevel(l); err == nil {
				lvl = append(lvl, level)
			}
		}
		h.lvl = lvl
	}
}
