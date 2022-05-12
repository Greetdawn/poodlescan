// @koi
// IP嗅探器
// 属于嗅探器的子类，主要用于嗅探和IP相关的资源
package asset_host

import (
	"poodle/pkg/common"
	"poodle/pkg/logger"
)

type ipSniffer struct {
	// 父类
	super *Sniffer
}

// StartSniff 实现iSniffer的接口:StartSniff
// 嗅探IP的信息
// 这里是流程控制代码，一般情况下不需要更改
func (this *ipSniffer) StartSniff() {
	for i := 0; i < len(this.super.TargetIPs); i++ {
		currentIP := this.super.TargetIPs[i]
		// 1. 通过IP嗅探主机存活情况
		if !common.IsHostAlived(currentIP) {
			// 主机不存活
			logger.LogInfo(currentIP+" 不存活", logger.LOG_TERMINAL_FILE)
			// 添加到不存活资产主机列表中
			var asset AssetHost
			asset.IsAlived = false
			this.super.AppendDiedAssetHost(asset)
		} else {
			// 主机存活情况
			// 2. 嗅探端口信息
			ports := this.sniffPort(currentIP)
			// 3. 资产信息
			var asset AssetHost
			// 是IP主机
			asset.IsIP = true
			// 设置IP
			asset.RealIP = currentIP
			// 设置IP存活
			asset.IsAlived = true
			// 设置开放端口信息
			asset.Ports = ports
			// 添加到资产主机列表中
			this.super.AppendAlivedAssetHost(asset)
		}
	}
}

// 实现iSniffer的接口:SaveInfo
func (this *ipSniffer) SaveInfo() {

}

// IP端口嗅探的具体实现代码。
func (this *ipSniffer) sniffPort(ip string) (ports []int) {
	ports = make([]int, 10)
	ports[0] = 80
	ports[1] = 81
	ports[2] = 8080
	return ports
}
