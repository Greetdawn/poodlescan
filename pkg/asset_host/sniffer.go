package asset_host

import (
	"fmt"
	"poodle/pkg/common"
	"poodle/pkg/logger"
	"sync"

	"github.com/liushuochen/gotable"
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

	// 域名扫描器
	domainSniffer

	// ip扫描器
	ipSniffer

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

// 实现iSniffer的接口:StartSniff
func (this *Sniffer) StartIPSniff() {
	logger.LogInfo("start single ip sniff...", logger.LOG_TERMINAL_FILE)
	this.ipSniffer.super = this
	this.ipSniffer.StartSniff()
}

// StartDomainSniff: 域名嗅探器
func (this *Sniffer) StartDomainSniff() {
	logger.LogInfo("start single domain sniff...", logger.LOG_TERMINAL_FILE)
	this.domainSniffer.super = this
	this.domainSniffer.StartSniff()
}

// 打印所有的资产信息
func (this *Sniffer) PrintAssetHostList() {
	common.InitSqlite("123456789.db")

	for _, asset := range this.AlivedAssetHosts {
		var first bool = true
		tab, _ := gotable.Create("主机IP", "存活性", "开放端口", "服务信息")

		if len(asset.OpenedPorts) == 0 {
			tab.AddRow([]string{asset.RealIP, "存活", "", ""})
		} else {
			for key, value := range asset.OpenedPorts {
				if first {
					tab.AddRow([]string{asset.RealIP, "存活", key, value})
					first = false
				} else {
					tab.AddRow([]string{" ", " ", key, value})
				}
				common.AppendAsset2Sql(asset.RealIP, key, value, "tcp")
			}
		}
		fmt.Println(tab)
	}

	common.CloseDB()

	//******** 处理不存活情况  *****************
	// for _, asset := range this.DiedAssetHosts {
	// 	var first bool = true
	// 	tab, _ := gotable.Create("主机IP", "存活性", "开放端口", "服务信息")

	// 	if len(asset.OpenedPorts) == 0 {
	// 		tab.AddRow([]string{asset.RealIP, "不存活", "", ""})
	// 	} else {
	// 		for key, value := range asset.OpenedPorts {
	// 			if first {
	// 				tab.AddRow([]string{asset.RealIP, "不存活", key, value})
	// 				first = false
	// 			} else {
	// 				tab.AddRow([]string{" ", " ", key, value})
	// 			}
	// 		}
	// 	}
	// 	fmt.Println(tab)
	// }
	//******** 处理不存活情况  end *****************
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
func (this *Sniffer) SniffSubDomain(domain string) []string {
	return ScanSubDomain(domain)
}
