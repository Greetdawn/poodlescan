package cmdinit

// 命令行的全局变量
type CMDPara struct {
	Targets               []string // 目标
	IsReadTargetsFromFile bool     // 目标是否从文件中加载
	Threads               int      // 总线程数，同时扫描几个目标
	PortScanPara                   // 端口扫描相关参数定义
}

// 端口扫描相关参数定义
type PortScanPara struct {
	Threads    int    // 端口扫描线程数，同时扫描多少端口
	Kind       string // 扫描类型，UDP扫描，TCP扫描，SYN扫描
	IsPingScan bool   // 是否先ping确认主机是否存活
}

// 全局初始化
type PoodleInit interface {
	CMDPause()           // 先使用默认参数初始化
	GetTargets(*CMDPara) // 解析并获取目标列表
}

// 命令行参数解析，先使用默认参数初始化命令行
func (CMDPara) CMDPause() *CMDPara {
	CMDInit := &CMDPara{
		Threads: 1, // 默认同时只对一个目标进行扫描
		PortScanPara: PortScanPara{
			IsPingScan: true,  // 默认先ping确认主机是否存活
			Threads:    200,   // 默认端口的扫描线程200
			Kind:       "TCP", // 默认TCP
		},
	}
	return CMDInit
}
