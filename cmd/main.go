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
		WG.Add(1)
		go func() {
			defer WG.Done()
			for targetTest := range CmdParas.TargetChan {

				var childwg sync.WaitGroup
				childwg.Add(1)

				go func(targetTest cmdparser.TargetInput) {
					common.ScanHostAlived(targetTest)
					childwg.Done()
				}(targetTest)

				childwg.Wait()

			}
		}()
	}

	WG.Wait()

}
