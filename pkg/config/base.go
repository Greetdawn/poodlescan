package config

import "strings"

// global variation
var GConfig StConfig

// 配置文件的文件名
const (
	file_name_config_json string = "config.json"
)

// 配置项
const (
	runTaskThreads string = "runTaskThreads"
	isPrintLogInfo string = "isPrintLogInfo"
)

// 配置命令的包
type ConfigPacket struct {
	Key   string // 对应的配置项
	Value string
}

// write struct into config file
func (c *ConfigPacket) UpdateConfig() (err error) {
	switch c.Key {
	case runTaskThreads:
		//KnelConfig.RunTaskThreads, err = strconv.Atoi(c.Value)
	case isPrintLogInfo:
		//KnelConfig.IsPrintLogInfo, err = strconv.ParseBool(c.Value)
	}
	return err
}

type StSubDomainConfig struct {
	// choose which method to use for subdomain scanning.
	// fofa
	Use        string
	FOFAConfig struct {
		Email string
		Key   string
	}
}

type StConfig struct {
	// whether to print the immediately
	IsPrintLogInfo bool `json:"isPrintLogInfo"`
	// the config of scan config
	ScanPortConfig struct {
		// the number of threads running the task
		RunTaskThreads int `json:"threads"`
		// default scan ports
		DefaultScan1000Ports string `json:"defaultScan1000Ports"`
		// the proto of scan ports
		// 指定扫描端口时使用的协议。可选项有：tcp、udp、syn、ack、fin
		Scan_Port_Proto string `json:"scanportproto"`
	}
	ScanDomainConfig struct {
		// 1. fofa
		// 2. subfinder
		Using      string `json:"scanapi"`
		FOFAConfig struct {
			Email string `json:"fafa_email"`
			Key   string `json:"fafa_key"`
		}
	}
}

// get default scan ports
func (cfg *StConfig) GetDefaultScanPorts() []string {
	port := strings.ReplaceAll(cfg.ScanPortConfig.DefaultScan1000Ports, "\n", "")
	return strings.Split(port, " ")
}
