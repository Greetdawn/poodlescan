/*
cmdParse的工作就只有一个：解析从客户端发来的命令。
根据命令的类型代码，把客户端的命令解析后，存放对应的参数到相应的对象中。
之后控制权交给对应的对象。
*/

package cmdParser

// cmdParse只对外暴露一个解析命令的方法。
// 这个方法会直接把解析出来的信息保存到对应的对象中。
// 返回对象特征码
// func ParseCmd(cmd string) int {
// 	return 1
// }
