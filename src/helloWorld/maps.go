package main

import "fmt"

type Vertex2 struct {
	Lat, Long float64
}

// 声明创建
var m map[string]Vertex2

// 创建带数值
var m1 = map[string]Vertex2{
	"Bell Labs": Vertex2{
		40.68433, -74.39967,
	},
	"Google": Vertex2{
		37.42202, -122.08408,
	},
}
// 简洁写法
var m2 = map[string]Vertex2{
	"hh1": {11, 22},
	"hh2": {33, 44},
}

func main33() {
	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{
		40.68433, -74.39967,
	}
	m["Google"] = Vertex2{
		37.42202, -122.08408,
	}
	fmt.Println(m["Bell Labs"])  	  // {40.68433 -74.39967}
	fmt.Println(m["Bell Labs"].Lat)   // 40.68433
	fmt.Println(m["Bell Labs"].Long)  // -74.39967

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fabo() func(int) int {
	a := 0
	b := 1
	return func(x int) int {
		 a = b
		 b = a + b
		 return b
	}
}
