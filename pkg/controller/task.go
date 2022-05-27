package controller

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"net"
	"poodle/pkg/asset_host"
	"poodle/pkg/common"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

//************ 全局变量定义区 ************
// 任务通道
var G_AssetTaskChannal chan *TASKUint

// 控制项,是否生成任务，用于取消任务
var isGenerateTask bool = true

//**************************************

type stTarget struct {
	Target string //目标
	IsIp   bool   //目标是否是IP
}

type TASKUint struct {
	// TUID       int               //当前任务编号
	// TotalNum   int               //任务总数
	Target     string            // 用户输入的扫描目标
	TargetIsIP bool              // 目标表示类型，IP或者域名
	Params     map[string]string // 用到的参数
}

// 命令行解析，生成任务，将任务写入任务通道中
func GenrateTasks(tp *common.TaskPacket) (tasknumber int, err error) {
	fmt.Printf("tp: %v\n", tp)
	// 3. 解析用户的输入目标
	targets, err := parseUserInputTargetString(tp.UserInputTargetString)
	if err != nil {
		// 解析用户的输入目标发生错误。
		return tasknumber, err
	}

	// 4. 将目标转化成 Task ，写入通道中。
	// 新建任务通道
	G_AssetTaskChannal = make(chan *TASKUint, 10000)
	for tasknumber = 0; isGenerateTask && tasknumber < len(targets); tasknumber++ {
		newTask := new(TASKUint)
		newTask.Params = make(map[string]string)
		newTask.Target = targets[tasknumber].Target
		// newTask.ControlCode = cc
		newTask.Params["ports"] = tp.PortList
		if targets[tasknumber].IsIp {
			newTask.TargetIsIP = true
		} else {
			newTask.TargetIsIP = false
		}
		G_AssetTaskChannal <- newTask
		fmt.Print(".")
	}
	close(G_AssetTaskChannal)
	fmt.Println()
	return tasknumber, nil
}

// 解析用户输入的“目标”字符串
func parseUserInputTargetString(targetInput string) (targets []stTarget, err error) {
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
			return append(targets, stTarget{targetInput, true}), nil
		} else {
			// 检查是否为域名
			// 当目标主机以常见域名格式结尾时，认为是个域名
			parts := strings.Split(targetInput, ".")
			endDomain := "." + parts[len(parts)-1]
			index := sort.SearchStrings(common.DOMAINARRAY, endDomain)
			if index < len(common.DOMAINARRAY) && common.DOMAINARRAY[index] == "."+parts[len(parts)-1] {
				// 认为合法域名
				return append(targets, stTarget{targetInput, false}), nil
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
func ipSeg2Targets(netw string) []stTarget {
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
	var targets []stTarget
	// loop through addresses as uint32.
	// I used "start + 1" and "finish - 1" to discard the network and broadcast addresses.
	for i := start + 1; i <= finish-1; i++ {
		// convert back to net.IPs
		// Create IP address of type net.IP. IPv4 is 4 bytes, IPv6 is 16 bytes.
		ip := make(net.IP, 4)
		binary.BigEndian.PutUint32(ip, i)
		targets = append(targets, stTarget{ip.String(), true})
	}
	// return a slice of strings containing IP addresses
	return targets
}

// 将用户输入的IP范围转成目标
// 匹配ip段，192.168.1.1-20，返回有效ip列表
func ipRange2Targets(s string) (targets []stTarget) {
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
			targets = append(targets, stTarget{v, true})
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

// 将端口的描述性字符串转成端口列表
func getPorts(str string) (ports []string, err error) {
	if str == "top1000" {
		ports = asset_host.SCAN_PORT_POODLE_COMMON_PORTS
		return
	}

	// 端口段
	if strings.Contains(str, "-") {
		strs := strings.Split(str, "-")
		p0, e0 := strconv.Atoi(strs[0])
		p1, e1 := strconv.Atoi(strs[1])
		if e0 != nil || e1 != nil {
			err = errors.New("端口段格式错误。端口段输入例如：1000-2000")
			return
		}
		if p0 > p1 {
			tmp := p1
			p1 = p0
			p0 = tmp
		}
		for ; p0 <= p1; p0++ {
			ports = append(ports, strconv.Itoa(p0))
		}
		return
	}
	return
}
