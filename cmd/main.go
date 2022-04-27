package main

import (
	"fmt"
	"poodle/pkg/asset_host"
)

func main() {
	sniffer := asset_host.Sniffer{}

	// 解析输入参数
	// 现在假设要嗅探 baidu.com
	sniffer.CmdCode = 2

	sniffer.TargetDomains = append(sniffer.TargetDomains, asset_host.Domain{"11111", "22222"})

	fmt.Println(len(sniffer.TargetDomains))
	// 开始嗅探
	sniffer.StartSniff()
	// 打印资产信息
	sniffer.AssetHosts[0].ToString()
}
