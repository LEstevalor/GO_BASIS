package practice

// 节5 - 流程控制语句 for if switch defer

import (
	"fmt"
	"runtime"
)

func main() {
	sum := 0
	stringConcat := "ai"
	for i := 0; i < 10; i++ {
		sum += i
		stringConcat += "ai"
	}
	for sum < 50 {
		sum++
	}
	fmt.Println(sum, stringConcat)

	/*
		for循环（与其他语言不一样的在于没 “()”，但外的”{}“必须在）：
			把分号分开的三个语句分别称为：init 语句，条件表达式，post 语句
			同理的，init和post可要可不要
		有个特殊的点，甚至可以把前后的”; ;“分号去了，这时候就同C语言的while循环 （GO没有专门的while语句，单纯只有条件表达式的for就是）

		for {
			// 无限循环
		}
	*/

	if sum == 50 {
		fmt.Println("我舍不得离开哎")
	}

	if sum++; sum == 51 {
		fmt.Println("能比我知道")
	} else if sum == 52 {
		fmt.Println("你的微笑像羽毛")
	} else {
		fmt.Println("多想藏着你的好", sum)
	}

	/*
		if语句也类似，不需要“()”但一定要“{}”
		但有个特殊的地方，类似for， if 语句可以在条件之前以语句执行 （甚至能在else里用）
	*/

	var word string = "晴天"
	switch word {
	case "晴天":
		fmt.Println("希望再见，一定是", word)
	case "雨天":
		fmt.Println("多希望话题不断")
	case "阴天":
		fmt.Println("只想为你撑伞")
	default:
		fmt.Println("末日")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("我在等待重来")
	case "windows":
		fmt.Println("欠你的宠爱")
	case "mac":
		fmt.Println("天空仍灿烂")
	default:
		fmt.Println("爱成了阻碍")
	}

	/*
		switch 语句是编写一连串 if - else 语句的简便方法
		类似 if，switch 语句可以在条件之前以语句执行
	*/

	for i := 1; i <= 3; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("爱你穿越时间")

	defer fmt.Println("我要的只是你在我身边")
	fmt.Println("让爱渗透了地面")

	/*
			defer 语句推迟函数的执行，直到周围的函数返回.
			defer栈 （后进先出）

		爱你穿越时间
		让爱渗透了地面
		我要的只是你在我身边
		3
		2
		1
	*/

}
