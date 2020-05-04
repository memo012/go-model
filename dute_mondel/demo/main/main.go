package main

import (
	"designMode/dute_mondel/demo"
	"designMode/dute_mondel/demo/enum"
)

func main() {
	logger := demo.GetLoggerFactory()

	logger.LogMessage(enum.CONSOLE, "this is console", nil)

	logger.LogMessage(enum.ERROR, "this is error", nil)

	logger.LogMessage(enum.FILE, "this is file", nil)

}
