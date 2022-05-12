package main

import (
	cmdparser "poodle/internal/cmd_parser"
	"poodle/pkg/common"
	"sync"
)

var WG sync.WaitGroup

func main() {

	// 初始化参数
	CmdParas := cmdparser.CMDParseInit()
	// 解析命令行参数
	CmdParas.CMDUserInputParse()
	//fmt.Println(CmdParas.IpList)

	// 生成目标
	CmdParas.TargetChan = make(chan cmdparser.TargetInput)
	WG.Add(1)
	go func() {
		CmdParas.ProduceTargets()
		close(CmdParas.TargetChan)
		WG.Done()
	}()

	// 全局并发控制

	for i := 0; i <= CmdParas.Threads; i++ {
		for targetTest := range CmdParas.TargetChan {
			WG.Add(1)
			go func(targetTest cmdparser.TargetInput) {
				WG.Add(1)
				go func() {
					common.ScanHostAlived(targetTest)
					WG.Done()
				}()
				WG.Done()
			}(targetTest)
		}
	}

	WG.Wait()

}
