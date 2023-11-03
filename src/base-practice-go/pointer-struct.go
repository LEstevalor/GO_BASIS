package practice

import "fmt"

// 节6 指针-结构体

type Vertex struct {
	X int
	Y int
}

var (
	aa = Vertex{}  // 默认
	bb = Vertex{X: 1}  // 默认Y
	cc = Vertex{Y: 2}  // 默认X
	dd = &aa
)

func main()  {
	var num *int
	fmt.Println(num)  // <nil>
	ui := 3
	num = &ui
	fmt.Println(ui, &ui, num, *num)  // 3 0xc00000e0c0 0xc00000e0c0 3
	*num++
	fmt.Println(ui, &ui, num, *num)  // 4 0xc00000e0c0 0xc00000e0c0 4


	structOne := Vertex{1, 2}
	structOne.X++
	fmt.Println(structOne)  // {2 2}

	fmt.Println(aa, bb, cc, dd)  // {0 0} {1 0} {0 2} &{0 0}

}
/*
指针
	Go 指针：保存一个值的内存地址
	类型 *T 是指向T的指针，默认零值为 nil

	& 操作符 生成一个指向其操作数的指针.
	* 操作符 表指针指向的底层值.

struct 是字段的集合

*/
