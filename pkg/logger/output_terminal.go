package logger

import (
	"fmt"
	"time"
)

// 打印空行
func OutputNullLine() {
	fmt.Println()
}

// 无格式打印
func OutputNoFormat(log string) {
	fmt.Print(log)
}

// 打印INFO日志
func OutputInfo(tag, log string) {
	fmt.Printf("[INFO] [%s] %s %s\n", tag, time.Now().Format("2006-01-02 15:04:05"), log)
}

// 打印INFO日志
func OutputError(tag, log string) {
	fmt.Printf("[ERROR] [%s] %s %s\n", tag, time.Now().Format("2006-01-02 15:04:05"), log)
}
