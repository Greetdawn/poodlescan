package cmdparser

import (
	"fmt"
	"poodle/pkg/asset_host"
	"poodle/pkg/common"

	"github.com/liushuochen/gotable"
	"github.com/liushuochen/gotable/table"
)

var isFirst bool = true

// 打印所有的资产信息
func PrintAssetHostList(assets []asset_host.AssetHost) {
	for _, asset := range assets {
		tab, _ := gotable.Create("目标", "子域名", "存活状态", "开放端口", "服务信息", "爬虫结果")

		isFirst = true
		var targetStr string
		if asset.IsIP {
			targetStr = asset.RealIP
		} else {
			targetStr = asset.Domain.Name
		}
		// 增加端口信息
		appendTableOfPorts(tab, &asset.OpenedPorts, targetStr, asset.IsAlived)
		//增加子域信息
		appendTableOfSubdomains(tab, asset.SubDomains, targetStr, asset.IsAlived)

		fmt.Println(tab)
	}
}

// 打印端口
func appendTableOfPorts(table *table.Table, ports *map[string]string, target string, isAlived bool) {
	for k, v := range *ports {
		if isFirst {
			if isAlived {
				table.AddRow([]string{target, "", "存活", k, v, ""})
			} else {
				table.AddRow([]string{"", "", "不存活", k, v, ""})
			}
			isFirst = false
		} else {
			table.AddRow([]string{"", "", "", k, v, ""})
		}
	}
}

// 打印子域信息
func appendTableOfSubdomains(table *table.Table, subdomains []common.Domain, target string, isAlived bool) {
	for _, v := range subdomains {
		if isFirst {
			if isAlived {
				appendTableOfPorts(table, &v.OpenPorts, target, isAlived)
				table.AddRow([]string{target, v.Name, "存活", "", "", ""})
			} else {
				appendTableOfPorts(table, &v.OpenPorts, target, isAlived)
				table.AddRow([]string{"", v.Name, "不存活", "", "", ""})
			}
			isFirst = false
		} else {
			appendTableOfPorts(table, &v.OpenPorts, target, isAlived)
			table.AddRow([]string{"", v.Name, "", "", "", ""})
		}
	}
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
