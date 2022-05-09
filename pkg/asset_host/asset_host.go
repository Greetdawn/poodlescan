package asset_host

import (
	"fmt"
	"poodle/pkg/common"
	"sync"
)

type AssetHost struct {
	// 源
	// IP或者域名
	// 主要描述输入信息
	SrcTarget string

	// 是否是IP
	// 资产嗅探主要分两个分支，这里表明这个资产是通过IP或者域名嗅探
	IsIP bool

	// 本资产的IP地址
	// 多个资产IP地址相同时，合并成一个资产
	RealIP string

	// 资产域名
	// IsIP == true 	:RealIP嗅探出的对应域名
	// IsIP == false 	: 为要嗅探的域名
	Domain common.Domain

	// 存活情况
	// IsIP == true 	: RealIP存活情况
	// IsIP == false 	: 无效
	IsAlived bool

	// 子域列表
	SubDomains []common.Domain

	// 开放的端口
	// <int, string> <开放端口号, 对应端口信息>
	Ports []int

	// 域名备案信息
	// IPC string

	// web目标的结果
	AssetWeb
}

// 此处存放的是web目标的结果
// 将资产主机中关于Web内容单独出来
type AssetWeb struct {
	// 标签。例如：cms、xx框架等
	Tag []string

	// 存活Web服务的响应头
	Header sync.Map

	// 泄露的敏感路径等等
	// 例如后台管理路径、某些敏感文件路径等
	WeekAddress []string

	// 特殊后缀地址文件比如.zip.xlsx后缀
	SpecialSuffic sync.Map
}

func (this *AssetHost) ToString() (info string) {
	info = fmt.Sprintf("\tSrcTarget : %s\n\tIsIP :%t\n\tRealIP : %s\n\tDomain : %#v\n\tIP存活情况：%t\n\t子域列表 : %#v\n\tPorts : %#v\n\n",
		this.SrcTarget, this.IsIP, this.RealIP, this.Domain, this.IsAlived, this.SubDomains, this.Ports)
	return info
}
