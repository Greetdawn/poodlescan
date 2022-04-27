package asset_host

// sniffer的接口类
type ISniffer interface {
	// 开始执行嗅探的工作。
	StartSniff()

	// 保存嗅探出来的结果。目前版本只需要把结果打印出来即可。
	//	SaveInfo()

	// 将对象转化成字符串输出
	//	toString()
}

// 嗅探器的父类，对外只暴露这个类。
type Sniffer struct {
	// 命令码
	CmdCode int

	// 需要嗅探的IP列表
	TargetIPs []string

	// 需要嗅探的域名列表
	TargetDomains []Domain

	// 域名扫描器
	domainSniffer

	// ip扫描器
	ipSniffer

	// 嗅探的资产结果
	AssetHosts []AssetHost
}

// 实现iSniffer的接口:StartSniff
func (this *Sniffer) StartSniff() {
	// ip sniffer
	if this.CmdCode == 1 {
		this.ipSniffer.StartSniff()
	} else {
		this.domainSniffer.super = this
		this.domainSniffer.StartSniff()
	}
}
