// @koi
// 域名嗅探器
// 属于嗅探器的子类，主要用于嗅探和域名相关的资源
package sniffer

type domainSniffer struct {
	// 父类
	sniffer Sniffer

	// 所有域名列表（本域名+子域名）
	domainList []StDomain
}

// 实现iSniffer的接口:StartSniff
func (this *domainSniffer) StartSniff() {

}

// 实现iSniffer的接口:SaveInfo
func (this *domainSniffer) SaveInfo() {

}

// 获取域名的备案信息
func (this *domainSniffer) getDomainRecordInfo(domain *StDomain) (info string) {
	return info
}

// 扫描子域信息
// 返回值类型为 域名切片
func (this *domainSniffer) sniffSubDomain(domain *StDomain) (domains []StDomain) {
	return domains
}

// 域名端口嗅探
func (this *domainSniffer) sniffPort(domain *StDomain) []int {
	return domain.SniffPort()
}

// 通过域名获取探测真实IP
func (this *domainSniffer) sniffRealIP(domain *StDomain) string {
	return domain.SniffRealIP()
}
