package common

import (
	"os"
	"os/exec"
	"runtime"
)

// 清屏
func Clear() {
	optSys := runtime.GOOS
	if optSys == "linux" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func Exec(cmd string) {
	optSys := runtime.GOOS
	if optSys == "linux" {
		cmd := exec.Command(cmd)
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("cmd", "/c", cmd)
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
