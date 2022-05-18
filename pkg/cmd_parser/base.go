package cmdparser

import (
	"fmt"
	"poodle/pkg/common"
)

// 全局变量，保存终端命令行参数结构体
var G_TerminalParam TerminalParams

// 命令行参数结构体
// 在用户输入后，通过flag模块，将用户输入的命令行转换成结构体保存
type TerminalParams struct {
	// 指定全流程扫描
	Full bool
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
	Vulscan bool
	// 指定具体漏洞名称或者编号扫描
	Vulscanid string
	// 指定具体漏洞类型扫描
	VulscanType string
	// 指定用户自定义脚本扫描
	VulscanDefined string
	// 指定目录探测功能
	VulscanDirsearch bool
	// 指定服务端口爆破
	VulscanBurst bool
	// 指定具体漏洞利用脚本
	Vulexploit string
	// 命令执行需要的参数
	Command string
}

// 根据终端参数结构体生成控制码
func (this *TerminalParams) GenerateControlCode() (controlCode uint) {
	controlCode = 0

	// -F 定义全流程扫描
	if this.Full {
		this.Sniffer = true
		this.Vulscan = true
	}

	// -sn 定义嗅探扫描全流程
	if this.Sniffer {
		this.PortScan = true
		this.SubDomain = true
		this.Spider = true
		this.Fingerprint = true
	}

	// -vs 定义全漏洞扫描流程
	if this.Vulscan {
		this.VulscanDirsearch = true
		this.VulscanBurst = true
	}

	// -Pn 	跳过主机存活检测
	// 默认不跳过
	if !this.IsPn {
		controlCode |= common.CC_PING_SCAN
	}

	// 指定端口扫描功能
	if this.PortScan {
		controlCode |= common.CC_PORT_SCAN
	}

	// 子域探测
	if this.SubDomain {
		controlCode |= common.CC_SUB_DOMAIN_SCAN
	}

	// 爬虫功能
	if this.Spider {
		controlCode |= common.CC_SPIDER
	}

	// 指纹识别
	if this.Fingerprint {
		controlCode |= common.CC_FINGERPRINT
	}

	// 全漏洞扫描
	if this.Vulscan {
		controlCode |= common.CC_VULSCAN
		controlCode = controlCode ^ common.CC_PING_SCAN
	}

	// 指定目录探测扫描
	if this.VulscanDirsearch {
		controlCode |= common.CC_VULSCAN_DIRSEARCH
		controlCode = controlCode ^ common.CC_PING_SCAN
	}

	// 指定服务端口爆破
	if this.VulscanBurst {
		controlCode |= common.CC_VULSCAN_BURST
		controlCode = controlCode ^ common.CC_PING_SCAN
	}

	fmt.Println(controlCode)
	return controlCode
}
