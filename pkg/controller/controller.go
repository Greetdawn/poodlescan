package controller

import (
	"fmt"
	"os"
	"poodle/pkg/asset_host"
	"poodle/pkg/common"
	"poodle/pkg/config"
	"poodle/pkg/logger"
	"sync"
)

// 多线程同步追加资产信息
var mutexOfAppendAsset sync.Mutex
var mutexOfAppendOpenedPorts sync.Mutex

//(alivedList asset_host.AssetHost, diedList asset_host.AssetHost, err error)
func Run(tp common.TaskPacket) {
	var tmps sync.WaitGroup
	tmps.Add(config.GConfig.ScanPortConfig.RunTaskThreads)
	for i := 0; i < config.GConfig.ScanPortConfig.RunTaskThreads; i++ {
		go func() {
			defer tmps.Done()
			var end = true
			for end {
				select {
				case task, ok := <-G_AssetTaskChannal:
					if !ok {
						// logger.LogWarn("TaskChannal 通道已关闭", logger.LOG_TERMINAL)
						end = false
					} else {
						run(task, tp)
					}
				default:
				}
			}
		}()
	}
	tmps.Wait()
}

// 执行功能
func run(task *TASKUint, tp common.TaskPacket) {
	// 嗅探器对象
	sniffer := asset_host.GetSnifferObj()
	// 资产主机
	var assetHost asset_host.AssetHost
	assetHost.Init()

	// 设置资产的初步信息，判断是不是IP还是域名，填入对应的字段中
	if task.TargetIsIP {
		assetHost.IsIP = true
		assetHost.RealIP = task.Target
	} else {
		assetHost.IsIP = false
		assetHost.Domain = common.Domain{Name: task.Target}
	}

	// ping扫功能
	if !tp.IsPn {
		assetHost.IsAlived = sniffer.IsHostAlived(task.Target)
		if assetHost.IsAlived {
			logger.LogInfo(fmt.Sprintf("%s 存活。", task.Target), logger.LOG_TERMINAL)
		} else {
			logger.LogWarn(fmt.Sprintf("%s 不存活。", task.Target), logger.LOG_TERMINAL)
			// 如果资产不存活，其他功能都不用执行
			mutexOfAppendAsset.Lock()
			asset_host.GetSnifferObj().AppendDiedAssetHost(assetHost)
			mutexOfAppendAsset.Unlock()
			return
		}
	}

	// 嗅探模块
	if tp.Sniffer {
		if err := doSniffer(task, tp, sniffer, &assetHost); err != nil {
			logger.LogError(err.Error(), logger.LOG_TERMINAL)
			os.Exit(0)
		}
	}

	// 扫描漏洞模块
	if tp.Vulscan {

	}
	// 多线程同步写入资产信息
	mutexOfAppendAsset.Lock()
	asset_host.GetSnifferObj().AppendAlivedAssetHost(assetHost)
	mutexOfAppendAsset.Unlock()
}

var onceParsePortRange sync.Once

func setPortsPool(task *TASKUint) {
	onceParsePortRange.Do(func() {
		var err error
		portStr := task.Params["ports"]
		asset_host.SCAN_PORT_PORTS, err = getPorts(portStr)
		if err != nil {
			logger.LogError(err.Error(), logger.LOG_TERMINAL_FILE)
		}
	})
}
func doSniffer(task *TASKUint, tp common.TaskPacket, sniffer *asset_host.Sniffer, asset *asset_host.AssetHost) error {
	var runSync sync.WaitGroup
	var err error
	// 1. 端口扫描功能
	if tp.PortScan {
		// 分析端口列表
		setPortsPool(task)

		// 执行端口扫描
		runSync.Add(1)
		go func() {
			defer runSync.Done()
			asset.AppendOpenedPortMap(sniffer.SnifferHostOpenedPorts(task.Target))
		}()
		runSync.Wait()
	}

	// 2. 子域探测
	if tp.SubDomain {
		logger.LogInfo("执行子域探测功能", logger.LOG_TERMINAL)
		// 分析端口列表
		setPortsPool(task)

		domains, err := sniffer.SniffSubDomain(task.Target)
		if err != nil {
			return err
		}
		asset.SubDomains = append(asset.SubDomains, domains...)
		if len(asset.SubDomains) == 0 {
			logger.LogWarn("Not found any subdomain.", logger.LOG_TERMINAL)
		}
		// ptr1 := &asset.SubDomains[0]
		// (*ptr1).IsAlived = true
		// 针对子域进行端口扫描和存活性探测
		for i := 0; i < len(asset.SubDomains); i++ {
			ptrOfSubDomain := &asset.SubDomains[i]
			(*ptrOfSubDomain).IsAlived = sniffer.IsHostAlived(ptrOfSubDomain.Name)
			(*ptrOfSubDomain).OpenPorts = asset_host.SyncMap2Map(sniffer.SnifferHostOpenedPorts(ptrOfSubDomain.Name))
		}
	}

	// 3. 爬虫
	if tp.Spider {
		// 执行爬虫功能
		logger.LogInfo("执行爬虫功能", logger.LOG_TERMINAL)
	}

	// 4. 指纹探测
	if tp.Fingerprint {
		// 执行指纹探测功能
		logger.LogInfo("执行指纹探测功能", logger.LOG_TERMINAL)
	}
	return err
}
