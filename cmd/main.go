package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"poodle/pkg/asset_host"
	"poodle/pkg/common"
	controllor "poodle/pkg/controller"
	"poodle/pkg/logger"
	pb "poodle/pkg/mygrpc"
	"sort"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
)

var lis net.Listener

type server struct {
	pb.UnimplementedKernelServer
}

func (s *server) SendOrder(req *pb.SendOrderRequest, srv pb.Kernel_SendOrderServer) error {
	logger.SRV = &srv
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
	num, err := controllor.GenrateTasks(&recvTaskPacket)
	if err != nil {
		// 解析用户的输入目标发生错误。
		logger.LogError(fmt.Sprintf("err: %v\n", err), logger.LOG_TERMINAL_FILE)
		return err
	}
	logger.LogInfo(fmt.Sprintf("共生成了%d个主机任务。", num), logger.LOG_TERMINAL)

	logger.LogInfo("开始执行主机任务。", logger.LOG_TERMINAL)
	controllor.Run(recvTaskPacket)

	var syncs sync.WaitGroup
	syncs.Add(1)
	go func() {
		syncs.Done()
		prit()
		// for _, v := range asset_host.GetSnifferObj().AlivedAssetHosts {
		// 	tmp, _ := json.Marshal(v)
		// 	err = srv.Send(&pb.SendOrderReply{
		// 		Info: string(tmp),
		// 	})
		// 	time.Sleep(time.Duration(1) * time.Second)
		// }
		asset_host.GetSnifferObj().AlivedAssetHosts = nil
	}()
	syncs.Wait()
	return err
}

func main() {
	initSetKernel()

	var err error
	// 创建监听器
	lis, err = net.Listen("tcp", fmt.Sprintf(":%d", 50088))
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

func initSetKernel() {
	// 输出banner信息
	common.ShowBanner()
	fmt.Println()
	fmt.Print("正在启动内核，请稍后")
	sort.Strings(common.DOMAINARRAY)
	for i := 0; i < 10; i++ {
		fmt.Print(".")
		time.Sleep(time.Duration(200) * time.Millisecond)
	}
	fmt.Println()
	fmt.Println("现在开始配置内核，请选择配置策略：")
	fmt.Println("1. 使用默认配置")
	fmt.Println("2. 自定义配置")
	var id int
	fmt.Scanf("%d", &id)
	if id == 2 {
		fmt.Println("自定义内核配置项：[1/2]")
		fmt.Print("运行线程数： ")
		fmt.Scanf("%d", &common.G_RunTaskThreads)
		fmt.Println("自定义内核配置项：[2/2]")
		fmt.Print("是否开启内核日志?(yes/no)：  ")
		logger.IsPrintLogInfo = scanYesOrNo()
	}
	fmt.Print("正在完成配置，请稍后")
	for i := 0; i < 5; i++ {
		fmt.Print(".")
		time.Sleep(time.Duration(200) * time.Microsecond)
	}
	fmt.Println()
	fmt.Print(" *******************  内核初始化完成  *******************\n")
}

// 输入 yes 或者y 视为true。不区分大小写
func scanYesOrNo() bool {
	var s string
	fmt.Scanf("%s", &s)
	//s = strings.ReplaceAll(s, "\r", "")
	//s = strings.ReplaceAll(s, "\n", "")
	if strings.ToLower(s) == "yes" || strings.ToLower(s) == "y" {
		return true
	}
	return false
}

func prit() {

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
		for i, v := range v.SubDomains {
			subDomain[i][0] = v.Name
			if v.IsAlived {
				subDomain[i][1] = "存活"
			} else {
				subDomain[i][1] = "不存活"
			}
		}
		//
		logger.PrintAssetHostList(targ, targAlived, subDomain, &v.OpenedPorts)
	}
}
