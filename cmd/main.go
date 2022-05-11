package main

import (
	"poodle/pkg/common"
	"strconv"
)

func main() {
	// 单独调试
	temp := "192.168.15."
	for i := 0; i <= 255; i++ {
		ip := temp + strconv.Itoa(i)
		common.IsHostAlived(ip)
	}

	// 单IP嗅探
	// parser.Parseing(10010000, []string{"192.168.1.1", "ssd"})

	// 单域名嗅探测试
	// parser.Parseing(10020000, []string{"baidu.com"})
}
