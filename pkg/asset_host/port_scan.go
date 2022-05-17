// @author: greetdawn
// @date: 2022-05-10
// 该文件主要实现端口扫描功能

package asset_host

import (
	"net"
	"strconv"
	"sync"
	"time"
)

var (
	proto         string = "tcp"
	timeoutSecond int    = 5
	threads       int    = 200
)

func ScanHostOpenedPorts(target string) sync.Map {
	return _TCPOrUDPPortScan(target)
}

// 如果ports长度为0，则进行全端口扫描
// 默认200并发
func _TCPOrUDPPortScan(target string, ports ...string) sync.Map {
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
	} else {
		// 否则扫描全端口
		wg.Add(1)
		go func() {
			for i := 1; i <= 65535; i++ {
				productPortChan <- strconv.Itoa(i)
			}
			close(productPortChan)
			wg.Done()
		}()
	}

	// 并发扫描
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for p := range productPortChan {
				// fmt.Printf("\r[*]正在扫描:%s端口", p)
				// 拼接地址
				addr := net.JoinHostPort(target, p)
				// 建立连接
				conn, err := net.DialTimeout(proto, addr, time.Duration(timeoutSecond)*time.Second)
				// 如果连接正常建立
				if err == nil {
					var tmp [512]byte
					conn.SetReadDeadline(time.Now().Add(time.Second * 10))
					conn.Read(tmp[:])
					resMap.Store(p, string(tmp[:]))
				}
			}

		}()
	}
	wg.Wait()

	return resMap
}
