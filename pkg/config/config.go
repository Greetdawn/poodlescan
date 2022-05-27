package config

import "strconv"

// 内核配置
var KnelConfig KernelConfig

// 配置项
const (
	runTaskThreads string = "runTaskThreads"
	isPrintLogInfo string = "isPrintLogInfo"
)

// 内核配置
type KernelConfig struct {
	// 执行任务的线程数
	RunTaskThreads int
	// 是否打开日志
	IsPrintLogInfo bool
}

// 配置命令的包
type ConfigPacket struct {
	Key   string // 对应的配置项
	Value string
}

func NewDefaultKernelConfig() KernelConfig {
	return KernelConfig{
		RunTaskThreads: 200,
	}
}

func (c *ConfigPacket) UpdateConfig() (err error) {
	switch c.Key {
	case runTaskThreads:
		KnelConfig.RunTaskThreads, err = strconv.Atoi(c.Value)
	case isPrintLogInfo:
		KnelConfig.IsPrintLogInfo, err = strconv.ParseBool(c.Value)
	}
	return err
}
