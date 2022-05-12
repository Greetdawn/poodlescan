package logic

import (
	cmdparser "poodle/internal/cmd_parser"
	"sync"
)

var WG sync.WaitGroup

func PoodleLogic(c *cmdparser.CMDPara, isOrder bool, extFunc ...func(cmdparser.TargetInput)) {
	// 生成目标
	c.TargetChan = make(chan cmdparser.TargetInput)
	WG.Add(1)
	go func() {
		c.ProduceTargets()
		close(c.TargetChan)
		WG.Done()
	}()

	// 全局并发控制
	for i := 0; i <= c.Threads; i++ {

		WG.Add(1)
		go func() {
			defer WG.Done()
			for ipOrDomainTarget := range c.TargetChan {
				var childwg sync.WaitGroup

				// 如果顺序执行
				if isOrder {
					childwg.Add(1)
					go func(targetTest cmdparser.TargetInput) {
						for _, f := range extFunc {
							f(targetTest)
						}
						childwg.Done()
					}(ipOrDomainTarget)
					childwg.Wait()
				} else {

					// 如果并发执行
					for _, f := range extFunc {
						childwg.Add(1)
						go func(ipOrDomainTarget cmdparser.TargetInput, f func(cmdparser.TargetInput)) {
							f(ipOrDomainTarget)
							childwg.Done()
						}(ipOrDomainTarget, f)
						childwg.Wait()
					}
				}

			}
		}()
	}

	WG.Wait()
}
