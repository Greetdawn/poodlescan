package asset_host

type Domain struct {
	// 域名名称，不加www，如 baidu.com
	Name string

	// 域名备案信息
	DomainRecordInfo string
}

// 域名端口探测实现的地方
func (this *Domain) SniffPort() []int {
	return []int{1, 2}
}

// 通过域名嗅探域名的真实IP
func (this *Domain) SniffRealIP() (realIP string) {
	return realIP
}

// 获取域名的备案信息
func (this *Domain) SniffDomainRecordInfo() (info string) {
	info = "备案信息"
	return info
}

// 嗅探域名的子域信息
func (this *Domain) SniffSubDomain() []Domain {
	ds := make([]Domain, 1)
	ds[0].Name = "baidu.sub.com"
	ds[0].DomainRecordInfo = "京 2022"
	return ds
}
