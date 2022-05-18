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
	Params      map[string]string   // 用到的参数
}

////////////////////////////////////////////////////////////////////////////////////
// 定义常见的域名后缀
// 新增.top后缀
var DOMAINARRAY []string = []string{
	".com", ".edu", ".gov", ".int", ".mil", ".net", ".org", ".biz", ".info", ".pro", ".name", ".museum", ".coop", ".aero", ".xxx", ".idv", ".ac", ".ad", ".ae", ".af", ".ag", ".ai", ".al", ".am", ".an", ".ao", ".aq", ".ar", ".as", ".at", ".au", ".aw", ".az", ".ba", ".bb", ".bd", ".be", ".bf", ".bg", ".bh", ".bi", ".bj", ".bm", ".bn", ".bo", ".br", ".bs", ".bt", ".bv", ".bw", ".by", ".bz", ".ca", ".cc", ".cd", ".cf", ".cg", ".ch", ".ci", ".ck", ".cl", ".cm", ".cn", ".co", ".cr", ".cu", ".cv", ".cx", ".cy", ".cz", ".de", ".dj", ".dk", ".dm", ".do", ".dz", ".ec", ".ee", ".eg", ".eh", ".er", ".es", ".et", ".eu", ".fi", ".fj", ".fk", ".fm", ".fo", ".fr", ".ga", ".gd", ".ge", ".gf", ".gg", ".gh", ".gi", ".gl", ".gm", ".gn", ".gp", ".gq", ".gr", ".gs", ".gt", ".gu", ".gw", ".gy", ".hk", ".hm", ".hn", ".hr", ".ht", ".hu", ".id", ".ie", ".il", ".im", ".in", ".io", ".iq", ".ir", ".is", ".it", ".je", ".jm", ".jo", ".jp", ".ke", ".kg", ".kh", ".ki", ".km", ".kn", ".kp", ".kr", ".kw", ".ky", ".kz", ".la", ".lb", ".lc", ".li", ".lk", ".lr", ".ls", ".lt", ".lu", ".lv", ".ly", ".ma", ".mc", ".md", ".mg", ".mh", ".mk", ".ml", ".mm", ".mn", ".mo", ".mp", ".mq", ".mr", ".ms", ".mt", ".mu", ".mv", ".mw", ".mx", ".my", ".mz", ".na", ".nc", ".ne", ".nf", ".ng", ".ni", ".nl", ".no", ".np", ".nr", ".nu", ".nz", ".om", ".pa", ".pe", ".pf", ".pg", ".ph", ".pk", ".pl", ".pm", ".pn", ".pr", ".ps", ".pt", ".pw", ".py", ".qa", ".re", ".ro", ".ru", ".rw", ".sa", ".sb", ".sc", ".sd", ".se", ".sg", ".sh", ".si", ".sj", ".sk", ".sl", ".sm", ".sn", ".so", ".sr", ".st", ".sv", ".sy", ".sz", ".tc", ".td", ".tf", ".tg", ".th", ".tj", ".tk", ".tl", ".tm", ".tn", ".to", ".tp", ".tr", ".tt", ".tv", ".tw", ".tz", ".ua", ".ug", ".uk", ".um", ".us", ".uy", ".uz", ".va", ".vc", ".ve", ".vg", ".vi", ".vn", ".vu", ".wf", ".ws", ".ye", ".yt", ".yu", ".yr", ".za", ".zm", ".zw", ".top",
}

// uint32每个bit代表的功能
// CC -- Control Code
const (
	// 0: Ping扫功能
	CC_PING_SCAN uint = 1
	// 1: 端口扫描
	CC_PORT_SCAN uint = 2
	// 2: 子域探测功能
	CC_SUB_DOMAIN_SCAN uint = 4
	// 3: 爬虫功能
	CC_SPIDER uint = 8
	// 4: 指纹识别功能
	CC_FINGERPRINT uint = 16
	// 5: 保留
	// 6: 保留
	// 7: 保留

	// 漏洞扫描模块
	// 8: 专项漏洞扫描功能, 针对具体漏洞编号扫描
	CC_VULSCAN_ID uint = 256
	// 9: 专项漏洞扫描功能, 针对具体的漏洞类型扫描（例如：weblogic）
	CC_VULSCAN_TYPE uint = 512
	//10: 专项漏洞扫描功能, 根据用户自定义漏洞文件路径扫描
	CC_VULSCAN_FILE uint = 1024
	//11: 专项漏洞扫描功能, 针对web应用开启目录探测功能
	CC_VULSCAN_DIRSEARCH uint = 2048
	//12: 服务端口爆破功能, 针对应用服务端口进行口令爆破
	CC_VULSCAN_BURST uint = 4096
	//13: 全漏洞扫描功能，针对单一目标实现平台全poc扫描
	CC_VULSCAN uint = 8192
	//14: 保留
	//15: 保留

	// 漏洞利用模块
	//16:漏洞利用功能
	CC_VULEXPLOIT uint = 65536
)
