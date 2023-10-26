package practice

// 节7 数组-切片

import (
	"fmt"
	"strings"
)

func main()  {
	/*
		数组：类型写法 [n]T
	*/
	var arrayOne [2]string
	arrayOne[0] = "想念还在等待救援"
	arrayOne[1] = "天空灰得想哭过"

	arrayTwo := [3]int{1, 2}
	arrayThree := [2]string{"信誓旦旦给了承诺", "却被时间扑了空"}  // 注：短变量里要注定类型

	fmt.Println(arrayOne)
	fmt.Println(arrayTwo, arrayThree)  // [1 2 0] [信誓旦旦给了承诺 却被时间扑了空]


	/*
		-------- 切片 ---------
	*/
	// （1）有指定值，可以不指定长度 -> 切片
	arrayFour := []string{"试着让故事继续", "分散时间的注意", "不知来不来得及"}
	fmt.Println(arrayFour)

	// （2）切片 - 切片表达式
	fmt.Println(arrayTwo[1])  // 2
	fmt.Println(arrayTwo[:1])  // [1]  切片表达式
	fmt.Println(arrayTwo[:2])  // [1 2]
	fmt.Println(arrayTwo[0:])  // [1 2 0]
	fmt.Println(arrayTwo[1:3])  // [2 0] 切片  切片表达式 左闭右开

	// （3）切片零值
	var arrayNil []int
	fmt.Println(arrayNil, len(arrayNil), cap(arrayNil))  // [] 0 0
	if arrayNil == nil {
		fmt.Println("nil!")  // nil!
	}

	// （4）make
	//   切片创建（动态创建数组大小的方式）
	//   make源码中第二位就是size
	arrayA := make([]int, 3)
	arrayB := make([]int, 0, 3)
	showLenAndCap(arrayA)  // array:[0 0 0], len: 3，cap：3
	showLenAndCap(arrayB)  // array:[], len: 0，cap：3

	// （5）len 和 cap
	//	 数组的 len 和 cap 相等，因数组的长度和容量是固定的。当你创建一个数组时，它的长度和容量就已确定了
	//   切片（slice）len 和 cap 可能不同，创建一切片时，len 和 cap 可能不同，因切片是一动态的数据结构，长度和容量可在运行时改变
	showLenAndCapInt(arrayTwo)  // len: 3，cap：3
	showLenAndCap(arrayTwo[1:3])  // len: 2，cap：2

	/*
		（6）切片可以包含任何类型，包括其他切片（切片的切片）
			X _ _
			_ _ O
	*/
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[1][2] = "O"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// （7）append
	//    切片追加元素
	var arrayAppend []int
	arrayAppend = append(arrayAppend, 1, 2, 3)  // [1 2 3]

	fmt.Println(arrayAppend)
	fmt.Println(append(arrayAppend, 4, 5))  // [1 2 3 4 5]

	/* （8）range
	     for 循环的 range 形式可遍历切片或映射.

		0 9
		1 8
		2 7
		3 6
		4 5

		for i, _ := range arrayRange       // 只要索引
		for _, value := range arrayRange   // 只有值
	 */
	var arrayRange = []int{9, 8, 7, 6, 5}

	for i, v := range arrayRange {
		// 每次迭代都会返回两个值。第一个是索引，第二个是该索引中元素的副本
		fmt.Println(i, v)
	}


}

// 注意参数[n]的n是确认的，比如声明好大小的数组
func showLenAndCapInt(arrays [3]int)  {
	// 获取长度和容量  len & cap
	fmt.Printf("array:%v, len: %d，cap：%d\n", arrays, len(arrays), cap(arrays))
}

// 注意参数[n]的n是没确认的，比如切片（包括未声明大小的数组）
func showLenAndCap(arrays []int)  {
	// 获取长度和容量  len & cap
	fmt.Printf("array:%v, len: %d，cap：%d\n", arrays, len(arrays), cap(arrays))
}

// 获取长度和容量  len & cap
func showLenAndCapAnyNum(data interface{}) {
	switch v := data.(type) {
	case []int:
		fmt.Printf("array:%v, len: %d，cap：%d\n", v, len(v), cap(v))
	case [3]int:
		fmt.Printf("array:%v, len: %d，cap：%d\n", v, len(v), cap(v))
	default:
		fmt.Println("unsupported data type")
	}
}