// @author: greetdawn
// @date: 2022-05-10
// 该文件主要实现端口扫描功能

package asset_host

import (
	"context"
	"fmt"
	"log"
	"net"
	"poodle/pkg/logger"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Ullaakut/nmap/v2"
)

type ScanPortMethod int

const (
	SPM_POODLE ScanPortMethod = 100
	SPM_NMAP   ScanPortMethod = 200
)

var (
	// 指定扫描端口时使用的协议。可选项有：tcp、udp、syn、ack、fin
	Scan_Port_Proto string = "tcp"

	// 扫描端口时，默认的扫描的端口
	SCAN_PORT_POODLE_COMMON_PORTS []string = []string{
		"1", "7", "9", "13", "19", "21", "22", "23", "25", "37", "42", "49", "53", "69", "79", "80", "81", "85", "105", "113", "123", "135",
		"137", "138", "139", "143", "161", "179", "222", "264", "384", "389", "402", "407", "443", "445", "446", "465", "500", "502", "512",
		"513", "514", "515", "523", "524", "540", "548", "554", "587", "617", "623", "689", "705", "771", "783", "873", "888", "902", "910",
		"912", "921", "993", "995", "998", "1000", "1024", "1030", "1035", "1090", "1098", "1099", "1100", "1101", "1102", "1103", "1128",
		"1129", "1158", "1199", "1211", "1220", "1234", "1241", "1300", "1311", "1352", "1433", "14344", "1435", "1440", "1494", "1521",
		"1530", "1533", "1581", "1582", "1604", "1720", "1723", "1755", "1811", "1900", "2000", "2001", "2049", "2082", "2083", "2100", "2103",
		"2121", "2199", "2207", "2222", "2323", "2362", "2375", "2380", "2381", "2525", "2533", "2598", "2601", "2604", "2638", "2809", "2947",
		"2967", "3000", "3037", "3050", "3057", "3128", "3200", "3217", "3273", "3299", "3306", "3311", "3312", "3389", "3460", "3500", "3628",
		"3632", "3690", "3780", "3790", "3817", "4000", "4322", "4433", "4444", "4445", "4659", "4679", "4848", "5000", "5038", "5040", "5051",
		"5060", "5061", "5093", "5168", "5247", "5250", "5351", "5353", "5355", "5400", "5405", "5432", "5433", "5498", "5520", "5521", "5554",
		"5555", "5560", "5580", "5601", "5631", "5632", "5666", "5800", "5814", "5900", "5901", "5902", "5903", "5904", "5905", "5906", "5907",
		"5908", "5909", "5910", "5920", "5984", "5985", "5986", "6000", "6050", "6060", "6070", "6080", "6082", "6101", "6106", "6112", "6262",
		"6379", "6405", "6502", "6503", "6504", "6542", "6660", "6661", "6667", "6905", "6988", "7001", "7021", "7071", "7080", "7144", "7181",
		"7210", "7443", "7510", "7579", "7580", "7700", "7770", "7777", "7778", "7787", "7800", "7801", "7879", "7902", "8000", "8001", "8008",
		"8014", "8020", "8023", "8028", "8030", "8080", "8081", "8082", "8087", "8090", "8095", "8161", "8180", "8205", "8222", "8300", "8303",
		"8333", "8400", "8443", "8444", "8503", "8800", "8812", "8834", "8880", "8888", "8889", "8890", "8899", "8901", "8902", "8903", "9000",
		"9002", "9060", "9080", "9081", "9084", "9090", "9099", "9100", "9111", "9152", "9200", "9390", "9391", "9443", "9495", "9809", "9810",
		"9811", "9812", "9813", "9814", "9815", "9855", "9999", "10000", "10001", "10008", "10050", "10051", "10080", "10098", "10162", "10202",
		"10203", "10443", "10616", "10628", "11000", "11099", "11211", "11234", "11333", "12174", "12203", "12221", "12345", "12397", "12401",
		"13364", "13500", "13838", "14330", "15200", "16102", "17185", "17200", "18881", "19300", "19810", "20010", "20031", "20034", "20101",
		"20111", "20171", "20222", "22222", "23472", "23791", "23943", "25000", "25025", "26000", "26122", "27000", "27017", "27888", "28222",
		"28784", "30000", "30718", "31001", "31099", "32764", "32913", "34205", "34443", "37718", "38080", "38292", "40007", "41025", "41080",
		"41523", "41524", "44334", "44818", "45230", "46823", "46824", "47001", "47002", "48899", "49152", "50000", "50001", "50002", "50003",
		"50004", "50013", "50500", "50501", "50502", "50503", "50504", "52302", "55553", "57772", "62078", "62514", "65535"}
	// 扫描端口时，用户指定使用的端口
	SCAN_PORT_PORTS []string

	// 线程数
	Scan_Port_Threads int = 200
	// 超时时间（秒）
	Scan_Port_Time_Out int = 5

	// 选择用那种方法来扫描端口
	// 默认下用泰迪自身的方法扫描
	Scan_Port_Method ScanPortMethod = SPM_NMAP
)

func ScanHostOpenedPorts(target string) (portMap sync.Map) {
	switch Scan_Port_Method {
	case SPM_POODLE:
		return _TCPOrUDPPortScan_POODLE(target, Scan_Port_Proto, SCAN_PORT_PORTS...)
	case SPM_NMAP:
		return _TCPOrUDPPortScan_NMAP(target, Scan_Port_Proto, SCAN_PORT_PORTS...)
	default:
		return _TCPOrUDPPortScan_POODLE(target, Scan_Port_Proto)
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
	for i := 0; i < Scan_Port_Threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for p := range productPortChan {
				// fmt.Printf("[*]正在扫描[%s]:%s端口\n", target, p)
				// 拼接地址
				addr := net.JoinHostPort(target, p)
				// 建立连接
				conn, err := net.DialTimeout(Scan_Port_Proto, addr, time.Duration(Scan_Port_Time_Out)*time.Second)
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
