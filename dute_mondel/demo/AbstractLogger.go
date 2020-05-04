package demo

type AbstractLogger struct {
	Level      int
	NextLogger Logger
}

// 设置下一个handler
func (al *AbstractLogger) SetNextLogger(logger Logger) {
	al.NextLogger = logger
}

func (al *AbstractLogger) LogMessage(level int, message string, logger Logger) {
	if al.Level <= level {
		logger.write(message)
	} else {
		if al.NextLogger != nil {
			al.NextLogger.LogMessage(level, message, al.NextLogger)
		}
	}
}
