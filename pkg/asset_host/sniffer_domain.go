package asset_host

import "fmt"

// @koi
// 域名嗅探器
// 属于嗅探器的子类，主要用于嗅探和域名相关的资源
type domainSniffer struct {
	// 父类
	super *Sniffer

	// 所有域名列表（本域名+子域名）
	domainList []Domain
}

// 实现iSniffer的接口:StartSniff
func (this *domainSniffer) StartSniff() {
	fmt.Println("开始 域名嗅探...")

	fmt.Println(len(this.super.TargetDomains))

	for i := 0; i < len(this.super.TargetDomains); i++ {
		// 创建一个资产
		var curAssetHost = new(AssetHost)

		// 复制域名信息
		curAssetHost.Domain = this.super.TargetDomains[i]

		// 嗅探域名真实IP
		curAssetHost.RealIP = this.sniffRealIP(&this.super.TargetDomains[i])

		// 通过域名探测开放的端口号
		curAssetHost.Ports = append(curAssetHost.Ports, this.sniffPort(&this.super.TargetDomains[i])...)

		// 嗅探域名备案信息
		curAssetHost.IPC = this.super.TargetDomains[i].SniffDomainRecordInfo()

		// 嗅探域名子域信息
		curAssetHost.SubDomains = this.super.TargetDomains[i].SniffSubDomain()

		// 将当前资产保存到父类的资产列表中
		this.super.AssetHosts = append(this.super.AssetHosts, *curAssetHost)
	}

}

// 实现iSniffer的接口:SaveInfo
func (this *domainSniffer) SaveInfo() {

}

// 扫描子域信息
// 返回值类型为 域名切片
func (this *domainSniffer) sniffSubDomain() (domains []Domain) {
	var curSniffDomain = this.super.domainList[0].Name
	fmt.Printf("当前需要嗅探的域名为：%s\r\n", curSniffDomain)
	domains = make([]Domain, 1)
	domains[1] = Domain{"baidu.com", "备案"}
	return domains
}

// 通过域名嗅探端口
func (this *domainSniffer) sniffPort(domain *Domain) []int {
	return domain.SniffPort()
}

// 通过域名获取探测真实IP
func (this *domainSniffer) sniffRealIP(domain *Domain) string {
	var curSniffDomain = domain.Name
	fmt.Printf("当前需要嗅探的域名为：%s\r\n", curSniffDomain)
	return "162.27.12.34"
}
