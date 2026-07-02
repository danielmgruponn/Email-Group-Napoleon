package logger

import (
	"napoleon-email/src/config/logger"

	"github.com/sirupsen/logrus"
)

type LogStruct struct {
	User   uint   `json:"user"`
	Action string `json:"action"`
	Data   any    `json:"data,omitempty"`
}

func LogInfo(message string, logData LogStruct) {
	logger.Log.WithFields(logrus.Fields{
		"action": logData.Action,
		"user":   logData.User,
		"data":   logData.Data,
	}).Info(message)
}

func LogError(message string, err error, logData LogStruct) {
    fields := logrus.Fields{
        "action": logData.Action,
        "user":   logData.User,
    }
    if err != nil {
        fields["error"] = err.Error()
    }
    logger.Log.WithFields(fields).Error(message)
}

func LogDebug(message string, logData LogStruct) {
	logger.Log.WithFields(logrus.Fields{
		"action": logData.Action,
		"user":   logData.User,
		"data":   logData.Data,
	}).Debug(message)
}

func LogWarn(message string, logData LogStruct) {
	logger.Log.WithFields(logrus.Fields{
		"action": logData.Action,
		"user":   logData.User,
	}).Warn(message)
}