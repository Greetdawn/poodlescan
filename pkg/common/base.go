package common

type TASKUint_TargetType byte

const (
	TASKUint_TargetType_IP     TASKUint_TargetType = 1
	TASKUint_TargetType_Domain TASKUint_TargetType = 2
)

// 全局变量，任务通道
var G_TaskChannal chan *TASKUint

type TASKUint struct {
	Target      string              // 用户输入的扫描目标
	TargetType  TASKUint_TargetType // 目标表示类型，IP或者域名
	ControlCode uint                // 控制码
}

////////////////////////////////////////////////////////////////////////////////////
// 定义常见的域名后缀
// 新增.top后缀
var DOMAINARRAY []string = []string{
	".com", ".edu", ".gov", ".int", ".mil", ".net", ".org", ".biz", ".info", ".pro", ".name", ".museum", ".coop", ".aero", ".xxx", ".idv", ".ac", ".ad", ".ae", ".af", ".ag", ".ai", ".al", ".am", ".an", ".ao", ".aq", ".ar", ".as", ".at", ".au", ".aw", ".az", ".ba", ".bb", ".bd", ".be", ".bf", ".bg", ".bh", ".bi", ".bj", ".bm", ".bn", ".bo", ".br", ".bs", ".bt", ".bv", ".bw", ".by", ".bz", ".ca", ".cc", ".cd", ".cf", ".cg", ".ch", ".ci", ".ck", ".cl", ".cm", ".cn", ".co", ".cr", ".cu", ".cv", ".cx", ".cy", ".cz", ".de", ".dj", ".dk", ".dm", ".do", ".dz", ".ec", ".ee", ".eg", ".eh", ".er", ".es", ".et", ".eu", ".fi", ".fj", ".fk", ".fm", ".fo", ".fr", ".ga", ".gd", ".ge", ".gf", ".gg", ".gh", ".gi", ".gl", ".gm", ".gn", ".gp", ".gq", ".gr", ".gs", ".gt", ".gu", ".gw", ".gy", ".hk", ".hm", ".hn", ".hr", ".ht", ".hu", ".id", ".ie", ".il", ".im", ".in", ".io", ".iq", ".ir", ".is", ".it", ".je", ".jm", ".jo", ".jp", ".ke", ".kg", ".kh", ".ki", ".km", ".kn", ".kp", ".kr", ".kw", ".ky", ".kz", ".la", ".lb", ".lc", ".li", ".lk", ".lr", ".ls", ".lt", ".lu", ".lv", ".ly", ".ma", ".mc", ".md", ".mg", ".mh", ".mk", ".ml", ".mm", ".mn", ".mo", ".mp", ".mq", ".mr", ".ms", ".mt", ".mu", ".mv", ".mw", ".mx", ".my", ".mz", ".na", ".nc", ".ne", ".nf", ".ng", ".ni", ".nl", ".no", ".np", ".nr", ".nu", ".nz", ".om", ".pa", ".pe", ".pf", ".pg", ".ph", ".pk", ".pl", ".pm", ".pn", ".pr", ".ps", ".pt", ".pw", ".py", ".qa", ".re", ".ro", ".ru", ".rw", ".sa", ".sb", ".sc", ".sd", ".se", ".sg", ".sh", ".si", ".sj", ".sk", ".sl", ".sm", ".sn", ".so", ".sr", ".st", ".sv", ".sy", ".sz", ".tc", ".td", ".tf", ".tg", ".th", ".tj", ".tk", ".tl", ".tm", ".tn", ".to", ".tp", ".tr", ".tt", ".tv", ".tw", ".tz", ".ua", ".ug", ".uk", ".um", ".us", ".uy", ".uz", ".va", ".vc", ".ve", ".vg", ".vi", ".vn", ".vu", ".wf", ".ws", ".ye", ".yt", ".yu", ".yr", ".za", ".zm", ".zw", ".top",
}

// uint32每个bit代表的功能
// CC -- Control Code
// 0: Ping扫功能
const CC_PING_SCAN uint = 1

// 全局变量，保存终端命令行参数结构体
var G_TerminalParam TerminalParams

// 命令行参数结构体
// 在用户输入后，通过flag模块，将用户输入的命令行转换成结构体保存
type TerminalParams struct {
	// 用户输入的扫描目标，原始字符串
	UserInputTargetString string
	// 标志;一些特殊的标志
	// 000000000 默认情况，保留
	// 000000001 目标从文件中读取
	Flag byte
	// 用户设置的线程数
	ThreadsNumber int // 总线程数，同时扫描几个目标
	// Pn "跳过Ping扫"，默认不跳过Ping扫,false
	IsPn bool
}

// 获取一个TerminalParams对象
func GetTerminalParamObj() *TerminalParams {
	return &TerminalParams{}
}

// 根据终端参数结构体生成控制码
func (this *TerminalParams) GenerateControlCode() (controlCode uint) {
	controlCode = 0

	// -Pn 	跳过主机存活检测
	// 默认不跳过
	if !this.IsPn {
		controlCode |= CC_PING_SCAN
	}

	return controlCode
}
