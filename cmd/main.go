package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"poodle/pkg/asset_host"
	cmdparser "poodle/pkg/cmd_parser"
	"poodle/pkg/common"
	controllor "poodle/pkg/controller"
	"poodle/pkg/logger"
	"poodle/pkg/pb"
	"sync"
	"time"

	"google.golang.org/grpc"
)

var mainWaitGroup sync.WaitGroup

var lis net.Listener

type server struct {
	pb.UnimplementedKernelServer
}

func (s *server) SendControlPkg(req *pb.ControlRequest, srv pb.Kernel_SendControlPkgServer) error {
	logger.SRV = &srv
	var err error
	err = json.Unmarshal([]byte(req.ParamKeys[0]), &cmdparser.G_TerminalParam)
	if err != nil {
		logger.LogError("接收到的数据转成json发生错误。", logger.LOG_TERMINAL)
	}
	// 多线程处理，开启2个子线程同时运行。
	// 线程1：解析终端参数结构体，生成对应的目标，将目标转化为任务结构体，传入通道中
	// 线程2：从通道中获取任务，开启-T指定的线程数并发处理任务
	mainWaitGroup.Add(2)
	go func() {
		defer mainWaitGroup.Done()
		// 3. 解析终端参数结构体，生成对应的目标，将目标转化为任务结构体，传入通道中
		err := cmdparser.GenrateTasks(&cmdparser.G_TerminalParam, uint(req.ControlCode))
		if err != nil {
			// 解析用户的输入目标发生错误。
			logger.LogError(fmt.Sprintf("err: %v\n", err), logger.LOG_TERMINAL_FILE)
			os.Exit(0)
			return
		}
	}()
	go func() {
		defer mainWaitGroup.Done()
		// 4. 从通道中获取任务，开启-T指定的线程数并发处理任务
		controllor.Run(cmdparser.G_TerminalParam.ThreadsNumber)
	}()

	mainWaitGroup.Wait()
	cmdparser.PrintAssetHostList(asset_host.GetSnifferObj().AlivedAssetHosts)
	for _, v := range asset_host.GetSnifferObj().AlivedAssetHosts {
		tmp, _ := json.Marshal(v)
		err = srv.Send(&pb.HelloReply{
			Asset: string(tmp),
		})
		time.Sleep(time.Duration(1) * time.Second)
	}
	asset_host.GetSnifferObj().AlivedAssetHosts = nil
	return err
}

func main() {
	// 输出banner信息
	common.ShowBanner()
	var err error
	// 创建监听器
	lis, err = net.Listen("tcp", fmt.Sprintf(":%d", 50041))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 创建服务器
	s := grpc.NewServer()

	// 注册服务器
	pb.RegisterKernelServer(s, &server{})
	logger.LogInfo(fmt.Sprintf("server listening at %v", lis.Addr()), logger.LOG_TERMINAL)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
