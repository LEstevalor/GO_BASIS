package practice
// 节13, 类型参数及泛型类型（/泛型函数及泛型类型）
import "fmt"

func IndexT[T comparable](s []T, x T) int {
	// comparable关键字用于指定类型参数T必须是可比较的类型。这意味着在使用==运算符进行比较时，类型T的值必须支持相等比较操作
	// 目的是为了在编译时进行类型检查，以可比较
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func printTest[T any](s T) {  // 将IndexT简易化演示类型参数，使用[T any]来定义类型参数，没有[T any]将无法使用泛型函数
	fmt.Println(s)
}

type List[T any] struct {  // any关键字是必要的，表示T可以是任意类型，跟以上comparable
	next *List[T]
	val  T
}

func main()  {
	/**
		类型参数
			类型参数编写 Go 函数以处理多种类型，比如 func Index[T comparable](s []T, x T) int
			comparable是一个有用的约束，意味着它是满足内置约束。 也是同一类型的值

		泛型类型
			除了泛型函数之外，Go 还支持泛型类型，可实现泛型数据结构

	 */
	nums := []int{1, 2, 3, 4}
	fmt.Println(IndexT(nums, 1))  // 0

	strings := []string{"你总是太晚明白", "最后才把话说开", "宁愿没出息"}
	fmt.Println(IndexT(strings, "最后才把话说开"))  // 1

	testList := List[int]{next: nil, val: 12}
	fmt.Println(testList)  // &{<nil> 12}

	printTest(1)  // 1
}
