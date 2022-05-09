package common

import (
	"fmt"
	"time"
)

// 打印INFO日志
func PrintInfoLog(log string) {
	fmt.Println()
	fmt.Printf("[INFO] %s %s\n", time.Now().Format("2006-01-02 15:04:05"), log)
}
