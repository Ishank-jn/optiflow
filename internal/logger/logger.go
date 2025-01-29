package logger

import (
    "github.com/sirupsen/logrus"
    "os"
)

var Log = logrus.New()

func Init() {
    Log.SetFormatter(&logrus.TextFormatter{})
    Log.SetOutput(os.Stdout)
    Log.SetLevel(logrus.InfoLevel)
}

func Info(message string) {
    Log.Info(message)
}

func Warn(message string) {
    Log.Warn(message)
}

func Error(message string) {
    Log.Error(message)
}

func Fatal(message string) {
    Log.Fatal(message)
}