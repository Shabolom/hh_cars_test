package tools

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
	"hh_test_autho/config"
	"os"
	"runtime"
)

var InfoLog *log.Logger

func InitLogger() error {

	logsFile, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	defer logsFile.Close()

	log.SetReportCaller(true)

	if config.Env.Production {
		log.SetLevel(log.WarnLevel)
		log.SetOutput(logsFile)
		log.SetFormatter(&nested.Formatter{
			ShowFullLevel: true,
			HideKeys:      true,
			FieldsOrder:   []string{"component", "category"},
		})
	} else {
		log.SetLevel(log.DebugLevel)
		log.SetOutput(os.Stdout)
		log.SetFormatter(&nested.Formatter{
			ShowFullLevel: true,
			HideKeys:      true,
			FieldsOrder:   []string{"component", "category"},
		})
	}
	return nil
}

func InfoLogs() *log.Logger {
	infoLog := log.New()
	infoLog.SetLevel(log.DebugLevel)
	infoLog.SetOutput(os.Stdout)
	infoLog.SetFormatter(&nested.Formatter{
		ShowFullLevel: true,
		HideKeys:      true,
		FieldsOrder:   []string{"component", "middleware"},
		CustomCallerFormatter: func(frame *runtime.Frame) string {
			return ""
		},
	})

	InfoLog = infoLog

	return infoLog
}
