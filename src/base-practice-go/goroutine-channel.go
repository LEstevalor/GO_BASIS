package practice
// 节13 goroutine和channel
import (
	"fmt"
	"time"
)

func say(string2 string)  {
	for i := 0; i < 4; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(string2)
	}
	/**
		hello
		world
		world
		hello
		world
		hello
		hello
	 */
}

func sum(s []int, c chan int)  {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func fibonacci(n int, c chan int)  {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	close(c)
}

func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:  // 第一个case是c <- x，它表示将当前的斐波那契数x发送到通道c。然后，我们更新x和y的值，计算下一个斐波那契数
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main()  {
	/**
		Goroutine
			由Go运行时管理的轻量级线程（一种协程）
			由Go运行时系统进行管理和调度。与传统的线程相比，Goroutine具有更小的开销和更高的并发性能。
			使用Goroutine和通道（channel），Go语言提供了一种简洁而高效的并发编程模型，使得编写并发程序变得更加容易和安全。

	 */
	go say("world")
	say("hello")


	/**
		Channels
			通道是一个类型化管道，可通过它使用通道操作符 <- 发送和接收值
			默认情况下，发送和接收会阻塞，直到另一端准备就绪。这允许 goroutine 在没有显式锁或条件变量的情况下进行同步

		ch := make(chan int)  // 与map和切片一样，通道必须在使用前创建
		ch <- v    // 将 v 发送到通道 ch(数据沿箭头方向流动)
		v := <-ch  // 从 ch 接收，且给 v 赋值
	 */
	s := []int{1, 2, 3, 4}
	c := make(chan int)  // 无缓冲的通道（发送和接收操作同步，即发送操作和接收操作会相互阻塞，直到双方都准备好）
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x + y) /** 7 3 10 或 3 7 10
		由于通道c是无缓冲的，发送和接收操作是同步的。
		这意味着当第一个Goroutine执行完毕并发送sum值到通道c时，第一个接收操作<-c会解除阻塞并接收到这个值。
		然后，当第二个Goroutine执行完毕并发送sum值到通道c时，第二个接收操作<-c会解除阻塞并接收到这个值。
	 */

	// 通道可以被 缓冲，将缓冲长度作为第二个参数提供给 来初始化一个带缓冲的通道
	// 仅当通道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞.
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//ch <- 3  // 报错fatal error: all goroutines are asleep - deadlock! 溢出缓冲区
	fmt.Println(<-ch)  // 1
	fmt.Println(<-ch)  // 2

	/**
		range、close
			发送方可以 close通道 以指示不再发送任何值。接收方可通过为接收表达式分配第二个参数来测试通道是否已关闭 v, ok := <-ch（关闭为false）
			循环 会反复从通道接收值，直到它关闭。
			注：只有发送者应关闭channel，而不是接收者。在关闭的频道上发送会引起panic
				通道不像文件；常不需关闭它们。仅当必须告诉接收者没有更多值时才需要关闭，例如终止一个range循环
	*/
	cc := make(chan int, 10)
	go fibonacci(cap(cc), cc)
	for i := range cc {
		fmt.Println(i)
	}

	/**
	select
		select 语句使一个goroutine可以等待多个通信操作
		select 会阻塞，直到其中一个case可以行，然后执行该case。如果多个准备就绪，它会随机选择一个执行

	默认选择
		当 select中的其它分支都没有准备好时， default分支就会执行
		为了在尝试发送或者接收时不发生阻塞，可使用default分支
			select {
			case i := <-c:
				// 使用 i
			default:
				// 从 c 接收会阻塞
			}
	*/
	ccc := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ccc)  // 每次循环迭代，会阻塞在这个接收操作上，直到fibonacciSelect函数通过ccc <- x将新的斐波那契数发送到通道ccc
		}
		quit <- 0
	}()
	fibonacciSelect(ccc, quit)


	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}

}

