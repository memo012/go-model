package demo

import "designMode/dute_mondel/demo/enum"

func GetLoggerFactory() Logger {
	errorLogger := new(ErrorLogger)
	errorLogger.abstractLogger.Level = enum.ERROR

	fileLogger := new(FileLogger)
	fileLogger.abstractLogger.Level = enum.FILE

	consoleLogger := new(ConsoleLogger)
	consoleLogger.abstractLogger.Level = enum.CONSOLE

	consoleLogger.abstractLogger.SetNextLogger(errorLogger)
	errorLogger.abstractLogger.SetNextLogger(fileLogger)

	return consoleLogger
}
