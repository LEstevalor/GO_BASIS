package main

import (
	"fmt"
	"math"
)

type Vertex3 struct {
	X, Y float64
}

func (v Vertex3) Abs() float64 {
	v.X = 4
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex3) Abss() float64 {
	v.X = 3
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main6() {
	v := Vertex3{2, 4}
	fmt.Println(v.Abs())  // 5
	fmt.Println(v.X)      // 2
	fmt.Println(v.Abss()) // 5
	fmt.Println(v.X)      // 3

	v1 := Vertex4{3, 4}
	v1.Scale(2)
	ScaleFunc(&v1, 10)   // 带有指针参数的函数必须接受一个指针，ScaleFunc(v1, 10)报红
	fmt.Println(v1)        // {60 80}

	p := &Vertex4{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	var i I = T{"hello"}
	i.M()
}

type Vertex4 struct {
	X, Y float64
}

func (v *Vertex4) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex4, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type I interface {
	M()
}

type T struct {
	S string
}
func (t T) M() {
	fmt.Println(t.S)
}
