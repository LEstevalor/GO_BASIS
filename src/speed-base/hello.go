package main

import "fmt"

// 导入一个系统包fmt用于输出

func main1() {
	/*
	var num int
	num = 100
	fmt.Printf("num:%d，内存地址：%p ", num, &num) //打印

	num = 200
	fmt.Printf("num:%d，内存地址：%p ", num, &num) //打印
	// num:100，内存地址：0xc00000e0a8 num:200，内存地址：0xc00000e0a8
	// 发现结果是同一内存地址，相当于num变量的地址是取定的，只是改变值而已，相当于python中的可变类型

	var str string
	str = "Hello"
	fmt.Printf("%T, %s\n", str, str)  // string, Hello

	// 单引号
	v1 := 'A'  // 编码表 ASCII 'A'对应的就是65（GBK所有中国字编码表，Unicode全世界编码表）
	v2 := "A"
	fmt.Printf("%T, %s\n", v1, v1)  // int32, %!s(int32=65)
	fmt.Printf("%T, %s\n", v2, v2)  // string, A

	// 字符串拼接（\是转义字符）
	fmt.Println("hello" + ", \"haha")  // hello, "haha

	var a int = 10
	a++
	fmt.Println(a)

	// 输入输出包 fmt包
	// 输出
	fmt.Println()  // 打印 + 换行
	fmt.Print()    // 打印
	fmt.Printf()   // 格式化打印

	// 输入
	fmt.Scanln()
	fmt.Scan()
	fmt.Scanf()

	var x int
	var y float64

	fmt.Println("请输入整数x,浮点数y")
	fmt.Scan(&x, &y)

	var a int = 20
	if a > 20 {
		fmt.Println("hhh")
	} else if a == 20 {
		fmt.Println("xixi")
	} else {

	}

	switch a {
	case 20:
		fmt.Println("20")  // go语言的switch把case默认都自带break
	case 2:
		fmt.Println("2")
	default:
		fmt.Printf("%s", "ad")

	}

	var b int = 1
	switch b {
	case 1:
		fmt.Print("1 ")
		fallthrough
	case 2:
		fmt.Print("2 ")
		fallthrough
	case 3:
		fmt.Print("3 ")
	default:
		fmt.Print("default")
	}

	var arr int = 12
	for i := 0; i < arr; i++ {
		// ...
	}

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, j*i)
		}
		fmt.Println()
	}*/

	str := "hello"
	fmt.Println("len获取字符串长度：", len(str))
	fmt.Println("字节打印：", str[0])
	fmt.Printf("获取指定的字节，比如索引0处：%c\n", str[0])

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])  // 直接用print打出来的是字节，对应的ASCII码
	}
	fmt.Println("\n循环方式一（常规）↑，循环方式二（range）↓")
	for i, v := range str {  // 下标, 字符 := range 字符串
		fmt.Printf("%d%c ", i, v)
	}

}
