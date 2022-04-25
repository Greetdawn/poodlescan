// 主要进行信息收集，最终结果存放在PoodleSnif类型结构体中

// @koi 本包进行资产的嗅探
package sniffer

// @koi
// sniffer.go
// 声明 sniffer 的接口声明
// 声明 sniffer struct 。 所有嗅探模块的父类，对外只暴露父类。

// @koi
// sniffer_domain
// 嗅探器的子类。
// 如果是探测域名，则使用域名探测器

// @koi
// sniffer_ip
// 嗅探器的子类。
// 如果是探测IP，则使用IP探测器
