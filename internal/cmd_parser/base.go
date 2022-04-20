// 本文件主要用来定义一些初始化参数
package cmdparser

// 为全局变量初始化
type PoodleInit interface {
	CMDUserInputParse(*CMDPara) // 再根据用户输入初始化
	GetTargets(*CMDPara)        // 解析并获取目标列表
}

// 命令行参数初始化的全局变量类型定义，用来作为输入
// 为后续的模块提供目标以及初始化
type CMDPara struct {
	UserInputTargetString string      // 用户输入的扫描目标
	IsReadTargetsFromFile bool        // 目标是否从文件中加载
	Threads               int         // 总线程数，同时扫描几个目标
	Target                chan string // 存放解析后参数，作为后续扫描的输入
	ExitFlag              chan bool   // 全局终止参数
	SnifferPara                       // 存放嗅探器的初始化参数
}

// 端口扫描相关参数定义
type PortScanPara struct {
	Threads    int    // 端口扫描线程数，同时扫描多少端口
	Kind       string // 扫描类型，UDP扫描，TCP扫描，SYN扫描
	IsPingScan bool   // 是否先ping确认主机是否存活
}

// 存放嗅探器的初始化参数
type SnifferPara struct {
	PortScanPara
}

// 命令行参数解析，先使用默认参数初始化命令行
func CMDParseInit() *CMDPara {
	return &CMDPara{}
}
