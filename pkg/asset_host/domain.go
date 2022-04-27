package asset_host

type Domain struct {
	// 域名名称，不加www，如 baidu.com
	name string

	// 域名备案信息
	domainRecordInfo string
}

// 域名端口探测实现的地方
func (this *Domain) SniffPort() (ports []int) {
	return ports
}

// 域名端口探测实现的地方
func (this *Domain) SniffRealIP() (realIP string) {
	return realIP
}
