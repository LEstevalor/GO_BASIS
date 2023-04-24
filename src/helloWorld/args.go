package main

import "fmt"

func main3() {
	/*
	arr1 := []int{1, 2, 3, 4}  // go中可拓展数组属于引用传递
	fmt.Println("arr1原数据：", arr1)
	arrChange(arr1)
	fmt.Println("arr1原数据：", arr1)
	fmt.Printf("%T", f)  // 结果：func()
	 */
	var f6 func(int, int) int  // func(参数列表) 返回类型  如果不要参数列表/返回类型就可以不写
	f6 = f1
	fmt.Println(f1)            // 0x889780
	fmt.Println(f6)            // 0x889780
}
func f1(a, b int) int {
	return a + b
}
func f()  {

}
func arrChange(arr []int)  {  // 这里是引用传递了
	arr[0] = 100
	fmt.Println("arr数据：", arr)
}
