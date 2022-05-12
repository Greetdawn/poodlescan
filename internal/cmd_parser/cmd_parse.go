package cmdparser

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"regexp"
	"strconv"
	"strings"
)

// todo 根据命令行输入初始化全局变量
func (c *CMDPara) CMDUserInputParse() {

	// 接收参数：10.1.0.0/16或者域名
	flag.StringVar(&c.UserInputTargetString, "t", "", "设置扫描目标")
	flag.BoolVar(&c.BreakPingScan, "Pn", false, "跳过ping扫")
	flag.IntVar(&c.Threads, "T", 5, "设置并发，允许同时扫描几个目标")
	// todo Paras其他参数
	flag.Parse()
	tmpSlice, isIP := targetParse(c.UserInputTargetString)
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
func targetParse(targetInput string) (getTarget []string, isIP bool) {

	_, _, err := net.ParseCIDR(targetInput)
	if err != nil {
		tmp := net.ParseIP(targetInput)
		fmt.Println(tmp.String())
		if tmp == nil { // 不是IP，可能是网址或者xxx.xxx.xxx.xxx-xxx
			if len(matchDomain(targetInput)) != 0 {
				return matchDomain(targetInput), false
			} else {
				return matchIPRange(targetInput), true
			}

		} else { // 是ip
			return append(getTarget, tmp.String()), true
		}
	} else { // 是ip段,xxx.xxx.xxx.xxx/xx
		return cidrHosts(targetInput), true
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

// 匹配网址
func matchDomain(s string) []string {
	re := `([\w]([\w]{0,63}[\w])?\.)+[a-zA-Z]{2,6}\/?(\w+)?`
	reg := regexp.MustCompile(re)
	regRes := reg.FindAllString(s, -1)
	return regRes
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
