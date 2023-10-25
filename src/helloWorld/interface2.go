package main

import (
	"fmt"
	"math"
)

type III interface {
	M()
}

type TTT struct {
	S string
}

func (t *TTT) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main10() {
	var i III

	i = &TTT{"Hello"}
	describe(i)   // (&{Hello}, *main.TTT)  Hello
	i.M()

	i = F(math.Pi)
	describe(i)   // (3.141592653589793, main.F)  3.141592653589793
	i.M()
}

func describe(i III) {
	fmt.Printf("(%v, %T)\n", i, i)
}
