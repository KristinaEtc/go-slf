package flogrus

import (
	"experimental/logger"
	log "github.com/Sirupsen/logrus"
	"github.com/gonotes/lfshook"
)

type LogrusLoggerFactory struct {
	logger *log.Logger
}

type LoggerFacade struct {
	logger *log.Entry
}

// TODO: add configfile
func NewLogger() (logger.LoggerFactory, error) {

	l := log.New()
	l.Hooks.Add(lfshook.NewHook(lfshook.PathMap{
		log.InfoLevel:  "info.log",
		log.ErrorLevel: "error.log",
	}))

	logFactory := LogrusLoggerFactory{logger: l}
	return &logFactory, nil
}

func (lFactory *LogrusLoggerFactory) GetLogger(name string) logger.Logger {
	logger := lFactory.logger.WithFields(log.Fields{"module": name})
	return &LoggerFacade{logger: logger}
}

//todo: change name - logrusfacade
func (lFacade *LoggerFacade) Log(level int, msg string) {
	// TODO: date, caller, time
	lFacade.logger.Printf("heh: %d %s", level, msg)
}
