package controllor

import (
	"encoding/binary"
	"errors"
	"log"
	"net"
	"poodle/pkg/asset_host"
	"poodle/pkg/common"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
)

//(alivedList asset_host.AssetHost, diedList asset_host.AssetHost, err error)
func Run(threadNum int) {
	var tmps sync.WaitGroup
	tmps.Add(threadNum)
	for i := 0; i < threadNum; i++ {
		go func() {
			defer tmps.Done()
			var end = true
			for end {
				select {
				case target, ok := <-common.G_TaskChannal:
					if !ok {
						//logger.LogWarn("TaskChannal 通道已关闭", logger.LOG_TERMINAL)
						end = false
					} else {
						run(target)
					}
				default:
				}
			}
		}()
	}
	tmps.Wait()
}

// 解析用户输入的“目标”字符串
func ParseUserInputTargetString(targetInput string) (targets []st_Target, err error) {
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
			return append(targets, st_Target{targetInput, true}), nil
		} else {
			// 检查是否为域名
			// 当目标主机以常见域名格式结尾时，认为是个域名
			parts := strings.Split(targetInput, ".")
			endDomain := "." + parts[len(parts)-1]
			index := sort.SearchStrings(common.DOMAINARRAY, endDomain)
			if index < len(common.DOMAINARRAY) && common.DOMAINARRAY[index] == "."+parts[len(parts)-1] {
				// 认为合法域名
				return append(targets, st_Target{targetInput, false}), nil
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

// 解析用户输入的“目标”字符串
func target2Task(targetInput string, cc uint) (err error) {
	var targetsBuffer []st_Target
	// 如果目标字符串中没有.(dot)认为目标无效
	if !strings.Contains(targetInput, ".") {
		err = errors.New("invalid target.1")
		return err
	}

	_, _, cidrError := net.ParseCIDR(targetInput)
	if cidrError != nil {
		// 尝试将目标解析为IP
		result := net.ParseIP(targetInput)
		if result != nil {
			// 成功将目标解析为IP，保存目标的IP
			targetsBuffer = append(targetsBuffer, st_Target{targetInput, true})
		} else {
			// 检查是否为域名
			// 当目标主机以常见域名格式结尾时，认为是个域名
			parts := strings.Split(targetInput, ".")
			endDomain := "." + parts[len(parts)-1]
			index := sort.SearchStrings(common.DOMAINARRAY, endDomain)
			if index < len(common.DOMAINARRAY) && common.DOMAINARRAY[index] == "."+parts[len(parts)-1] {
				// 认为合法域名
				targetsBuffer = append(targetsBuffer, st_Target{targetInput, false})
			} else if isIpRange(targetInput) {
				targetsBuffer = ipRange2Targets(targetInput)
			} else {
				err = errors.New("invalid target.")
				return err
			}
		}
	} else {
		// 是ip段,xxx.xxx.xxx.xxx/xx
		targetsBuffer = ipSeg2Targets(targetInput)
	}

	// 生成任务，将目标转化成 Task ，写入通道中。
	common.G_TaskChannal = make(chan *common.TASKUint, 10000) // 新建任务通道
	for _, target := range targetsBuffer {
		if target.IsIp {
			common.G_TaskChannal <- &common.TASKUint{Target: target.Target, TargetType: common.TASKUint_TargetType_IP, ControlCode: cc}
		} else {
			common.G_TaskChannal <- &common.TASKUint{Target: target.Target, TargetType: common.TASKUint_TargetType_Domain, ControlCode: cc}
		}
	}
	close(common.G_TaskChannal)
	return nil
}

// 将用户输入的IP段转成目标
func ipSeg2Targets(netw string) []st_Target {
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
	var targets []st_Target
	// loop through addresses as uint32.
	// I used "start + 1" and "finish - 1" to discard the network and broadcast addresses.
	for i := start + 1; i <= finish-1; i++ {
		// convert back to net.IPs
		// Create IP address of type net.IP. IPv4 is 4 bytes, IPv6 is 16 bytes.
		ip := make(net.IP, 4)
		binary.BigEndian.PutUint32(ip, i)
		targets = append(targets, st_Target{ip.String(), true})
	}
	// return a slice of strings containing IP addresses
	return targets
}

// 将用户输入的IP范围转成目标
// 匹配ip段，192.168.1.1-20，返回有效ip列表
func ipRange2Targets(s string) (targets []st_Target) {
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
			targets = append(targets, st_Target{v, true})
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

// 执行功能
func run(task *common.TASKUint) {
	// 嗅探器对象
	sniffer := asset_host.GetSnifferObj()

	var assetHost asset_host.AssetHost
	if task.TargetType&common.TASKUint_TargetType_IP == common.TASKUint_TargetType_IP {
		assetHost.IsIP = true
		assetHost.RealIP = task.Target
	} else {
		assetHost.IsIP = false
		assetHost.Domain = common.Domain{Name: task.Target}
	}

	var alived = true
	// ping扫功能
	if task.ControlCode&common.CC_PING_SCAN == common.CC_PING_SCAN {
		alived = sniffer.IsHostAlived(task.Target)
	}

	// 子域扫描功能
	if task.ControlCode&common.CC_SUB_DOMAIN_SCAN == common.CC_SUB_DOMAIN_SCAN {
		// 执行子域扫描功能
	}

	// 根据存活信息分别存放
	if alived {
		// 端口扫描功能
		if task.ControlCode&common.CC_PORT_SCAN == common.CC_PORT_SCAN {
			assetHost.AppendOpenedPortMap(sniffer.SnifferHostOpenedPorts(task.Target))
		}

		asset_host.GetSnifferObj().AppendAlivedAssetHost(assetHost)
	} else {
		asset_host.GetSnifferObj().AppendDiedAssetHost(assetHost)
	}
}
