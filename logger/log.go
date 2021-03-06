package logger

type Log interface {
	// 初始化
	Init()

	// 设置日志级别
	SetLevel(level int)

	// 各级别日志
	Debug(format string, args ...interface{})
	Trace(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
}
