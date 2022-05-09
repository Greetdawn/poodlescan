package asset_host

import "fmt"

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
	TargetDomains []Domain

	// 域名扫描器
	domainSniffer

	// ip扫描器
	ipSniffer

	// 嗅探的资产结果
	AssetHosts []AssetHost
}

// 追加资产信息
func (this *Sniffer) AppendAssetHost(asset AssetHost) {
	fmt.Println("[I] append an asset host to the list of asset host.")
	this.AssetHosts = append(this.AssetHosts, asset)
	fmt.Printf("[I] the current number of asset host is %d.\n", len(this.AssetHosts))
}

// 实现iSniffer的接口:StartSniff
func (this *Sniffer) StartIPSniff() {
	fmt.Println("[I] start single ip sniff...")
	this.ipSniffer.super = this
	this.ipSniffer.StartSniff()
}

// StartDomainSniff: 域名嗅探器
func (this *Sniffer) StartDomainSniff() {
	this.domainSniffer.super = this
	this.domainSniffer.StartSniff()
}

// 打印所有的资产信息
func (this *Sniffer) PrintAssetHostList() {
	fmt.Println("\n[I] 所有的资产主机信息：")
	for _, asset := range this.AssetHosts {
		fmt.Println()
		asset.ToString()
	}
}
