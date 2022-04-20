package test

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("此处进行单元测试")
	fmt.Println("在此处使用go test命令进行单元测试")
	fmt.Println("测试其他子方法")
	m.Run()
}

func TestExample1(t *testing.T) {
	fmt.Println("其他测试函数1")
}

func TestExample2(t *testing.T) {
	fmt.Println("其他测试函数2")
}
