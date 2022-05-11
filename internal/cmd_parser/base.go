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
	UserInputTargetString string           // 用户输入的扫描目标
	IsReadTargetsFromFile bool             // 目标是否从文件中加载
	Threads               int              // 总线程数，同时扫描几个目标
	IpList                []string         // 存放解析后参数
	DomainList            []string         // 存放解析后参数
	TargetChan            chan TargetInput // 存放等待扫描的目标，是扫描器的输入
	isIP                  bool             //是否是IP
	SnifferPara                            // 存放嗅探器的初始化参数
}

// 端口扫描相关参数定义
type PortScanPara struct {
	Threads       int    // 端口扫描线程数，同时扫描多少端口
	Kind          string // 扫描类型，UDP扫描，TCP扫描，SYN扫描
	BreakPingScan bool   // 是否先ping确认主机是否存活
}

// 存放嗅探器的初始化参数
type SnifferPara struct {
	PortScanPara
}

// 目标输入，并发控制使用，作为通道传参
type TargetInput struct {
	Target string
	IsIP   bool
}

// 命令行参数解析，先使用默认参数初始化命令行
func CMDParseInit() *CMDPara {
	return &CMDPara{}
}
