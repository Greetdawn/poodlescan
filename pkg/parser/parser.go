package parser

import (
	"fmt"
	"poodle/pkg/asset_host"
	"strings"
)

const (
	// 类型码：
	// 01:	嗅探模块
	SNIFFER_MODULE int = 1
	// 01:100	单独IP嗅探
	SNIFFER_MODULE_SINGLE_IP int = 100
	// 01:200	单独IP嗅探
	SNIFFER_MODULE_SINGLE_DOMAIN int = 200
)

// 02:	漏洞扫描模块

// Parseing 解析命令行
// 解析器对外只暴露这个函数。
// 这个函数解析控制码。控制码分为两个部分，前面为模块码，后面为模块字码。
// 模块码确定接下来的控制权交给哪个模块执行，模块子码确定具体的执行流程。
func Parseing(controlCode int, param string) {
	params := strings.Split(param, " ")
	switch controlCode / 10000 {
	case SNIFFER_MODULE:
		switch controlCode % 10000 {
		case SNIFFER_MODULE_SINGLE_IP:
			// 检查参数数量，应该为1个参数
			if len(params) != 1 {
				fmt.Println("[E] The Number of params is not 1.")
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
				fmt.Println("[E] The Number of params is not 1.")
				return
			}

			sniffer := asset_host.Sniffer{}
			// 解析输入参数
			// 现在假设要嗅探 baidu.com
			sniffer.TargetDomains = append(sniffer.TargetDomains, asset_host.Domain{Name: params[0]})
			// 开始嗅探
			sniffer.StartDomainSniff()
			// 打印资产信息
			sniffer.AssetHosts[0].ToString()
		}
	}
}
