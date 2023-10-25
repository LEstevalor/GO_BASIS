package practice

// 节4 - 变量

import (
	"container/list"
	"fmt"
	"go/types"
	"hash"
)

// 节4 - 变量

var a, b bool
var live = "能不能把我的愿望还给我" // 和 var live string = "能不能把我的愿望还给我" 同意
const DROP = "但故事的最后你还是说了"

func main() {
	// 默认值
	var i int
	var u uint
	var k byte
	var r rune
	var e error
	var j float64
	var l complex64
	var name string
	var s list.List
	var t hash.Hash
	var y types.Array
	fmt.Println(a, b) // false false
	fmt.Println(i, u, k, r, e, j, l, name, s, t, y)
	// 0 0 0 0 <nil> 0 (0+0i)  {{<nil> <nil> <nil> <nil>} 0} <nil> {0 <nil>} 对应看看各个默认值

	// 变量声明
	var isNo = false // 同 var isNo bool = false
	fmt.Println(live, isNo)

	// 短变量声明（注：函数外无法使用短变量）
	quiet := "为什么天这么安静"
	numA, numB := 1, 2
	fmt.Printf("%s, %d, %d %T\n", quiet, numA, numB, quiet) // %T 会打印出类型
	// 当右值声明了类型时，新变量的类型与其相同:
	var sea int
	flower := sea                     // flower is an int
	love := 0.867 + 0.5i              // complex128
	fmt.Printf("%T %T", flower, love) // int complex128

	// 类型转换
	var ii int = 42
	var ff float64 = float64(ii)
	var uu uint = uint(ff)
	fmt.Println(ii, ff, uu) // 42 42 42

	//	常量，使用const关键字，不能使用短变量声明、
	const (
		Big    = 1 << 100
		RESULT = "拜拜"
	)
	//fmt.Println(Big)  // int 最多可以存储 64 位整数，打印会报错
	fmt.Println(DROP, RESULT)
}

/**
（知识点）
Basic types（基本类型），参考自 https://go.p2hp.com/tour/basics/11
	Go 的7个基本类型为：
	 布尔    bool
	 字节	byte	   (uint8的别名)
	 字符	rune       (int32的别名)
	 字符串  string
	 整型	int        (int8  int16  int32  int64)  uint (uint8 uint16 uint32 uint64 uintptr)
	 浮点	float32    (float64)
	 复数	complex64  (complex128)

注：int, uint, and uintptr 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽

默认零值是: 数字类型为 0， 布尔类型为 false， 字符串为 "" (空字符串）。

*/
