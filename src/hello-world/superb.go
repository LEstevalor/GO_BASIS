package main

import "fmt"

func main5() {
	c := high(1, 2, add)
	fmt.Println(c)   // 3

	// 匿名函数的骚操作
	c1 := high(1, 2, func(a, b int) int {
		return a + b
	})
	fmt.Println(c1)  // 3
}
// 高阶函数
func high(a, b int, fun func(int, int) int) int {
	return fun(a, b)
}

func add(a, b int) int {
	return a + b
}