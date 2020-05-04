package demo

import "fmt"

type FileLogger struct {
	abstractLogger AbstractLogger
}

func (c *FileLogger) write(message string)  {
	fmt.Println("consoleLogger :" + message)
}

func (c *FileLogger) LogMessage(level int, message string, logger Logger)  {
	c.abstractLogger.LogMessage(level, message, c)
}
