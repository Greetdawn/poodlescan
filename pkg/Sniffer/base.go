package sniffer

import (
	"sync"
)

// 只接受IP地址和域名
// 判断输入类型，是ip[false]还是域名[true]；
// 判断是否输入域名，或者ip上是否绑定域名，查询备案号；
// 如果输入是域名则根据域名探测子域，以及根据域名探测端口；
// 如果输入是ip则探测端口，以及ip上绑定的域名，
// 此时因为输入的是IP，所以探测到ip所绑定的域名后不再进行子域探测，
// 以防止还有子域绑定到其他的ip上所导致的扫描偏离主体。
type TeddySniffer interface {
	// 先格式化，然后判断输入类型，是ip[false]还是域名[true]
	TargetType() bool

	// 如果有域名，则保存域名，并且HasDomain==true
	HasDomain() bool

	// 判断HasDomain是否为true备案号，需要稳定爬虫
	GetIPC(bool, string) string

	// 如果是域名<IsIP[false]>则根据域名获取子域
	GetSubdomain(bool, string) []string

	// 根据ip或者域名获取端口
	GetPort(string) []string

	// 爬虫爬取的特殊后缀,且响应为200
	// map[<suffix>][]String
	// Example:map["zip"][]String{....}
	GetSpecialSuffix(string) sync.Map

	// 根据域名或者IP目录爆破
	GetWeekAddress(string) []string

	// 保存结果为json格式
	JsonSave()
}

// 实例化接口
type TeddySniff struct {
	SrcTarget     string   // 接受待检测的url或者IP
	IsIP          bool     // 判断收到的参数是IP还是域名
	SubDomain     []string // 子域列表
	Domain        string   // 域名
	IPAddress     string   // IP地址
	IPC           string   // 备案号
	Ports         []string // 开放端口\端口信息
	WeekAddress   []string // 后台地址等等
	SpecialSuffic sync.Map // 特殊后缀地址文件比如.zip.xlsx后缀
}

// 结构体构造函数
func TeddyInit() TeddySniff {
	var Teddy TeddySniff
	return Teddy
}

// func TeddyStartSniffer(c CMDPara, T TeddySniffer) TeddySniff {
// 	var ts TeddySniff
// 	ts.IsIP = !T.HasDomain()
// 	return ts
// }
