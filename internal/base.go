package main

// 全局变量定义
type CMDPara struct {
	Target  string // 目标
	Threads int    // 总线程数，同时扫描几个目标
	PortScanPara
}

type PortScanPara struct {
	Threads    int    // 端口扫描线程数，同时扫描多少端口
	Kind       string // 扫描类型，UDP扫描，TCP扫描，SYN扫描
	IsPingScan bool   // 是否先ping确认主机是否存活
}

var CMDInit = CMDPara{
	Threads: 1, // 默认同时只对一个目标进行扫描
	PortScanPara: PortScanPara{
		IsPingScan: true,  // 默认先ping确认主机是否存活
		Threads:    200,   // 默认端口的扫描线程200
		Kind:       "TCP", // 默认TCP
	},
}
