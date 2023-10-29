package practice
// 节8 - map

import "fmt"

type Vertexes struct {
	x, y int
}

func main() {
	// 短变量
	mapOne := make(map[string]Vertexes)
	mapOne["love"] = Vertexes{5, 2}
	mapOne["bye"] = Vertexes{8, 8}
	fmt.Println(mapOne)  // map[bye:{8 8} love:{5 2}]

	// 预定义形式
	var mapTwo map[int]Vertexes  // 此时为空
	fmt.Println(mapTwo)  // map[]
	mapTwo = make(map[int]Vertexes)  // 需要初始化
	fmt.Println(mapTwo)  // map[]
	mapTwo[1] = Vertexes{1, 1}
	mapTwo[2] = Vertexes{2, 2}
	fmt.Println(mapTwo)  // map[1:{1 1} 2:{2 2}]

	// 字面量形式
	var mapThree = map[int]Vertexes{
		1: Vertexes{1, 1},
		2: Vertexes{2, 4},
	}
	// 字面量中甚至可以忽略顶层类型
	var mapFour = map[int]Vertexes{
		1: Vertexes{1, 1},
		2: Vertexes{2, 4},
	}
	fmt.Println(mapThree, mapFour)  // map[1:{1 1} 2:{2 4}]

	/**
		map 操作
	      插入或更新： m[key] = elem
	  	  检索：elem = m[key]
		  删除：delete(m, key)
		  双赋值检测键是否存在：elem, ok := m[key]
			（在则ok为true，不在则ok为false，elem是map元素类型的零值，比如这里就是{0,0}）
	*/
	mapFour[3] = Vertexes{3, 9}
	elem := mapFour[3]
	delete(mapFour, 3)
	elemThree, ok3 := mapFour[3]
	elemTwo, ok2 := mapFour[2]
	fmt.Println(elem, elemThree, ok3, elemTwo, ok2)  // {3 9} {0 0} false {2 4} true
	fmt.Printf("%T", elemThree)  // main.Vertexes


}
/**
创建方式：map[]
map的零值是 nil
 */
