package cmdparser

import (
	"flag"
	"fmt"
	"poodle/pkg/common"
	controllor "poodle/pkg/controller"
	"poodle/pkg/logger"
	"sort"
	"strconv"
)

type _Target struct {
	Target string //目标
	IsIp   bool   //目标是否是IP
}

// 模块变量初始化
func init() {
	sort.Strings(common.DOMAINARRAY)
}

// 命令行解析，生成任务，将任务写入任务通道中
func GenrateTasks(tp *TerminalParams, cc uint) error {
	// 3. 解析用户的输入目标
	targets, err := controllor.ParseUserInputTargetString(G_TerminalParam.UserInputTargetString)
	if err != nil {
		// 解析用户的输入目标发生错误。
		return err
	}

	// 4. 将目标转化成 Task ，写入通道中。
	// 新建任务通道
	common.G_TaskChannal = make(chan *common.TASKUint, 10000)
	for _, target := range targets {
		if target.IsIp {
			common.G_TaskChannal <- &common.TASKUint{Target: target.Target, TargetType: common.TASKUint_TargetType_IP, ControlCode: cc}
		} else {
			common.G_TaskChannal <- &common.TASKUint{Target: target.Target, TargetType: common.TASKUint_TargetType_Domain, ControlCode: cc}
		}
	}
	close(common.G_TaskChannal)
	return nil
}

// 根据命令行输入初始化TerminalParams结构体
func ParsingUserTerminalLine() (terminalParams TerminalParams, err error) {
	// -t : 设置用户提供的目标
	flag.StringVar(&terminalParams.UserInputTargetString, "t", "", "设置扫描目标")

	// -Pn : 如果加了这个参数，则表示跳过Ping扫
	flag.BoolVar(&terminalParams.IsPn, "Pn", false, "true:跳过ping扫;false（默认）:不跳过Ping扫")

	// -T : 用户设置线程
	flag.IntVar(&terminalParams.ThreadsNumber, "T", 5, "设置并发，允许同时扫描几个目标")
	// todo Paras其他参数

	// 开始解析
	flag.Parse()

	// debug信息
	logger.LogInfo("参数：用户设置目标："+terminalParams.UserInputTargetString, logger.LOG_TERMINAL)
	logger.LogInfo(fmt.Sprintf("参数：是否跳过Ping扫描：%t", terminalParams.IsPn), logger.LOG_TERMINAL)
	logger.LogInfo("参数：线程数："+strconv.Itoa(terminalParams.ThreadsNumber), logger.LOG_TERMINAL)

	return
}
