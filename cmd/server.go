package main

import (
	"encoding/json"
	"fmt"
	"poodle/pkg/asset_host"
	"poodle/pkg/common"
	"poodle/pkg/controller"
	"poodle/pkg/logger"

	"github.com/liushuochen/gotable"
)

var G_ServerLogChannal chan string
var cp_srv *Kernel_SendOrderServer

type Server struct {
	UnimplementedKernelServer
}

func (s *Server) SendOrder(req *SendOrderRequest, srv Kernel_SendOrderServer) error {
	cp_srv = &srv
	common.G_LogInfoChannal = make(chan string, 1000)
	var err error
	var recvTaskPacket common.TaskPacket

	// 将收到的任务包json转成go结构体
	if e := json.Unmarshal([]byte(req.TerminalParamJson), &recvTaskPacket); e != nil {
		logger.LogError("接收到的数据转成json发生错误。", logger.LOG_TERMINAL)
	}

	// 多线程处理，开启2个子线程同时运行。
	// 线程1：解析终端参数结构体，生成对应的目标，将目标转化为任务结构体，传入通道中
	// 线程2：从通道中获取任务，开启-T指定的线程数并发处理任务

	// 3. 解析终端参数结构体，生成对应的目标，将目标转化为任务结构体，传入通道中
	logger.LogInfo("开始生成主机任务，请等待。", logger.LOG_TERMINAL)

	num, err := controller.GenrateTasks(&recvTaskPacket)
	if err != nil {
		// 解析用户的输入目标发生错误。
		logger.LogError(fmt.Sprintf("err: %v\n", err), logger.LOG_TERMINAL_FILE)
		return err
	}
	logger.LogInfo(fmt.Sprintf("共生成了%d个主机任务。", num), logger.LOG_TERMINAL)

	logger.LogInfo("开始执行主机任务。", logger.LOG_TERMINAL)
	controller.Run(recvTaskPacket)

	sendAssetInfo()
	for IsOpenLogInfoChannal {
		select {
		case logstring, ok := <-common.G_LogInfoChannal:
			if !ok {
				// logger.LogWarn("TaskChannal 通道已关闭", logger.LOG_TERMINAL)
				IsOpenLogInfoChannal = false
			} else {
				OrderServer2Client(logstring)
			}
		default:
		}
	}

	return err
}

var IsOpenLogInfoChannal = true

func CloseLogInfoChannal() {
	IsOpenLogInfoChannal = false
}
func OrderServer2Client(info string) {

	if cp_srv != nil {
		(*cp_srv).Send(&SendOrderReply{
			Info: info,
		})
		fmt.Println("string")
	}
}

func sendAssetInfo() {
	for _, v := range asset_host.GetSnifferObj().AlivedAssetHosts {
		var targ string
		var targAlived string
		if v.IsIP {
			targ = v.RealIP
		} else {
			targ = v.Domain.Name
		}

		if v.IsAlived {
			targAlived = "存活"
		} else {
			targAlived = "不存活"
		}

		var subDomain [][]string
		for _, v := range v.SubDomains {
			if v.IsAlived {
				subDomain = append(subDomain, []string{v.Name, "存活"})
			} else {
				subDomain = append(subDomain, []string{v.Name, "不存活"})
			}
		}

		tab, _ := gotable.Create("目标", "子域名", "存活状态", "开放端口", "服务信息", "爬虫结果")
		tab.AddRow([]string{targ, "", targAlived, "", "", ""})

		for _, v := range subDomain {
			tab.AddRow([]string{"", v[0], v[1], "", "", ""})
		}
		for k, v := range v.OpenedPorts {
			tab.AddRow([]string{"", "", "", k, v, ""})
		}
		fmt.Println(tab)
		OrderServer2Client(tab.String())
	}
}
