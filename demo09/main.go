package main

import (
	"demo09/myinit"
	"fmt"
)

func main() {
	fmt.Println(">>> main 函数开始执行")
	// 我们并没有调用 myinit.init()，它已经自动执行了
	fmt.Println("myinit.AB1 的值是:", myinit.AB1)
	fmt.Println("myinit.AB 的值是:", myinit.AB)

}
