package main

import (
"fmt"
"math"
)

type Abser interface {
	Abs() float64
}

func main8() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex5{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex5 implements Abser

	// In the following line, v is a Vertex5 (not *Vertex5)
	// and does NOT implement Abser.
	//a = v  // 报错

	fmt.Println(a.Abs())  // 5
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex5 struct {
	X, Y float64
}

func (v *Vertex5) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
