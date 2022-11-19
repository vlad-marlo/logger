package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
	"strings"
	"time"
)

const (
	JSONFormatter = "json"
	TextFormatter = "text"
)

var formatters = map[string]logrus.Formatter{
	JSONFormatter: &logrus.JSONFormatter{
		TimestampFormat:   time.RFC3339,
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		CallerPrettyfier:  defaultPrettier,
		PrettyPrint:       false,
	},
	TextFormatter: &logrus.TextFormatter{
		DisableColors:          true,
		DisableTimestamp:       false,
		FullTimestamp:          true,
		TimestampFormat:        time.RFC3339,
		DisableSorting:         false,
		DisableLevelTruncation: true,
		PadLevelText:           false,
		CallerPrettyfier:       defaultPrettier,
	},
}

type CallerPrettier func(*runtime.Frame) (function string, file string)

func defaultPrettier(f *runtime.Frame) (function string, file string) {
	function = f.Function
	stripped := strings.Split(function, "/")

	if len(stripped) >= 1 {
		function = stripped[len(stripped)-1]
	} else {
		function = ""
	}

	stripped = strings.Split(f.Function, "/")
	file = ""
	if len(stripped) < 1 {
		pkg := strings.Split(stripped[len(stripped)-1], ".")[0]
		file = fmt.Sprintf("%s/%s:%d", pkg, path.Base(f.File), f.Line)
	}
	return
}
