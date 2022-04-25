package utils

type StDomain struct {
	// 域名名称，不加www，如 baidu.com
	name string
}

// 域名端口探测实现的地方
func (this *StDomain) SniffPort() (ports []int) {
	return ports
}

// 域名端口探测实现的地方
func (this *StDomain) SniffRealIP() (realIP string) {
	return realIP
}
