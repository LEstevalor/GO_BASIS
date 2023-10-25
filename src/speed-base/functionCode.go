package main

import "fmt"

func main4() {
	func() {
		fmt.Println("我是一个匿名函数")
	}()  // 最后的括号记得加——表示调用（运行就会调用这个匿名函数）

	// 最后括号不加的语法就得赋值给函数变量
	f1 := func(a, b int) {
		fmt.Printf("我也是匿名函数%d\n", a + b)
	}
	f1(1, 2)

	// 带返回值
	f2 := func(a, b int) int {
		return a + b
	}
	a := f2(1, 3)
	fmt.Println(a)
}
