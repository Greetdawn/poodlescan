package os

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/Ullaakut/nmap/v2"
)

// 判断系统中是否安装nmap
func IsHavaNMap() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	nmapOpts := []nmap.Option{nmap.WithUDPScan(), nmap.WithTargets("127.0.0.1"), nmap.WithContext(ctx)}
	scanner, err := nmap.NewScanner(nmapOpts...)
	if err != nil {
		fmt.Printf("unable to create nmap scanner: %v\n", err)
		return false
	}
	_, warnings, err := scanner.Run()
	if err != nil {
		fmt.Printf("unable to run nmap scan: %v\n", err)
	}

	if warnings != nil {
		log.Printf("Warnings: \n %v", warnings)
	}
	return true
}

func InstallNmap() {
	optSys := runtime.GOOS
	fmt.Println(optSys)
}
