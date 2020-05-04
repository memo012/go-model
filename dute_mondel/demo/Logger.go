package demo

type Logger interface {
	LogMessage(level int, message string, logger Logger)
	write(message string) // 处理
}
