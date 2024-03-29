package src

import "fmt"

/*
功能：Go 语言的 fmt 包是 "format" 的简称。
	fmt 包提供了一系列函数，用于格式化和处理各种类型的数据，如数字、字符串、时间等。
	这些函数可将数据转为字符串，也可将字符串解析为各种数据类型。fmt 包是非常重要的基础包之一。
*/

func main() {
	fmt.Println("hello")
}

/*
语法：Go中如果导入的包或变量不使用，是会冒编译错误

命名规范（文件名、文件夹名和变量名）【命名规范的关键是简洁、清晰和一致。作用：提高代码的可读性和可维护性】：
	以下都有的一点：避免使用 Go 语言的关键字和保留字
	文件名和文件夹名：
		使用小写字母，不使用下划线或其他分隔符。
		多个单词之间用短横线（-）连接，例如：my-project。
	变量名：
		变量名应该使用驼峰命名法（CamelCase），即首字母小写，单词之间不使用下划线。
		避免使用单个字母作为变量名，除非它是一个简单的循环计数器。（比如只是简单的for间接变量）
	包名：
		包名应该使用简短的、有意义的名称，最好使用小写字母，不使用下划线或其他分隔符。
	常量名、枚举名：
		常量名、枚举名应该使用大写字母和下划线来分隔单词，例如：MAX_COUNT
	函数名：
		函数名应该使用驼峰命名法（camelCase），即首字母小写，单词之间不使用下划线。
	结构体名和接口名：
		结构体和接口名应该使用大写字母开头，驼峰命名法（CamelCase），例如：UserInfo 和 FileReader，上面的Println也是如此。

	注：在Go语言中，变量、函数、文件和包的命名可以包含数字。根据Go语言的命名约定，数字可以作为命名的一部分，但不应该作为命名的开头。
		（Go语言中的数字更偏向于英文使用，比如 usage_3d，plugin-v2 这类，而不是随意的 fun1 fun2 这种）

*/
