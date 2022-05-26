package common

import "strconv"

var (
	// 执行任务的线程数
	G_RunTaskThreads int = 200
)

// 配置项
const (
	runTaskThreads string = "runTaskThreads"
)

// 配置命令的包
type ConfigPacket struct {
	Key   string // 对应的配置项
	Value string // 值
}

func (c *ConfigPacket) UpdateConfig() (err error) {
	switch c.Key {
	case runTaskThreads:
		G_RunTaskThreads, err = strconv.Atoi(c.Value)
	}
	return err
}
