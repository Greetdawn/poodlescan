package cmdinit

import "strings"

// 根据CMDPara.IsReadTargetsFromFile标志
// 获取目标列表，文件中读或者从用户输入中读
// 从文件中读取支持 TXT / XLAM / XLSM / XLSX / XLTM / XLTX 格式
// 用户输入每个目标空格分隔
// 需要进行合法性检测，确定目标是正确输入
// 最终结果是纯IP，或者纯域名(不包含协议，比如http)
// 另外全局并发控制从此处开始
func (CMDPara) GetTargets(CMDParas *CMDPara) {
	if CMDParas.IsReadTargetsFromFile {
		// 从文件中解析，应该通过

	} else {
		targetSlice := strings.Split(CMDParas.UserInputTargetString, " ")
		go func() {
			for _, v := range targetSlice {
				if len(v) > 0 {
					CMDParas.Target <- v
				}
			}
		}()
	}
}
