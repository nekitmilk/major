package logger

import (
	"os"

	logrusFormatter "github.com/t-tomalak/logrus-easy-formatter"

	"github.com/sirupsen/logrus"
)

func SetLogger(debug bool) *logrus.Logger {
	var lvl logrus.Level

	if debug {
		lvl = logrus.DebugLevel
	} else {
		lvl = logrus.InfoLevel
	}

	logger := &logrus.Logger{
		Out:   os.Stderr,
		Level: lvl,
		Formatter: &logrusFormatter.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%time%] [%lvl%] %msg%\n",
		},
	}

	return logger
}
