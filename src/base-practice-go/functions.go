package practice

// 节3 - 函数

import (
	"fmt"
	"math"
)

func funOneAdd(x int, y int) int {
	return x + y
}

func funTwoMutiply(x, y int) int {
	return x * y
}

func funThreeSwap(x, y string) (string, string) {
	return y, x
}

func funFourAbs(sum int) (x, y int) {
	x = int(math.Abs(float64(sum)))
	y = -x
	return
}

func compute(fn func(float64, float64) float64) float64{
	return fn(3, 4)
}

func main() {
	/**
		函数：
			这里 funOneAdd 的类型定义，有点类型python中的 x:int, y:int -> int
			但有优化写法（比如 funTwoMutiply 的写法）：
				当两个或多个连续命名函数参数共享一个类型时，除了最后一个之外，您可以省略所有类型.
			这里的多类型返回也挺有意思，比如 funThreeSwap 这里的返回两个 (string, string)，这里跟python是一致的

			甚至返回值是可以用变量命名的，有点类似MATLAB中的写法，例如 funFourAbs。
			注：不带参数的 return 语句返回命名的返回值，被称“裸”返回。应只在短函数使用，会损害较长函数的可读性
	*/
	fmt.Println("This is functions")
	fmt.Println("Hello，9", funOneAdd(4, 5))     // Hello，9 9
	fmt.Println("Hello，9", funTwoMutiply(3, 3)) // Hello，9 9
	fmt.Println(funThreeSwap("9 9", "Hello，"))  // Hello， 9 9 （这里打印两个变量Println默认插入两个空格）
	fmt.Println(funFourAbs(-9))                 // 9 -9

	/**
		函数也是值。它们可以像其他值一样传递.
	 	函数值可以用作函数参数和返回值.
	 */
	fnVariable := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(fnVariable)  // 0x6d99e0
	fmt.Println(fnVariable(3, 5))  // 5.830951894845301
	fmt.Println(compute(fnVariable))  // 5

	/**
		函数闭包
			闭包是引用其主体外部的变量的函数值。该函数可以访问并分配给引用的变量

		以下例子结果：
			0 0
			1 -2
			3 -6
	 */
	pos, neg := adder(), adder()
	for i := 0; i < 3; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

}

// adder 函数返回一个闭包。每个闭包都绑定到它自己的 sum 变量.
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
