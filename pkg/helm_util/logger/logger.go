package logger

import (
	"fmt"
	"github.com/ggermis/helm-util/pkg/helm_util/cli"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

var logger *logrus.Logger

func init() {
	logger = &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.InfoLevel,
		Formatter: &formatter{logrus.TextFormatter{
			FullTimestamp:          true,
			TimestampFormat:        "2006-01-02 15:04:05.00000",
			ForceColors:            true,
			DisableLevelTruncation: true,
		},
		},
	}
}

type formatter struct {
	logrus.TextFormatter
}

func (f *formatter) Format(entry *logrus.Entry) ([]byte, error) {
	if cli.Colour {
		var levelColor int
		switch entry.Level {
		case logrus.DebugLevel, logrus.TraceLevel:
			levelColor = 34 // blue
		case logrus.WarnLevel:
			levelColor = 33 // yellow
		case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
			levelColor = 31 // red
		default:
			levelColor = 32 // cyan
		}
		return []byte(fmt.Sprintf("\u001B[1;%dm[%s] - %s - %s\u001B[0m\n", levelColor, entry.Time.Format(f.TimestampFormat), strings.ToUpper(entry.Level.String()), entry.Message)), nil
	} else {
		return []byte(fmt.Sprintf("[%s] - %s - %s\n", entry.Time.Format(f.TimestampFormat), strings.ToUpper(entry.Level.String()), entry.Message)), nil
	}
}

func SetLogLevel() {
	if cli.Debug {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}
