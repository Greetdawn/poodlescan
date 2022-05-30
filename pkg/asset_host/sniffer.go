package asset_host

import (
	"poodle/pkg/common"
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

func Assets2Strings(isRemove bool) (asstesString []string) {
	for _, v := range GetSnifferObj().AlivedAssetHosts {
		var targ string
		var targAlived string
		if v.IsIP {
			targ = v.RealIP
		} else {
			targ = v.Domain.Name
		}

		if v.IsAlived {
			targAlived = "存活"
		} else {
			targAlived = "不存活"
		}

		var subDomain [][]string
		for _, v := range v.SubDomains {
			if v.IsAlived {
				subDomain = append(subDomain, []string{v.Name, "存活"})
			} else {
				subDomain = append(subDomain, []string{v.Name, "不存活"})
			}
		}

		tab, _ := gotable.Create("目标", "子域名", "存活状态", "开放端口", "服务信息", "爬虫结果")
		tab.AddRow([]string{targ, "", targAlived, "", "", ""})

		for _, v := range subDomain {
			tab.AddRow([]string{"", v[0], v[1], "", "", ""})
		}
		for k, v := range v.OpenedPorts {
			tab.AddRow([]string{"", "", "", k, v, ""})
		}
		asstesString = append(asstesString, tab.String())
	}
	if isRemove {
		GetSnifferObj().AlivedAssetHosts = nil
	}
	return
}
