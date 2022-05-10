package common

import "fmt"

// 判断主机是否存活
// host可以传入IP、域名
func IsHostAlived(host string) bool {
	// 运行前初始话当前系统类型
	OS := JudgeSystemType()
	fmt.Println(OS)
	return true
}
