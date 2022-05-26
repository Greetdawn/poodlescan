package logger

import (
	"fmt"
	pb "poodle/pkg/mygrpc"

	"github.com/liushuochen/gotable"
)

// 打印所有的资产信息
func PrintAssetHostList(target, isAlived string, subDomain [][]string, ports *map[string]string) {

	tab, _ := gotable.Create("目标", "子域名", "存活状态", "开放端口", "服务信息", "爬虫结果")
	tab.AddRow([]string{target, "", isAlived, "", "", ""})

	for _, v := range subDomain {
		tab.AddRow([]string{"", v[0], v[1], "", "", ""})
	}
	for k, v := range *ports {
		tab.AddRow([]string{"", "", "", k, v, ""})
	}
	if SRV != nil {
		_ = (*SRV).Send(&pb.SendOrderReply{
			Info: tab.String(),
		})
	}
	fmt.Println(tab)
}

// common.CloseDB()

//******** 处理不存活情况  *****************
// for _, asset := range this.DiedAssetHosts {
// 	var first bool = true
// 	tab, _ := gotable.Create("主机IP", "存活性", "开放端口", "服务信息")

// 	if len(asset.OpenedPorts) == 0 {
// 		tab.AddRow([]string{asset.RealIP, "不存活", "", ""})
// 	} else {
// 		for key, value := range asset.OpenedPorts {
// 			if first {
// 				tab.AddRow([]string{asset.RealIP, "不存活", key, value})
// 				first = false
// 			} else {
// 				tab.AddRow([]string{" ", " ", key, value})
// 			}
// 		}
// 	}
// 	fmt.Println(tab)
// }
//******** 处理不存活情况  end *****************
