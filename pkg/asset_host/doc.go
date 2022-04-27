// 资产分类模块，主要根据扫描结果识别主机资产
package asset_host

// asset_host.go
// 定义 资产主机
// 嗅探的结果保存到资产主机对象中

// domain.go
// 定义 域名

// Sniffer 嗅探器
//
// sniffer.go
// 声明 sniffer 的接口声明
// 声明 sniffer struct 。 所有嗅探模块的父类，对外只暴露父类。
//
// sniffer_domain
// 嗅探器的子类。
// 如果是探测域名，则使用域名探测器
//
// sniffer_ip
// 嗅探器的子类。
// 如果是探测IP，则使用IP探测器
