package practice

// 节1 - 只是时间

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("最美的爱情在回忆里待续")

	fmt.Print("时至今时：", time.Now()) // 时至今时：2023-10-25 22:02:50.9913746 +0800 CST m=+0.003758801
	// CST代表"China Standard Time"，即中国标准时间。CST是UTC+8的时区，它是中国大陆、台湾、香港和澳门等地的常用时区。
	// m=+xxx表示相对于某个基准时间的时间差 为0.003758801分钟（约为0.226秒）
}

/**
import导入模块有点奇怪（喜欢写成一行的人真难顶emmm），
可以看到接口都是大写字母开头
这个print的拼接方式，有点像python print("时至今时：", time.now())  // 时至今时：2023-10-25 22:05:53.705069
不过注意，Go的Print在拼接会插入一个空格，而且代码句尾强制不需要”;“结束，类似python的写法
*/
