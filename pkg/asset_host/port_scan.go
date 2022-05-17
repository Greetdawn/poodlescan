// @author: greetdawn
// @date: 2022-05-10
// 该文件主要实现端口扫描功能

package asset_host

import (
	"context"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/Ullaakut/nmap/v2"
)

var (
	proto string = "tcp"
)

func ScanHostOpenedPorts(target string) (portMap sync.Map) {
	// ports := []string{"22", "80", "8090", "8091"}
	// return _TCPOrUDPPortScan(target, ports...)

	return TCPOrUDPPortScan(target, proto)
}

// scanPort接收tcp/udp/ack/syn/fin，设置扫描类型
// 指定范围扫：asset_host.TCPOrUDPPortScan("127.0.0.1", "tcp","1-65535")
// 指定端口扫：asset_host.TCPOrUDPPortScan("127.0.0.1", "tcp","443,444,445")
// 默认/精简端口扫：asset_host.TCPOrUDPPortScan("127.0.0.1", "tcp")
func TCPOrUDPPortScan(target string, scanProto string, ports ...string) sync.Map {

	// resMap数据保存格式：{443:[TCP unfiltered https]}
	// resMap数据保存格式：map[int][]string
	var resMap sync.Map

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	nmapOpts := []nmap.Option{}
	switch {
	// 默认tcp
	case strings.ToLower(scanProto) == "tcp":
	case strings.ToLower(scanProto) == "udp":
		nmapOpts = append(nmapOpts, nmap.WithUDPScan())
	case strings.ToLower(scanProto) == "syn":
		nmapOpts = append(nmapOpts, nmap.WithSYNScan())
	case strings.ToLower(scanProto) == "ack":
		nmapOpts = append(nmapOpts, nmap.WithACKScan())
	case strings.ToLower(scanProto) == "fin":
		nmapOpts = append(nmapOpts, nmap.WithTCPFINScan())
	}

	if len(ports) != 0 {
		nmapOpts = append(nmapOpts, nmap.WithPorts(ports...))
	}
	nmapOpts = append(nmapOpts, nmap.WithTargets(target))
	nmapOpts = append(nmapOpts, nmap.WithContext(ctx))

	// with a 5 minute timeout.
	scanner, err := nmap.NewScanner(nmapOpts...)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	result, warnings, err := scanner.Run()
	if err != nil {
		log.Fatalf("unable to run nmap scan: %v", err)
	}

	if warnings != nil {
		log.Printf("Warnings: \n %v", warnings)
	}

	// Use the results to print an example output
	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		for _, port := range host.Ports {
			portRes := []string{port.Protocol, port.State.String(), port.Service.Name}
			resMap.Store(int(port.ID), portRes)
			//fmt.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
		}
	}

	return resMap
}
