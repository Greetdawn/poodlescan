// @author: greetdawn
// @date: 2022-05-10
// 该文件主要实现主机存活检测基本方法

package common

import (
	"bytes"
	"encoding/binary"
	"net"
	"os/exec"
	"poodle/pkg/logger"
	"runtime"
	"strings"
	"time"
)

// 判断主机是否存活
// host可以传入IP、域名
func IsHostAlived(host string) bool {
	// 通过ICMP方法判断当前主机是否存活
	if Icmp(host) {
		logger.LogInfo("JudgeHost", "目标主机: "+host+" 存活")
		return true
	} else {
		// 使用系统本机ping判断当前主机是否存活
		if Ping(host) {
			logger.LogInfo("JudgeHost", "目标主机: "+host+" 存活")
			return true
		} else {
			logger.LogWarning("JudgeHost", "目标主机: "+host+" 未存活")
			return false
		}
	}
}

var (
	BufferByteMax           = 65535
	timeout       int64     = 1200 //1200毫秒
	command       *exec.Cmd        //命令执行
)

// 构造ICMP数据包格式
type ICMP struct {
	Type        uint8
	Code        uint8
	Checksum    uint16
	Identifier  uint16
	SequenceNum uint16
}

func CheckSum(data []byte) (rt uint16) {
	var (
		sum    uint32
		length = len(data)
		index  int
	)
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[index]) << 8
	}
	rt = uint16(sum) + uint16(sum>>16)
	return ^rt
}

// 定义icmp探测功能
func Icmp(host string) bool {
	SuccessTimes := 0
	conn, _ := net.DialTimeout("ip:icmp", host, time.Duration(timeout)*time.Millisecond)
	if conn == nil {
		return false
	}
	// 执行完毕后将连接关闭
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			logger.LogError("Icmp", err.Error())
		}
	}(conn)

	// 设施icmp头部信息
	icmp := ICMP{8, 0, 0, 0, 0}

	var buffer bytes.Buffer
	_ = binary.Write(&buffer, binary.BigEndian, icmp)

	data := make([]byte, 1)
	buffer.Write(data)
	data = buffer.Bytes()

	for i := 1; i < 11; i++ {
		icmp.SequenceNum = uint16(1) //检验和设为0
		data[2], data[3] = byte(0), byte(0)

		data[6], data[7] = byte(icmp.SequenceNum>>8), byte(icmp.SequenceNum)

		icmp.Checksum = CheckSum(data)
		data[2], data[3] = byte(icmp.Checksum>>8), byte(icmp.Checksum)

		tmpTimeNow := time.Now()
		_ = conn.SetReadDeadline(tmpTimeNow.Add(time.Duration(time.Duration(timeout) * time.Millisecond)))

		_, err := conn.Write(data)
		if err != nil {
			logger.LogError("Icmp", err.Error())
		}

		buf := make([]byte, BufferByteMax)
		n, err := conn.Read(buf)
		if err != nil {
			continue
		}

		SuccessTimes++
		if n > 0 {
		}
	}

	if SuccessTimes > 0 {
		return true
	} else {
		return false
	}
}

// 调用系统ping命令实现主机存活探测
func Ping(host string) bool {
	// 运行前判断当前系统类型
	OS := runtime.GOOS
	if OS == "linux" || OS == "darwin" {
		command = exec.Command("/bin/bash", "-c", "ping -c 1 -w 1 "+host+" >/dev/null && echo true || echo false")
	} else if OS == "windows" {
		command = exec.Command("cmd", "/c", "ping -n 1 -w 100 "+host+" && echo true || echo false")
	}
	outinfo := bytes.Buffer{}
	command.Stdout = &outinfo
	// 非阻塞性执行命令
	err := command.Start()
	if err != nil {
		return false
	}
	if err = command.Wait(); err != nil {
		return false
	} else {
		if strings.Contains(outinfo.String(), "true") {
			return true
		} else {
			return false
		}
	}
}
