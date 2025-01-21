package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

const logFolderPath = "./../log/"

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	createLogFolderIfNotExists()
	fileName := createLogFileName()
	file, err := os.Create(logFolderPath + fileName)
	if err != nil {
		Log.Errorln("Error when create log folder, %q", err.Error())
	}
	mw := io.MultiWriter(os.Stdout, file)
	Log.SetOutput(mw)
	Log.SetFormatter(
		&logrus.TextFormatter{
			FullTimestamp: true,
		})
	Log.SetLevel(logrus.TraceLevel)
}

func createLogFolderIfNotExists()  {
	_, err := os.Stat(logFolderPath)
	if os.IsNotExist(err) {
		err := os.Mkdir(logFolderPath, os.ModePerm)
		if err != nil {
			Log.Infoln("Error create log folder")
		}
	} else if err != nil {
		Log.Errorf("Error checking log folder: %q", err.Error())
		return
	} else {
		Log.Infoln("Log folder already exists")
	}
}

func createLogFileName() string {
	now := time.Now()
	return fmt.Sprintf(
		"log_%d-%02d-%02d_%02d-%02d-%02d.log",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
}
