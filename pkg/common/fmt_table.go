package common

import (
	"github.com/liushuochen/gotable"
	"github.com/liushuochen/gotable/table"
)

var CmdTable *table.Table

func CreateTable() {
	if CmdTable == nil {
		CmdTable, _ = gotable.Create("主机IP", "存活性", "开放端口")
	}
}

func AddAssetHostRecord(ip string, isAlived bool, openPorts string) {
	var alivedStr string
	if isAlived {
		alivedStr = "存活"
	} else {
		alivedStr = "不存活"
	}
	CmdTable.AddRow([]string{ip, alivedStr, openPorts})
}
