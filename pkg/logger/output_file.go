package logger

import (
	"fmt"
	"os"
	"time"
)

var LoggerFilePt *os.File

func init() {
	OutputInfo("logger", "Init logger module")
	// 这里创建文件
	// LoggerFilePt, err := os.Create("log.txt")
}

// 打印空行
func FoutputNullLine() {
	fmt.Println()
}

// 无格式打印
func FoutputNoFormat(log string) {
	fmt.Print(log)
}

// 打印INFO日志
func FoutputInfo(tag, log string) {
	fmt.Printf("[INFO] [%s] %s %s\n", tag, time.Now().Format("2006-01-02 15:04:05"), log)
}

// 打印Error日志
func FoutputError(tag, log string) {
	fmt.Printf("[ERROR] [%s] %s %s\n", tag, time.Now().Format("2006-01-02 15:04:05"), log)
}
