// @koi
// IP嗅探器
// 属于嗅探器的子类，主要用于嗅探和IP相关的资源
package sniffer

type ipSniffer struct {
	// 父类
	sniffer Sniffer
}

// 实现iSniffer的接口:StartSniff
func (this *ipSniffer) StartSniff() {

}

// 实现iSniffer的接口:SaveInfo
func (this *ipSniffer) SaveInfo() {

}

// 探测该IP 是否绑定了域名
func (this *ipSniffer) sniffBindDomain(ip string) (domain StDomain) {
	return domain
}

// IP端口嗅探
func (this *ipSniffer) sniffPort(ip string) (ports []int) {
	return ports
}
