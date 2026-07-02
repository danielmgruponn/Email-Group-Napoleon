package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()

	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	if err := os.MkdirAll("src/pkg/logs", 0755); err != nil {
		os.Stderr.WriteString("WARNING: Could not create log directory, logging to stdout only: " + err.Error() + "\n")
		Log.SetOutput(os.Stdout)
	} else {
		logFile := &lumberjack.Logger{
			Filename:   "src/pkg/logs/app.log",
			MaxSize:    10,
			MaxBackups: 7,
			MaxAge:     30,
		}

		multiWriter := io.MultiWriter(os.Stdout, logFile)
		Log.SetOutput(multiWriter)
	}

	Log.SetLevel(logrus.DebugLevel)

	Log.Info("Logger Initialization")
}
