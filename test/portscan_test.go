package test

// func TestMain(m *testing.M) {

// 	// scanPort接收tcp/udp/ack/syn/fin，设置扫描类型
// 	// 指定范围扫：asset_host.TCPOrUDPPortScan("127.0.0.1", "tcp","1-65535")
// 	// 指定端口扫：asset_host.TCPOrUDPPortScan("127.0.0.1", "tcp","443,444,445")
// 	// 默认/精简端口扫：asset_host.TCPOrUDPPortScan("127.0.0.1", "tcp")
// 	asset_host.SCAN_PORT_PORTS = asset_host.SCAN_PORT_POODLE_COMMON_PORTS

// 	res := asset_host.ScanHostOpenedPorts("172.16.15.200")

// 	// 输出结果集
// 	fmt.Println()
// 	res.Range(func(key, value interface{}) bool {
// 		k := key.(int)
// 		v := value.([]string)
// 		fmt.Println(k, " ", v)
// 		return true
// 	})
// }
