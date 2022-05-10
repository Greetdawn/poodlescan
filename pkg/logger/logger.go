package logger

const (
	INFO_LEVEL int = 0
)

// 打印空行
func LogNullLine() {
	OutputNullLine()
}

// 无格式打印
func LogNoFormat(log string) {
	OutputNoFormat(log)
}

// 打印INFO日志
func LogInfo(tag, log string) {
	OutputInfo(tag, log)
}

// 打印INFO日志
func LogError(tag, log string) {
	OutputError(tag, log)
}
