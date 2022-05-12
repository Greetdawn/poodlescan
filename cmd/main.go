package main

import (
	cmdparser "poodle/internal/cmd_parser"
	"poodle/internal/logic"
	"poodle/pkg/common"
)

var (
	// 初始化参数
	CmdParas = cmdparser.CMDParseInit()
)

func main() {

	// 解析命令行参数
	CmdParas.CMDUserInputParse()

	// 开扫
	/*
		顺序扫描：
		logic.PoodleLogic(CmdParas,true, <第一步函数名>,<第二步函数名>...)
		并发扫描:
		logic.PoodleLogic(CmdParas,false, 函数名,函数名...)
	*/

	logic.PoodleLogic(CmdParas, false, common.ScanHostAlived)

}
