package common

type Domain struct {
	// 域名名称，不加www，如 baidu.com
	Name string

	// 域名备案信息
	IPC string

	// 本域名是否存活
	IsAlived bool
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
func (this *Domain) SniffIPC() (info string) {
	info = "备案信息"
	return info
}

// 嗅探域名的子域信息
// 在子域探测的同时需要进行主机存活的判断，如果主机存活加到返回值中，否则丢弃
func (this *Domain) SniffSubDomain() []Domain {
	ds := make([]Domain, 1)
	ds[0].Name = "baidu.sub.com"
	ds[0].IPC = "京 2022"
	return ds
}

// 嗅探域名主机是否存活
func (this *Domain) SniffDomainAlive() bool {
	return IsAlivedOfHostByDomain(*this)
}
