package common

import "fmt"

// 判断主机是否存活
// host可以传入IP、域名
func IsHostAlived(host string) bool {
	// 运行前初始化当前系统类型
	OS := JudgeSystemType()
	fmt.Println(OS)
	return true
}

// 构造ICMP数据包格式
type ICMP struct {
	Type        uint8
	Code        uint8
	Checksum    uint16
	Identifier  uint16
	SequenceNum uint16
}

func CheckSum(data []byte) (rt uint16) {
	var (
		sum    uint32
		length = len(data)
		index  int
	)
	for length > 1 {
		sum +=
	}

}
