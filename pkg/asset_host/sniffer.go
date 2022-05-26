package asset_host

import (
	"poodle/pkg/common"
	"sync"
)

// 嗅探器单例
var pSniffer *Sniffer
var once sync.Once

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

	// 需要嗅探的IP列表
	TargetIPs []string

	// 需要嗅探的域名列表
	TargetDomains []common.Domain

	// 嗅探的资产结果
	AlivedAssetHosts []AssetHost

	// 不存活的资产结果
	DiedAssetHosts []AssetHost
}

// 用单例的方式获取嗅探器对象
func GetSnifferObj() *Sniffer {
	once.Do(func() {
		pSniffer = &Sniffer{}
	})
	return pSniffer
}

// 追加存活资产信息
func (this *Sniffer) AppendAlivedAssetHost(asset AssetHost) {
	this.AlivedAssetHosts = append(this.AlivedAssetHosts, asset)
}

// 追加不存活资产信息
func (this *Sniffer) AppendDiedAssetHost(asset AssetHost) {
	this.DiedAssetHosts = append(this.DiedAssetHosts, asset)
}

// 嗅探目标主机是否存活
func (this *Sniffer) IsHostAlived(target string) bool {
	return common.IsHostAlived(target)
}

// 嗅探目标主机开放端口信息
func (this *Sniffer) SnifferHostOpenedPorts(target string) sync.Map {
	return ScanHostOpenedPorts(target)
}

// 嗅探域名的子域信息
func (this *Sniffer) SniffSubDomain(domain string) (domains []common.Domain) {
	for _, v := range ScanSubDomain(domain) {
		domains = append(domains, common.Domain{Name: v})
	}
	return
}
