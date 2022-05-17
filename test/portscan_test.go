package test

import (
	"fmt"
	"poodle/pkg/asset_host"
	"testing"
)

func TestMain(m *testing.M) {

	// tcp和udp必须小写
	// port不传参默认全端口
	//scanPort := []string{"80"}
	res := asset_host.TCPOrUDPPortScan("127.0.0.1", "tcp", 5, 2000)

	// 输出结果集
	fmt.Println()
	res.Range(func(key, value interface{}) bool {
		k := key.(string)
		v := value.(string)
		fmt.Println(k, " "+v)
		return true
	})
	//m.Run()
}

// func TestExample1(t *testing.T) {
// 	fmt.Println("其他测试函数1")
// }

// func TestExample2(t *testing.T) {
// 	fmt.Println("其他测试函数2")
// }
