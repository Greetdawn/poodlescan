package parser

import (
	"poodle/pkg/asset_host"
	"poodle/pkg/common"
	"poodle/pkg/logger"
)

const (
	// 类型码：
	// 01:	嗅探模块
	SNIFFER_MODULE uint32 = 1
	// 01:100	单独IP嗅探
	SNIFFER_MODULE_SINGLE_IP uint32 = 10000
	// 01:200	单独域名嗅探
	SNIFFER_MODULE_SINGLE_DOMAIN uint32 = 20000
)

// 02:	漏洞扫描模块

// Parseing 解析命令行
// 解析器对外只暴露这个函数。
// 这个函数解析控制码。控制码分为两个部分，前面为模块码，后面为模块字码。
// 模块码确定接下来的控制权交给哪个模块执行，模块子码确定具体的执行流程。
// uint32：表示 32 位无符号整型 大小：32 位 范围：0～4294967295
// 	 4294967295
// / 0010000000  得到大模块类，最大429个类别
// % 0010000000  得到子控制码。
// 控制字码9999999共7位，前3位表示子码的大类，后四位小类别（如果有）。
// 10010000
// 10000000
// 4101238765
func Parseing(controlCode uint32, params []string) {
	switch controlCode / 10000000 {
	case SNIFFER_MODULE:
		switch controlCode % 10000000 {
		case SNIFFER_MODULE_SINGLE_IP:
			// 检查参数数量，应该为1个参数
			if len(params) != 1 {
				logger.OutputError("Parseing", "The Number of params is not 1.")
				return
			}
			// 声明一个嗅探器
			sniffer := asset_host.Sniffer{}
			// 填入需要嗅探的目标
			sniffer.TargetIPs = append(sniffer.TargetIPs, params[0])
			// 嗅探器开始嗅探
			sniffer.StartIPSniff()
			// 打印嗅探出的所有资产信息
			sniffer.PrintAssetHostList()
		case SNIFFER_MODULE_SINGLE_DOMAIN:
			// 检查参数数量，应该为1个参数
			if len(params) != 1 {
				logger.OutputError("Parseing", "The Number of params is not 1.")
				return
			}

			sniffer := asset_host.Sniffer{}
			// 填入需要嗅探的目标
			sniffer.TargetDomains = append(sniffer.TargetDomains, common.Domain{Name: params[0]})
			// 开始嗅探
			sniffer.StartDomainSniff()
			// 打印资产信息
			sniffer.PrintAssetHostList()
		}
	}
}
