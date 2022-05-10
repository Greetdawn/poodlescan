package common

import "runtime"

// 判断当前程序运行的系统类型
func JudgeSystemType() string {
	winenv := "cmd"
	linuxenv := "/bin/bash"
	OS := runtime.GOOS
	if OS == "linux" || OS == "darwin" {
		return linuxenv
	} else if OS == "windows" {
		return winenv
	}
	return ""
}
