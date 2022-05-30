package mygrpc

import (
	"encoding/json"
	"fmt"
	"poodle/pkg/asset_host"
	"poodle/pkg/common"
	"poodle/pkg/controller"
	"poodle/pkg/logger"
)

var cp_srv *Kernel_SendOrderServer
var serverLogChannal chan string
var GPtr_ServerLogChannal *chan string

type Server struct {
	UnimplementedKernelServer
}

func (s *Server) SendOrder(req *SendOrderRequest, srv Kernel_SendOrderServer) error {
	cp_srv = &srv
	serverLogChannal = make(chan string, 1000)
	GPtr_ServerLogChannal = &serverLogChannal
	logger.GPtr_LogModuleInfoChannal = GPtr_ServerLogChannal
	var err error
	var recvTaskPacket common.TaskPacket

	// 将收到的任务包json转成go结构体
	if e := json.Unmarshal([]byte(req.TerminalParamJson), &recvTaskPacket); e != nil {
		//logger.LogError("接收到的数据转成json发生错误。", logger.LOG_TERMINAL)
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

	go func() {
		var end = true
		for end {
			select {
			case log, ok := <-serverLogChannal:
				if !ok {
					// logger.LogWarn("TaskChannal 通道已关闭", logger.LOG_TERMINAL)
					end = false
				} else {
					OrderServer2Client(log)
				}
			default:
			}
		}
	}()

	sendAssetInfo()
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
	}
}

func sendAssetInfo() {
	for _, v := range asset_host.Assets2Strings(true) {
		OrderServer2Client(v)
	}
}
