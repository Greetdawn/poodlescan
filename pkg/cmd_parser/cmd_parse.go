package cmdparser

import (
	"flag"
	"fmt"
	"os"
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
		newTask := new(common.TASKUint)
		newTask.Params = make(map[string]string)
		newTask.Target = target.Target
		newTask.ControlCode = cc
		newTask.Params["ports"] = tp.PortList
		if target.IsIp {
			newTask.TargetType = common.TASKUint_TargetType_IP
		} else {
			newTask.TargetType = common.TASKUint_TargetType_Domain
		}
		common.G_TaskChannal <- newTask
	}
	close(common.G_TaskChannal)
	return nil
}

// 根据命令行输入初始化TerminalParams结构体
func ParsingUserTerminalLine() (terminalParams TerminalParams, err error) {
	// 获取终端参数的数量
	numberArgs := len(os.Args)

	// -t : 设置用户提供的目标
	flag.StringVar(&terminalParams.UserInputTargetString, "t", "", "设置扫描目标")

	// -Pn : 如果加了这个参数，则表示跳过Ping扫
	flag.BoolVar(&terminalParams.IsPn, "Pn", false, "跳过ping扫")

	// -T : 用户设置线程
	flag.IntVar(&terminalParams.ThreadsNumber, "T", 5, "设置并发，允许同时扫描几个目标")

	// -sn : 指定完整嗅探扫描
	flag.BoolVar(&terminalParams.Sniffer, "sn", false, "指定完整嗅探扫描")

	// -sn-p ：指定端口扫描
	flag.BoolVar(&terminalParams.PortScan, "sn-p", false, "指定端口扫描")

	// -p : 指定端口扫描端口范围
	flag.StringVar(&terminalParams.PortList, "p", "top1000", "指定扫描端口范围 例如:80,8080,80-1000 (default top1000)")

	// -sn-sub : 指定子域探测
	flag.BoolVar(&terminalParams.SubDomain, "sn-sub", false, "指定子域探测")

	// -sn-sp : 指定网站爬虫
	flag.BoolVar(&terminalParams.Spider, "sn-sp", false, "指定网站爬虫")

	// -sn-fp : 指定指纹识别
	flag.BoolVar(&terminalParams.Fingerprint, "sn-fp", false, "指定指纹识别")

	// -vs : 指定全漏洞扫描
	flag.BoolVar(&terminalParams.vulscan, "vs", false, "指定全漏洞扫描")

	// -vs-id : 指定具体漏洞编号扫描 例如：cve-2021-44228 (default "空值")

	// 开始解析
	flag.Parse()

	fmt.Printf("%d\n", numberArgs)
	fmt.Printf("terminalParams: %v\n", terminalParams)

	// debug信息
	logger.LogInfo("参数：用户设置目标："+terminalParams.UserInputTargetString, logger.LOG_TERMINAL)
	logger.LogInfo(fmt.Sprintf("参数: 是否跳过Ping扫描: %t", terminalParams.IsPn), logger.LOG_TERMINAL)
	logger.LogInfo("参数：线程数："+strconv.Itoa(terminalParams.ThreadsNumber), logger.LOG_TERMINAL)

	return
}
