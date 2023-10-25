package practice

// 节2 - 包

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(5)) // My favorite number is 2

	fmt.Println("My favorite float is", math.Sqrt(2)) // My favorite float is 1.4142135623730951

	fmt.Println("My favorite circle number is", math.Pi) // My favorite circle number is 3.141592653589793
}

/*
package main：
	对于Go语言的可执行程序，必须从package main的main函数开始运行。main函数是Go程序的入口点，它是程序的起始位置，定义了程序的执行逻辑（且只能有一个main入口）。
	Go语言中，一个可执行程序必须包含一个package main，并且在该包中必须包含一个main函数。main函数的签名必须是func main()，没有参数和返回值。当程序运行时，Go编译器会自动查找并执行main函数。
	除main函数外，其他函数可在程序中定义和使用，但它们只能被其他函数调用，而不能作为程序的入口点。只有main函数才能作为程序的入口点，从而启动整个程序的执行。

package 命名使用：
	按约定，包名与导入路径的最后一个元素相同。例如，"math/rand" 包中的源码均以`package`rand` 语句开始，如这里的 rand.Intn
*/
