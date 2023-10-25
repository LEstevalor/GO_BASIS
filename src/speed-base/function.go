package main

import (
	"fmt"
)

func function2(a, b int, c string) int {
	fmt.Println(c)
	return a + b
}

func main2()  {
	x := "aaa"
	y := "bbb"
	x, y = swap(y, x)

	hhh(1, "hello", "world")
}
func hhh(h int, strs ...string)  {  // (arg ...type)  (参数组名 ...类型)
	for i := 0; i < len(strs); i++ {
		fmt.Printf("%s\t", strs[i])
	}
}
func swap(x, y string) (string,string) {
	return y, x
}
/**
func 函数名() 类型 {
	函数体
}
 */