// 本路径下主要用于单元测试
// 单元测试命名规则：test_<需要测试的单元>.go

// 单元测试入口函数：
// func TestMain(m *testing.M) {}

// 单元测试子函数：
// func TestExample1(t *testing.T) {}

// 入口函数中调用以下方法可以调用全部子函数
// m.Run()

// 使用如下命令开始单元测试，-v参数可选
// go test <-v>

// BUG：单元测试无法使用flag包指定参数，应该修改默认参数

package test
