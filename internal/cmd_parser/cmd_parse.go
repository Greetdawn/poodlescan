package cmdparser

import (
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// 模块变量初始化
func init() {
	sort.Strings(DomainArray)
}

// todo 根据命令行输入初始化全局变量
func (c *CMDPara) CMDUserInputParse() {

	// 接收参数：10.1.0.0/16或者域名
	flag.StringVar(&c.UserInputTargetString, "t", "", "设置扫描目标")
	flag.BoolVar(&c.BreakPingScan, "Pn", false, "跳过ping扫")
	flag.IntVar(&c.Threads, "T", 5, "设置并发，允许同时扫描几个目标")
	// todo Paras其他参数
	flag.Parse()

	fmt.Println(c.UserInputTargetString)
	fmt.Println(c.BreakPingScan)
	fmt.Println(c.Threads)

	tmpSlice, isIP, err := targetParse(c.UserInputTargetString)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	c.isIP = isIP
	if isIP {
		c.IpList = append(c.IpList, tmpSlice...)

	} else {
		c.DomainList = append(c.DomainList, tmpSlice...)
	}
}

// 解析如下输入：
// 192.168.0.1-20
// 192.168.0.1/24
// http://www.baidu.com
// www.baidu.com
func targetParse(targetInput string) (getTarget []string, isIP bool, err1 error) {
	// 如果目标字符串中没有.(dot)认为目标无效
	if !strings.Contains(targetInput, ".") {
		err1 = errors.New("invalid target.1")
		return nil, false, err1
	}

	_, _, cidrError := net.ParseCIDR(targetInput)

	if cidrError != nil {
		fmt.Printf("cidrError: %v\n", cidrError)
		// 判断是否存在输入错误
		// if strings.HasPrefix(strings.ToLower(cidrError.Error()), strings.ToLower("invalid CIDR address:")) {
		// 	err1 = errors.New("invalid target.2")
		// 	return nil, false, err1
		// }

		// 尝试将目标解析为IP
		result := net.ParseIP(targetInput)
		if result != nil {
			// 成功将目标解析为IP，保存目标的IP
			fmt.Println("ss:::IP")
			return append(getTarget, result.String()), true, nil
		} else {
			// 检查是否为域名
			// 当目标主机以常见域名格式结尾时，认为是个域名
			parts := strings.Split(targetInput, ".")
			endDomain := "." + parts[len(parts)-1]
			index := sort.SearchStrings(DomainArray, endDomain)
			if index < len(DomainArray) && DomainArray[index] == "."+parts[len(parts)-1] {
				// 认为合法域名
				fmt.Printf("%s\n", " 确认:::合法域名")
				return append(getTarget, targetInput), false, nil
			} else if isIpSeg(targetInput) {
				fmt.Println("确认:::合法IP段")
				return matchIPRange(targetInput), true, nil
			} else {
				err1 = errors.New("invalid target.")
				return nil, false, err1
			}
		}
	} else { // 是ip段,xxx.xxx.xxx.xxx/xx
		fmt.Println("猜测:::IP段")
		return cidrHosts(targetInput), true, nil
	}
}

func cidrHosts(netw string) []string {
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
	var hosts []string
	// loop through addresses as uint32.
	// I used "start + 1" and "finish - 1" to discard the network and broadcast addresses.
	for i := start + 1; i <= finish-1; i++ {
		// convert back to net.IPs
		// Create IP address of type net.IP. IPv4 is 4 bytes, IPv6 is 16 bytes.
		ip := make(net.IP, 4)
		binary.BigEndian.PutUint32(ip, i)
		hosts = append(hosts, ip.String())
	}
	// return a slice of strings containing IP addresses
	return hosts
}

// 匹配ip段，192.168.1.1-20，返回有效ip列表
func matchIPRange(s string) (ipSlice []string) {
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
			ipSlice = append(ipSlice, v)
		}
	}
	return
}

func isIpSeg(ipstr string) bool {
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
