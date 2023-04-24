package main

import "fmt"

func main() {
	r1 := increment() // 此时r1就是返回的内部函数
	fmt.Println(r1)   // 0xa29780

	v1 := r1()
	fmt.Println(v1)   // 1

	v2 := r1()
	fmt.Println(v2)   // 2
}
// 自增函数
func increment() func() int {  // 相当于返回类型为func() int
	i := 0  // 外层函数局部变量，就算外部函数使用结束，由于内部函数调用（返回的就是内部函数），还是存在而不被销毁
	fun := func() int {
		i++
		return i
	}
	return fun   // 返回函数，类似python的装饰器
}