package test

import (
	"fmt"
	"poodle/internal/cmdinit"
	"testing"
)

func TestMain(m *testing.M) {

	cmdPars := cmdinit.CMDParseInit()
	cmdPars.CMDUserInputParse(cmdPars)
	fmt.Println(cmdPars.IsPingScan)
	fmt.Println(cmdPars.UserInputTargetString)
	//m.Run()
}

// func TestExample1(t *testing.T) {
// 	fmt.Println("其他测试函数1")
// }

// func TestExample2(t *testing.T) {
// 	fmt.Println("其他测试函数2")
// }
