package common

// 通过IP来判断目标主机是否存活
// 可以考虑用接口让isAlivedOfHostByDomain这两个函数合并成一个函数
func IsAlivedOfHostByIP(ip string) bool {
	return true
}

// 通过域名来判断目标主机是否存活
func IsAlivedOfHostByDomain(domain Domain) bool {
	return true
}
