package cmdparser

import "flag"

// todo 根据命令行输入初始化全局变量
func (CMDPara) CMDUserInputParse(CMDParas *CMDPara) {
	CMDParas.UserInputTargetString = *flag.String("t", "", "扫描目标")
	CMDParas.IsPingScan = *flag.Bool("ping", true, "是否先ping确认目标存活")
	// todo Paras其他参数
	flag.Parse()

}
