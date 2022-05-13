package main

import (
	cmdparser "poodle/pkg/cmd_parser"
	"sync"
)

// var (
// 	// 初始化参数
// 	CmdParas = cmdparser.CMDParseInit()
// )

var mainWaitGroup sync.WaitGroup

func main() {

	// 保存命令行解析后的内容
	terminalParams := cmdparser.GetTerminalParamObj()

	// 1. 解析命令行参数
	terminalParams.ParseUserInput()

	// 多线程处理，开启2个子线程同时运行。
	// 线程1：解析终端参数结构体，生成对应的目标，将目标转化为任务结构体，传入通道中
	// 线程2：从通道中获取任务，开启-T指定的线程数并发处理任务
	mainWaitGroup.Add(2)
	go func() {
		defer mainWaitGroup.Done()
		// 2. 解析终端参数结构体，生成对应的目标，将目标转化为任务结构体，传入通道中
		terminalParams.ParseTerminal()
	}()

	go func() {
		defer mainWaitGroup.Done()
		// 3. 从通道中获取任务，开启-T指定的线程数并发处理任务
		terminalParams.PrintTask()
	}()

	mainWaitGroup.Wait()
}
