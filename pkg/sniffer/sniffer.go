package sniffer

// sniffer的接口类
type iSniffer interface {
	// 开始执行嗅探的工作。
	StartSniff()

	// 保存嗅探出来的结果。目前版本只需要把结果打印出来即可。
	SaveInfo()

	// 将对象转化成字符串输出
	toString()
}

// 嗅探器的父类，对外只暴露这个类。
type Sniffer struct {
	// 命令码
	CmdCode int
}
