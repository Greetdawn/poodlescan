package cmdparser

import (
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"poodle/pkg/common"
	"poodle/pkg/logger"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var TaskChannal chan *TASKUint

type _Target struct {
	Target string //目标
	IsIp   bool   //目标是否是IP
}

// 模块变量初始化
func init() {
	sort.Strings(DomainArray)
}

// todo 根据命令行输入初始化全局变量
func (this *TerminalParams) ParseUserInput() {
	// -t : 设置用户提供的目标
	flag.StringVar(&this.UserInputTargetString, "t", "", "设置扫描目标")

	// -Pn : 如果加了这个参数，则表示跳过Ping扫
	flag.BoolVar(&this.IsPn, "Pn", false, "true:跳过ping扫;false（默认）:不跳过Ping扫")

	// -T : 用户设置线程
	flag.IntVar(&this.ThreadsNumber, "T", 5, "设置并发，允许同时扫描几个目标")
	// todo Paras其他参数

	// 开始解析
	flag.Parse()

	// debug信息
	logger.LogInfo("参数：用户设置目标："+this.UserInputTargetString, logger.LOG_TERMINAL)
	logger.LogInfo(fmt.Sprintf("参数：是否跳过Ping扫描：%t", this.IsPn), logger.LOG_TERMINAL)
	logger.LogInfo("参数：线程数："+strconv.Itoa(this.ThreadsNumber), logger.LOG_TERMINAL)
}

// 解析TerminalParams结构体
func (this *TerminalParams) ParseTerminal() bool {
	// 1. 解析用户的输入目标
	targets, err := parseUserInputTargetString(this.UserInputTargetString)
	if err != nil {
		// 解析用户的输入目标发生错误。
		logger.LogError(fmt.Sprintf("err: %v\n", err), logger.LOG_TERMINAL_FILE)
		return false
	}
	// 2. 将目标转化成 Task ，写入通道中。
	TaskChannal = make(chan *TASKUint, 10000)
	for _, target := range targets {
		if target.IsIp {
			TaskChannal <- &TASKUint{target.Target, TASKUint_TargetType_IP, 1} //控制码1，暂时不用
		} else {
			TaskChannal <- &TASKUint{target.Target, TASKUint_TargetType_Domain, 1} //控制码1，暂时不用
		}
	}
	close(TaskChannal)
	return true
}

// 测试。执行命令
func (this *TerminalParams) PrintTask() {
	var tmps sync.WaitGroup
	tmps.Add(this.ThreadsNumber)
	for j := 0; j < this.ThreadsNumber; j++ {
		go func() {
			defer tmps.Done()
			var end = true
			for end {
				select {
				case target, ok := <-TaskChannal:
					if !ok {
						//logger.LogWarn("TaskChannal 通道已关闭", logger.LOG_TERMINAL)
						end = false
					} else {
						common.IsHostAlived(target.Target)
						// fmt.Println(target)
					}
				default:

				}

			}
		}()
	}
	tmps.Wait()
}

// 解析用户输入的“目标”字符串
func parseUserInputTargetString(targetInput string) (targets []_Target, err error) {
	// 如果目标字符串中没有.(dot)认为目标无效
	if !strings.Contains(targetInput, ".") {
		err = errors.New("invalid target.1")
		return nil, err
	}

	_, _, cidrError := net.ParseCIDR(targetInput)
	if cidrError != nil {
		// 尝试将目标解析为IP
		result := net.ParseIP(targetInput)
		if result != nil {
			// 成功将目标解析为IP，保存目标的IP
			return append(targets, _Target{targetInput, true}), nil
		} else {
			// 检查是否为域名
			// 当目标主机以常见域名格式结尾时，认为是个域名
			parts := strings.Split(targetInput, ".")
			endDomain := "." + parts[len(parts)-1]
			index := sort.SearchStrings(DomainArray, endDomain)
			if index < len(DomainArray) && DomainArray[index] == "."+parts[len(parts)-1] {
				// 认为合法域名
				return append(targets, _Target{targetInput, false}), nil
			} else if isIpRange(targetInput) {
				return ipRange2Targets(targetInput), nil
			} else {
				err = errors.New("invalid target.")
				return nil, err
			}
		}
	} else {
		// 是ip段,xxx.xxx.xxx.xxx/xx
		return ipSeg2Targets(targetInput), nil
	}
}

// 将用户输入的IP段转成目标
func ipSeg2Targets(netw string) []_Target {
	// convert string to IPNet struct
	_, ipv4Net, err := net.ParseCIDR(netw)
	if err != nil {
		log.Fatal(err)
	}
	// convert IPNet struct mask and address to uint32
	mask := binary.BigEndian.Uint32(ipv4Net.Mask)
	// find the start IP address
	start := binary.BigEndian.Uint32(ipv4Net.IP)
	// find the final IP address
	finish := (start & mask) | (mask ^ 0xffffffff)
	// make a slice to return host addresses
	var targets []_Target
	// loop through addresses as uint32.
	// I used "start + 1" and "finish - 1" to discard the network and broadcast addresses.
	for i := start + 1; i <= finish-1; i++ {
		// convert back to net.IPs
		// Create IP address of type net.IP. IPv4 is 4 bytes, IPv6 is 16 bytes.
		ip := make(net.IP, 4)
		binary.BigEndian.PutUint32(ip, i)
		targets = append(targets, _Target{ip.String(), true})
	}
	// return a slice of strings containing IP addresses
	return targets
}

// 将用户输入的IP范围转成目标
// 匹配ip段，192.168.1.1-20，返回有效ip列表
func ipRange2Targets(s string) (targets []_Target) {
	// 存放未检测有效性的ip地址池
	var tmpSlice []string

	re := `\d+\.\d+\.\d+\.\d+-\d+`
	reg := regexp.MustCompile(re)
	regRes := reg.FindAllString(s, -1)

	for _, v := range regRes {
		vSplitByLine := strings.Split(v, "-")
		vSplitByLine0_SplitByDot := strings.Split(vSplitByLine[0], ".")
		ipSuffix := vSplitByLine0_SplitByDot[0] + "." + vSplitByLine0_SplitByDot[1] + "." + vSplitByLine0_SplitByDot[2] + "."
		num1, _ := strconv.Atoi(vSplitByLine0_SplitByDot[3])
		num2, _ := strconv.Atoi(vSplitByLine[1])
		if num1 > num2 {
			for i := num2; i <= num1; i++ {
				tmpSlice = append(tmpSlice, ipSuffix+strconv.Itoa(i))
			}

		} else {
			for i := num1; i <= num2; i++ {
				tmpSlice = append(tmpSlice, ipSuffix+strconv.Itoa(i))
			}
		}
	}
	for _, v := range tmpSlice {
		if net.ParseIP(v) != nil {
			targets = append(targets, _Target{v, true})
		}
	}
	return targets
}

// 判断用户输入的是否是合法的IP范围
func isIpRange(ipstr string) bool {
	part2 := strings.Split(ipstr, "-")
	if len(part2) != 2 {
		return false
	}
	_, e1 := strconv.ParseFloat(part2[1], 64)
	if e1 != nil {
		return false
	}

	part4 := strings.Split(part2[0], ".")
	if len(part4) != 4 {
		return false
	}
	for _, v := range part4 {
		_, e1 := strconv.ParseFloat(v, 64)
		if e1 != nil {
			return false
		}
	}
	return true
}