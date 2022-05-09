package main

import (
	"poodle/pkg/parser"
)

func main() {
	// 单IP嗅探
	parser.Parseing(10100, []string{"192.168.1.1"})
}
