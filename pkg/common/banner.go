package common

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var banner = `
██████╗  ██████╗  ██████╗ ██████╗ ██╗     ███████╗███████╗ ██████╗ █████╗ ███╗   ██╗
██╔══██╗██╔═══██╗██╔═══██╗██╔══██╗██║     ██╔════╝██╔════╝██╔════╝██╔══██╗████╗  ██║
██████╔╝██║   ██║██║   ██║██║  ██║██║     █████╗  ███████╗██║     ███████║██╔██╗ ██║
██╔═══╝ ██║   ██║██║   ██║██║  ██║██║     ██╔══╝  ╚════██║██║     ██╔══██║██║╚██╗██║
██║     ╚██████╔╝╚██████╔╝██████╔╝███████╗███████╗███████║╚██████╗██║  ██║██║ ╚████║
╚═╝      ╚═════╝  ╚═════╝ ╚═════╝ ╚══════╝╚══════╝╚══════╝ ╚═════╝╚═╝  ╚═╝╚═╝  ╚═══╝
`

func ShowBanner() {
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

	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println(banner)
	fmt.Printf("\t\t\t\t\t\t\t\t\tBy Teams\n\n")
}
