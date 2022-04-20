package cmdinit

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

var (
	wg sync.WaitGroup
)

// 根据CMDPara.IsReadTargetsFromFile标志
// 获取目标列表，文件中读或者从用户输入中读
// 从文件中读取支持 TXT / XLAM / XLSM / XLSX / XLTM / XLTX 格式
// 用户输入每个目标空格分隔
// 需要进行合法性检测，确定目标是正确输入
// 最终结果是纯IP，或者纯域名(不包含协议，比如http)
// 另外全局并发控制从此处开始
func (CMDPara) GetTargets(CMDParas *CMDPara) {
	if CMDParas.IsReadTargetsFromFile {
		// todo
		// 从文件中解析，应该通过gorouting输入目标，并监听退出标志，

	} else {
		// 从手动设置的目标中解析，应该通过gorouting输入目标，并监听退出标志，
		targetSlice := strings.Split(CMDParas.UserInputTargetString, " ")
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, v := range targetSlice {
				for {
					select {
					case <-CMDParas.ExitFlag:
						fmt.Println("退出")
						os.Exit(0)
					default:
						CMDParas.Target <- v
					}
				}
			}
		}()
	}
	wg.Wait()
}
