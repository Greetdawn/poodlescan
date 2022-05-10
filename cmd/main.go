package main

import "poodle/pkg/parser"

func main() {
	// 单IP嗅探
	parser.Parseing(10010000, []string{"192.168.1.1", "ssd"})

	// 	// 单域名嗅探测试
	parser.Parseing(10020000, []string{"baidu.com"})
}
