package practice
// 节11 错误

import (
	"fmt"
	"strconv"
	"time"
)


type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%v, %s, %v\n", e.When, e.What, e.What)  // 2023-11-04 22:15:30.7155489 +0800 CST m=+0.029298901, it didn't work, it didn't work
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main()  {
	/**
		Go用 error值表错误状态
		error 类型是类似于 fmt.Stringer 的内置接口
			type error interface {
				Error() string
			}

	注意: 在 Error 方法内调用 fmt.Sprint(e) 会让程序陷入死循环。您可以通过首先转换 e 来避免这种情况，
		  比如当给定负数时，Sqrt 应返回非nil错误值，因为它不支持复数，可以打印fmt.Sprint(float64(e))。为什么？

		  因Error()方法内调用了fmt.Sprint(e)，而fmt.Sprint()会调用Error()方法来获取错误的字符串表示。
		  循环调用导致死循环。为避免，可先将e转换为其他类型，如float64类型（就不走Error方法了），然后再用fmt.Sprint()获取字符串表示
	 */
	if errTest := run(); errTest != nil {
		fmt.Println(errTest)  // 与 fmt.Stringer 一样，fmt 包在打印值时会查找 error 接口
	}

	i, err := strconv.Atoi("42")  // 一个nil error 表示成功;非 nil error 表示失败（通过测试错误是否等于 nil 来处理错误）
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}
	fmt.Println("Converted integer:", i)  // Converted integer: 42


	/**
		Go语言中的错误处理机制是通过返回错误值来处理异常情况，而不是使用异常机制。

		注意：recover函数只能在defer函数中使用。如果在没有发生panic的情况下调用recover函数，它将返回nil
			Go语言没有像其他语言那样的异常捕获机制，但可以使用defer和recover来实现类似的效果，以处理panic引发的错误

		Go语言中，panic用于引发异常（类似Exception），异常后开始执行调用栈中的延迟函数（defer）
	 */
	defer func() {
		// defer关键字定义了一个匿名函数，用内置recover函数捕获panic引发的错误。如果发生了panic错误，recover函数将返回错误信息
		if err := recover(); err != nil {
			fmt.Println("捕获到panic错误:", err)  // 捕获到panic错误: 发生了一个错误
		}
	}()

	// 引发panic错误
	panic("发生了一个错误")
}
