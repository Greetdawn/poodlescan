// @author: greetdawn
// @date: 2022-05-10
// 该文件主要实现端口扫描功能

package asset_host

import (
	"context"
	"fmt"
	"log"
	"net"
	"poodle/pkg/config"
	"poodle/pkg/logger"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Ullaakut/nmap/v2"
)

// 定义用哪种方法来扫描端口
type ScanPortMethod int

const (
	// 超时时间（秒）
	SCAN_PORT_TIME_OUT int = 5

	SPM_POODLE ScanPortMethod = 100
	SPM_NMAP   ScanPortMethod = 200
)

var (
	// 扫描端口时，用户指定使用的端口
	SCAN_PORT_PORTS []string

	// 选择用那种方法来扫描端口
	// 默认下用泰迪自身的方法扫描
	Scan_Port_Method ScanPortMethod = SPM_POODLE
)

func ScanHostOpenedPorts(target string) (portMap sync.Map) {
	switch Scan_Port_Method {
	case SPM_POODLE:
		return _TCPOrUDPPortScan_POODLE(target, config.GConfig.ScanPortConfig.Scan_Port_Proto, SCAN_PORT_PORTS...)
	case SPM_NMAP:
		return _TCPOrUDPPortScan_NMAP(target, config.GConfig.ScanPortConfig.Scan_Port_Proto, SCAN_PORT_PORTS...)
	default:
		return _TCPOrUDPPortScan_POODLE(target, config.GConfig.ScanPortConfig.Scan_Port_Proto)
	}
}

// 默认扫描方式
func _TCPOrUDPPortScan_POODLE(target, proto string, ports ...string) sync.Map {
	var (
		wg              sync.WaitGroup
		productPortChan chan string = make(chan string)
		resMap          sync.Map
	)
	// 如果传入端口则扫描指定端口
	if len(ports) != 0 {
		wg.Add(1)
		go func() {
			for _, p := range ports {
				productPortChan <- p
			}
			close(productPortChan)
			wg.Done()
		}()
	}

	// 并发扫描
	for i := 0; i < config.GConfig.ScanPortConfig.RunTaskThreads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for p := range productPortChan {
				// fmt.Printf("[*]正在扫描[%s]:%s端口\n", target, p)
				// 拼接地址
				addr := net.JoinHostPort(target, p)
				// 建立连接
				conn, err := net.DialTimeout(config.GConfig.ScanPortConfig.Scan_Port_Proto, addr, time.Duration(SCAN_PORT_TIME_OUT)*time.Second)
				// 如果连接正常建立
				if err == nil {
					var tmp [512]byte
					conn.SetReadDeadline(time.Now().Add(time.Second * 10))
					conn.Read(tmp[:])
					v := string(tmp[:])
					v = strings.ReplaceAll(v, "\r", "")
					v = strings.ReplaceAll(v, "\n", "")
					var data []byte
					for i := 0; i < len(v); i++ {
						if v[i] != 0 {
							data = append(data, byte(v[i]))
						} else {
							break
						}
					}

					v = string(data)
					resMap.Store(p, v)
					logger.LogInfo(fmt.Sprintf("%22s >> %15s端口 %s", logger.FgGreen(target), logger.FgGreen(p), logger.FgGreen(v)), logger.LOG_TERMINAL)
				}
			}
		}()
	}
	wg.Wait()
	return resMap
}

// scanPort接收tcp/udp/ack/syn/fin，设置扫描类型
// 指定范围扫：asset_host.TCPOrUDPPortScan("127.0.0.1", "tcp","1-65535")
// 指定端口扫：asset_host.TCPOrUDPPortScan("127.0.0.1", "tcp","443,444,445")
// 默认/精简端口扫：asset_host.TCPOrUDPPortScan("127.0.0.1", "tcp")
func _TCPOrUDPPortScan_NMAP(target string, scanProto string, ports ...string) sync.Map {

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
			//portRes := []string{port.Protocol, port.State.String(), port.Service.Name}
			resMap.Store(strconv.Itoa(int(port.ID)), port.Service.Name)
			//fmt.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
		}
	}

	return resMap
}
