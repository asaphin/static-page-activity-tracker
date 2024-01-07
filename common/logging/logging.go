package logging

import (
	"github.com/rs/zerolog"
	"os"
	"strings"
)

var logLevel = strings.ToLower(os.Getenv("LOG_LEVEL"))

var logLevelsMap = map[string]zerolog.Level{
	"trace":   zerolog.TraceLevel,
	"debug":   zerolog.DebugLevel,
	"info":    zerolog.InfoLevel,
	"warning": zerolog.WarnLevel,
	"warn":    zerolog.WarnLevel,
	"error":   zerolog.ErrorLevel,
	"fatal":   zerolog.FatalLevel,
	"panic":   zerolog.PanicLevel,
}

func Setup() {
	if lvl, ok := logLevelsMap[logLevel]; ok {
		zerolog.SetGlobalLevel(lvl)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}
