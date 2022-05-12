// 本文件主要用来定义一些初始化参数
package cmdparser

// 为全局变量初始化
type PoodleInit interface {
	CMDUserInputParse(*CMDPara) // 再根据用户输入初始化
	GetTargets(*CMDPara)        // 解析并获取目标列表
}

// 命令行参数初始化的全局变量类型定义，用来作为输入
// 为后续的模块提供目标以及初始化
type CMDPara struct {
	UserInputTargetString string           // 用户输入的扫描目标
	IsReadTargetsFromFile bool             // 目标是否从文件中加载
	Threads               int              // 总线程数，同时扫描几个目标
	IpList                []string         // 存放解析后参数
	DomainList            []string         // 存放解析后参数
	TargetChan            chan TargetInput // 存放等待扫描的目标，是扫描器的输入
	isIP                  bool             //是否是IP
	SnifferPara                            // 存放嗅探器的初始化参数
}

// 端口扫描相关参数定义
type PortScanPara struct {
	Threads       int    // 端口扫描线程数，同时扫描多少端口
	Kind          string // 扫描类型，UDP扫描，TCP扫描，SYN扫描
	BreakPingScan bool   // 是否先ping确认主机是否存活
}

// 存放嗅探器的初始化参数
type SnifferPara struct {
	PortScanPara
}

// 目标输入，并发控制使用，作为通道传参
type TargetInput struct {
	Target string
	IsIP   bool
}

// 命令行参数解析，先使用默认参数初始化命令行
func CMDParseInit() *CMDPara {
	return &CMDPara{}
}

// 定义常见的域名后缀
var DomainArray []string = []string{
	".com", ".edu", ".gov", ".int", ".mil", ".net", ".org", ".biz", ".info", ".pro", ".name", ".museum", ".coop", ".aero", ".xxx", ".idv", ".ac", ".ad", ".ae", ".af", ".ag", ".ai", ".al", ".am", ".an", ".ao", ".aq", ".ar", ".as", ".at", ".au", ".aw", ".az", ".ba", ".bb", ".bd", ".be", ".bf", ".bg", ".bh", ".bi", ".bj", ".bm", ".bn", ".bo", ".br", ".bs", ".bt", ".bv", ".bw", ".by", ".bz", ".ca", ".cc", ".cd", ".cf", ".cg", ".ch", ".ci", ".ck", ".cl", ".cm", ".cn", ".co", ".cr", ".cu", ".cv", ".cx", ".cy", ".cz", ".de", ".dj", ".dk", ".dm", ".do", ".dz", ".ec", ".ee", ".eg", ".eh", ".er", ".es", ".et", ".eu", ".fi", ".fj", ".fk", ".fm", ".fo", ".fr", ".ga", ".gd", ".ge", ".gf", ".gg", ".gh", ".gi", ".gl", ".gm", ".gn", ".gp", ".gq", ".gr", ".gs", ".gt", ".gu", ".gw", ".gy", ".hk", ".hm", ".hn", ".hr", ".ht", ".hu", ".id", ".ie", ".il", ".im", ".in", ".io", ".iq", ".ir", ".is", ".it", ".je", ".jm", ".jo", ".jp", ".ke", ".kg", ".kh", ".ki", ".km", ".kn", ".kp", ".kr", ".kw", ".ky", ".kz", ".la", ".lb", ".lc", ".li", ".lk", ".lr", ".ls", ".lt", ".lu", ".lv", ".ly", ".ma", ".mc", ".md", ".mg", ".mh", ".mk", ".ml", ".mm", ".mn", ".mo", ".mp", ".mq", ".mr", ".ms", ".mt", ".mu", ".mv", ".mw", ".mx", ".my", ".mz", ".na", ".nc", ".ne", ".nf", ".ng", ".ni", ".nl", ".no", ".np", ".nr", ".nu", ".nz", ".om", ".pa", ".pe", ".pf", ".pg", ".ph", ".pk", ".pl", ".pm", ".pn", ".pr", ".ps", ".pt", ".pw", ".py", ".qa", ".re", ".ro", ".ru", ".rw", ".sa", ".sb", ".sc", ".sd", ".se", ".sg", ".sh", ".si", ".sj", ".sk", ".sl", ".sm", ".sn", ".so", ".sr", ".st", ".sv", ".sy", ".sz", ".tc", ".td", ".tf", ".tg", ".th", ".tj", ".tk", ".tl", ".tm", ".tn", ".to", ".tp", ".tr", ".tt", ".tv", ".tw", ".tz", ".ua", ".ug", ".uk", ".um", ".us", ".uy", ".uz", ".va", ".vc", ".ve", ".vg", ".vi", ".vn", ".vu", ".wf", ".ws", ".ye", ".yt", ".yu", ".yr", ".za", ".zm", ".zw",
}
