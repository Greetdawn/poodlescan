package controllor

import (
	"poodle/pkg/common"
	"sync"
)

//(alivedList asset_host.AssetHost, diedList asset_host.AssetHost, err error)
func Run() {
	var tmps sync.WaitGroup
	tmps.Add(common.G_TerminalParam.ThreadsNumber)
	for i := 0; i < common.G_TerminalParam.ThreadsNumber; i++ {
		go func() {
			defer tmps.Done()
			var end = true
			for end {
				select {
				case target, ok := <-common.G_TaskChannal:
					if !ok {
						//logger.LogWarn("TaskChannal 通道已关闭", logger.LOG_TERMINAL)
						end = false
					} else {
						run(target)
					}
				default:

				}

			}
		}()
	}
	tmps.Wait()

}
func run(task *common.TASKUint) {
	// ping扫功能
	if task.ControlCode&common.CC_PING_SCAN == common.CC_PING_SCAN {
		common.IsHostAlived(task.Target)
	}
}
