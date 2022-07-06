package main

import (
	"fmt"
	"log"
	"net"
	"poodle/pkg/common"
	"poodle/pkg/config"
	"poodle/pkg/logger"
	"poodle/pkg/mygrpc"
	"poodle/pkg/os"

	"sort"
	"strings"
	"time"

	"google.golang.org/grpc"

	"github.com/fatih/color"
)

func main() {
	// 检查内核运行环境
	kernelEnvCheck()

	// 初始化内核
	initKernel()

	// 创建监听器
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50088))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 创建服务器
	s := grpc.NewServer()

	// 注册服务器
	mygrpc.RegisterKernelServer(s, &mygrpc.Server{})
	logger.LogInfo(fmt.Sprintf("server listening at %v", lis.Addr()), logger.LOG_TERMINAL)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initKernel() {
	fmt.Print("正在启动内核，请稍后")
	// 字符串数组排序
	sort.Strings(common.DOMAINARRAY)
	// 创建配置对象
	config.InitConfigModule()

	for i := 0; i < 10; i++ {
		fmt.Print(".")
		time.Sleep(time.Duration(50) * time.Millisecond)
	}
	fmt.Println()
	fmt.Println("现在开始配置内核，请选择配置策略：")
	fmt.Println("1. 使用默认配置")
	fmt.Println("2. 自定义配置")

	// only debug
	var id int = 1
	// fmt.Scanln(&id)
	if id == 2 {
		fmt.Println("自定义内核配置项：[1/2]")
		fmt.Print("运行线程数： ")
		fmt.Scanln(&config.GConfig.ScanPortConfig.RunTaskThreads)
		fmt.Println("自定义内核配置项：[2/2]")
		fmt.Print("是否开启内核日志?(yes/no)：  ")
		config.GConfig.IsPrintLogInfo = scanYesOrNo()
	}

	fmt.Print("正在完成配置，请稍后")
	for i := 0; i < 5; i++ {
		fmt.Print(".")
		time.Sleep(time.Duration(50) * time.Microsecond)
	}
	fmt.Println()
	fmt.Print(" *******************  内核初始化完成   *******************\n")
	fmt.Print(" *******************     Running     *******************\n")
}

func kernelEnvCheck() {
	fmt.Println("运行环境检查......")
	if os.IsHavaNMap() {
		color.Green("[+] namp")
	} else {
		os.InstallNmap()
		color.Yellow("[x] namp")
	}
}

// 输入 yes 或者y 视为true。不区分大小写
func scanYesOrNo() bool {
	var s string
	fmt.Scanf("%s", &s)
	if strings.ToLower(s) == "yes" || strings.ToLower(s) == "y" {
		return true
	}
	return false
}
