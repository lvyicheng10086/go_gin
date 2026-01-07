package myinit

import "fmt"

var AB1 int

// init 是一个特殊的函数
// 1. 不需要传入参数，也没有返回值
// 2. 不能被显式调用（myinit.init() 是非法的）
// 3. 在包被导入时自动执行
func init() {
	fmt.Println(">>> myinit 包的 init() 函数正在执行...")
	AB1 = 234 // 自动完成赋值
}
