package asset_host

import (
	"fmt"
	"poodle/pkg/common"
	"poodle/pkg/logger"
)

// sniffer的接口类
type ISniffer interface {
	// 开始执行嗅探的工作。
	StartSniff()

	// 保存嗅探出来的结果。目前版本只需要把结果打印出来即可。
	//	SaveInfo()

	// 将对象转化成字符串输出
	//	toString()
}

// 嗅探器的父类，对外只暴露这个类。
type Sniffer struct {
	// 命令码
	// CmdCode int

	// 需要嗅探的IP列表
	TargetIPs []string

	// 需要嗅探的域名列表
	TargetDomains []common.Domain

	// 域名扫描器
	domainSniffer

	// ip扫描器
	ipSniffer

	// 嗅探的资产结果
	AlivedAssetHosts []AssetHost

	// 不存活的资产结果
	DiedAssetHosts []AssetHost
}

// 追加存活资产信息
func (this *Sniffer) AppendAlivedAssetHost(asset AssetHost) {
	logger.OutputInfo("sniffer", "append an asset host to the alived list of asset host.")
	this.AlivedAssetHosts = append(this.AlivedAssetHosts, asset)
	logger.OutputInfo("sniffer", fmt.Sprintf("the current number of asset host is %d.", len(this.AlivedAssetHosts)))
}

// 追加不存活资产信息
func (this *Sniffer) AppendDiedAssetHost(asset AssetHost) {
	logger.OutputInfo("sniffer", "append an asset host to the died list of asset host.")
	this.DiedAssetHosts = append(this.DiedAssetHosts, asset)
	logger.OutputInfo("sniffer", fmt.Sprintf("the current number of asset host is %d.", len(this.DiedAssetHosts)))
}

// 实现iSniffer的接口:StartSniff
func (this *Sniffer) StartIPSniff() {
	logger.OutputInfo("sniffer", "start single ip sniff...")
	this.ipSniffer.super = this
	this.ipSniffer.StartSniff()
}

// StartDomainSniff: 域名嗅探器
func (this *Sniffer) StartDomainSniff() {
	logger.OutputInfo("sniffer", "start single domain sniff...")
	this.domainSniffer.super = this
	this.domainSniffer.StartSniff()
}

// 打印所有的资产信息
func (this *Sniffer) PrintAssetHostList() {
	logger.OutputInfo("sniffer", "资产主机信息：")
	for _, asset := range this.AlivedAssetHosts {
		logger.OutputNoFormat(asset.ToString())
	}
	logger.OutputInfo("sniffer", "不存活资产主机信息：")
	for _, asset := range this.DiedAssetHosts {
		logger.OutputNoFormat(asset.ToString())
	}
}
