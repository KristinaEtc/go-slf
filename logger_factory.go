package logger

import (
	"log"
	"os"
)

type Logger interface {
	// TODO: string + interface... instead of string
	// return
	Log(int, string)
}

type CustomLoggerFactory struct {
}

type CustomLoggerFacade struct {
	logger *log.Logger
}

func init() {
	factory = &CustomLoggerFactory{}
}

func (cFactory *CustomLoggerFactory) GetLogger(name string) Logger {
	logger := log.New(os.Stdout, "", log.Ldate|log.Lshortfile)
	l := CustomLoggerFacade{logger: logger}
	return &l
}

func (cFacade *CustomLoggerFacade) Log(level int, msg string) {
	cFacade.logger.Printf("heh-log: %d %s", level, msg)
}

// Implement a factory of loggers.
type LoggerFactory interface {
	// GetLogger based on the given library name, where would be used a logger
	// Returns logger pointer, nil (will used a custom golang log library) otherwise.
	GetLogger(name string) Logger
}

var factory LoggerFactory

func GetLogger(name string) Logger {
	return factory.GetLogger(name)
}

func InitLogger(l LoggerFactory) {
	factory = l
}
