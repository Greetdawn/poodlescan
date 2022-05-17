package cmdparser

import "poodle/pkg/common"

// 全局变量，保存终端命令行参数结构体
var G_TerminalParam TerminalParams

// 命令行参数结构体
// 在用户输入后，通过flag模块，将用户输入的命令行转换成结构体保存
type TerminalParams struct {
	// 用户输入的扫描目标，原始字符串
	UserInputTargetString string
	// 标志;一些特殊的标志
	// 000000000 默认情况，保留
	// 000000001 目标从文件中读取
	Flag byte
	// 用户设置的线程数
	ThreadsNumber int // 总线程数，同时扫描几个目标
	// Pn "跳过Ping扫"，默认不跳过Ping扫,false
	IsPn bool
	// 定义是否进行完整嗅探扫描
	Sniffer bool
	// 定义是否进行端口扫描
	PortScan bool
	// 定义端口扫描的范围
	PortList string
	// 定义是否进行子域探测
	SubDomain bool
	// 指定是否进行网站爬虫
	Spider bool
	// 指定是否进行指纹识别
	Fingerprint bool
	// 指定全漏洞扫描
	vulscan bool
}

// 获取一个TerminalParams对象
func GetTerminalParamObj() *TerminalParams {
	return &TerminalParams{}
}

// 根据终端参数结构体生成控制码
func (this *TerminalParams) GenerateControlCode() (controlCode uint) {
	controlCode = 0

	// -Pn 	跳过主机存活检测
	// 默认不跳过
	if !this.IsPn {
		controlCode |= common.CC_PING_SCAN
	}

	// 端口扫描
	controlCode |= common.CC_PORT_SCAN
	return controlCode
}
