package common

import (
	"fmt"
	"time"
)

// 打印空行
func PrintNullLine() {
	fmt.Println()
}

// 无格式打印
func Print(log string) {
	fmt.Print(log)
}

// 打印INFO日志
func PrintInfoLog(log string) {
	fmt.Printf("[INFO] %s %s\n", time.Now().Format("2006-01-02 15:04:05"), log)
}

// 打印INFO日志
func PrintErrorLog(log string) {
	fmt.Printf("[ERROR] %s %s\n", time.Now().Format("2006-01-02 15:04:05"), log)
}
