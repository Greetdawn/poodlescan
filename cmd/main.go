package main

import (
	"poodle/pkg/common"
	"poodle/pkg/parser"
)

func main() {
	// 单独调试
	common.IsHostAlived("192.168.1.1")

	// 单IP嗅探
	parser.Parseing(10010000, []string{"192.168.1.1"})

	// 	// 单域名嗅探测试
	parser.Parseing(10020000, []string{"baidu.com"})
}
