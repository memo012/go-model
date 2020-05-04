package demo

import "fmt"

type ErrorLogger struct {
	abstractLogger AbstractLogger
}

func (c *ErrorLogger) write(message string)  {
	fmt.Println("consoleLogger :" + message)
}

func (c *ErrorLogger) LogMessage(level int, message string, logger Logger)  {
	c.abstractLogger.LogMessage(level, message, c)
}

