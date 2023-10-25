package main

import (
	"fmt"
)

var (
	name     string = "hh"
	stream   byte   = 12
	skip     rune   = 12
	username string = "xx"
	MaxInt   uint64 = 1<<64 - 1
)

var (
	v1 = Vertex{1, 2}  // {1 2}
	v2 = Vertex{X: 1}  // Y:0 is implicit 结果{1 0}
	v3 = Vertex{}      // X:0 and Y:0 结果{0 0}
	p  = &Vertex{1, 2} // 0xc00000e0c8
)

func main11() {
	fmt.Println(addd(1, 2))

	i, j := 42, 2701

	p := &i

	fmt.Println(*p)

	*p = 21
	fmt.Println(i)

	p = &j
	fmt.Println(*p / 37)

	v := Vertex{1, 2}
	fmt.Printf("%d, %d", v.X, v.Y)

	fmt.Println(v1, v2, v3, p)

	var aa [10]int // 声明
	aa[1] = 44
	aa[4] = 66
	fmt.Println(aa) // [0 44 0 0 66 0 0 0 0 0]

	primes := [6]int{2, 3, 5, 7, 11, 13} // 赋值声明
	fmt.Println(primes)                  // [2 3 5 7 11 13]

	nums := [][]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println(nums) // [[1 2] [3 4]]

	//var numns [1][2] int

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r) // [true false true true false true]

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	// [{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}]

	var sss []int
	sss = primes[:]
	fmt.Printf("len=%d cap=%d \n", len(sss), cap(sss))
	// len=6 cap=6
	sss = append(sss, 1, 2)
	fmt.Printf("len=%d cap=%d \n", len(sss), cap(sss))
	// len=8 cap=12
	fmt.Println(primes) // [2 3 5 7 11 13]
	sss = append(sss, 1, 5, 4, 5, 6, 4)
	fmt.Printf("len=%d cap=%d \n", len(sss), cap(sss))
	// len=14 cap=24

	for i, v := range primes {
		fmt.Println(i, v)
	}
	/*
		0 2
		1 3
		2 5
		3 7
		4 11
		5 13
	*/

}

func addd(x, y int) (sum int) {
	//sum = x + y
	return
}

// 结构体，名字为Vertex
type Vertex struct {
	X int
	Y int
}

type Vertex1 struct {
	X, Y int
}