package main

import (
	"fmt"
	"os"
	"poodle/pkg/asset_host"
	cmdparser "poodle/pkg/cmd_parser"
	"poodle/pkg/common"
	controllor "poodle/pkg/controller"
	"poodle/pkg/logger"
	"sync"
)

var mainWaitGroup sync.WaitGroup

func main() {

	// 1. 解析用户输入的命令行
	var err error

	// 输出banner信息
	common.ShowBanner()

	err = cmdparser.ParsingUserTerminalLine(&cmdparser.G_TerminalParam)
	if err != nil {
		// 解析用户的输入目标发生错误。
		logger.LogError(fmt.Sprintf("err: %v\n", err), logger.LOG_TERMINAL_FILE)
		return
	}
	// 2. 生成CC 控制码
	cc := cmdparser.G_TerminalParam.GenerateControlCode()

	// 多线程处理，开启2个子线程同时运行。
	// 线程1：解析终端参数结构体，生成对应的目标，将目标转化为任务结构体，传入通道中
	// 线程2：从通道中获取任务，开启-T指定的线程数并发处理任务
	mainWaitGroup.Add(2)
	go func() {
		defer mainWaitGroup.Done()
		// 3. 解析终端参数结构体，生成对应的目标，将目标转化为任务结构体，传入通道中
		err := cmdparser.GenrateTasks(&cmdparser.G_TerminalParam, cc)
		if err != nil {
			// 解析用户的输入目标发生错误。
			logger.LogError(fmt.Sprintf("err: %v\n", err), logger.LOG_TERMINAL_FILE)
			os.Exit(0)
			return
		}
	}()

	go func() {
		defer mainWaitGroup.Done()
		// 4. 从通道中获取任务，开启-T指定的线程数并发处理任务
		controllor.Run(cmdparser.G_TerminalParam.ThreadsNumber)
	}()

	mainWaitGroup.Wait()
	asset_host.GetSnifferObj().PrintAssetHostList()
}
