package asset_host

import (
	"fmt"
	"poodle/pkg/common"
	"poodle/pkg/logger"
)

// @koi
// 域名嗅探器
// 属于嗅探器的子类，主要用于嗅探和域名相关的资源
type domainSniffer struct {
	// 父类
	super *Sniffer

	// 所有域名列表（本域名+子域名）
	domainList []common.Domain
}

// 实现iSniffer的接口:StartSniff
func (this *domainSniffer) StartSniff() {
	logger.OutputInfo("domain sniffer", "start sniff domain...")
	for i := 0; i < len(this.super.TargetDomains); i++ {
		// 创建一个资产
		var alivedAsset AssetHost
		// 1. 设置不是IP
		alivedAsset.IsIP = false
		// 2. 复制域名信息。如果不进行真实IP的探测，这里将是资产的主键
		alivedAsset.Domain = this.super.TargetDomains[i]
		// 3. 嗅探主域备案信息

		// 4. 嗅探域名子域信息
		subDomains := this.super.TargetDomains[i].SniffSubDomain()
		diedAsset := alivedAsset
		for _, v := range subDomains {
			if v.IsAlived {
				alivedAsset.SubDomains = append(alivedAsset.SubDomains, v)
			} else {
				diedAsset.SubDomains = append(diedAsset.SubDomains, v)
			}
		}

		// 保存存活资产列表
		this.super.AppendAlivedAssetHost(alivedAsset)
		// 保存不存活资产列表
		if len(diedAsset.SubDomains) > 0 {
			this.super.AppendDiedAssetHost(diedAsset)
		}
	}
}

// 实现iSniffer的接口:SaveInfo
func (this *domainSniffer) SaveInfo() {

}

// 扫描子域信息
// 需要同步探测每个子域是否存活
func (this *domainSniffer) sniffSubDomain() (domains []common.Domain) {
	domains = make([]common.Domain, 1)
	domains[1] = common.Domain{Name: "baidu.com", IPC: "备案"}
	return domains
}

// 通过域名嗅探端口
func (this *domainSniffer) sniffPort(domain *common.Domain) []int {
	return domain.SniffPort()
}

// 获取域名的备案信息
// 将获取到的信息保存到传入的域名结构体重
// 返回值为嗅探到的备案信息。
func SniffDomainRecordInfo(domain *common.Domain) (info string) {
	// todo 通过域名嗅探备案信息
	info = "备案信息"
	// 写入Domain结构体中
	domain.IPC = info
	// 返回嗅探的备案信息，一般情况下不用接收
	return info
}

// 通过域名获取探测真实IP
func (this *domainSniffer) sniffRealIP(domain *common.Domain) string {
	var curSniffDomain = domain.Name
	fmt.Printf("当前需要嗅探的域名为：%s\r\n", curSniffDomain)
	return "162.27.12.34"
}
