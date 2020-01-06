package main

import (
	"github.com/bhbosman/CodeGenerators/idlgenerator/AppImpl"
	"go.uber.org/fx"
	"log"
)

type loggerWrapper struct {
	applicationLogger *log.Logger
}

func (self *loggerWrapper) Printf(format string, v ...interface{}) {
	self.applicationLogger.Printf(format, v...)
}

func FxAppProvideFxAppOverrideLogger(logger *log.Logger) fx.Option {
	return fx.Logger(&loggerWrapper{
		applicationLogger: logger,
	})
}

type LoggerInstance struct {
	logger *log.Logger
}

func NewLoggerInstance(logger *log.Logger) *LoggerInstance {
	return &LoggerInstance{
		logger: logger,
	}
}

type LogFactory struct {
	logger *LoggerInstance
}

func NewLogFactory(logger *log.Logger) *LogFactory {
	return &LogFactory{
		logger: NewLoggerInstance(logger),
	}
}

func (l LogFactory) Create() *log.Logger {
	return l.logger.logger
}

func FxAppProvideApplicationLogger(logger *log.Logger) fx.Option {
	return fx.Provide(
		func() (AppImpl.ILogFactory, error) {
			return NewLogFactory(logger), nil
		})
}
