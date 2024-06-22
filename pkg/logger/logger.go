package logger

import (
	"os"
    prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
    Logger = NewLogger()        
}

// creates and returns a new logger instance
func NewLogger() *logrus.Logger {
	logger := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.InfoLevel,
		Formatter: &prefixed.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		},
	}

    return logger
}

func getLogEntryWithFields(err error, message string, info map[string]interface{}) *logrus.Entry {
    entry := Logger.WithFields(logrus.Fields{
        "error":   err,
        "msg": message,
    })
    for key, value := range info {
        entry = entry.WithField(key, value)
    }
    
    return entry
}

func LogError(err error, message string, info map[string]interface{}) {
    entry := getLogEntryWithFields(err, message, info)
    entry.Error("")
}

func LogFatal(err error, message string, info map[string]interface{}) {
    entry := getLogEntryWithFields(err, message, info)
    entry.Fatal("")
}

func LogInfo(err error, message string, info map[string]interface{}) {
    entry := getLogEntryWithFields(err, message, info)
    entry.Info("")
}

func LogWarn(err error, message string, info map[string]interface{}) {
    entry := getLogEntryWithFields(err, message, info)
    entry.Warn("")
} 
