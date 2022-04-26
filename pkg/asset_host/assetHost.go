package asset_host

import "poodle/pkg/sniffer"

type StAssetHost struct {
	// 本资产的IP地址
	// 多个资产IP地址相同时，合并成一个资产
	ip string
	// 资产域名
	// 当前为一个域名，后面根据需要可能会换成域名切片
	domain sniffer.StDomain
	// 资产开放的端口信息
	openedPorts []int
}
