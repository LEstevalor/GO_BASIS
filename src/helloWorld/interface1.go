package main

import "fmt"

type II interface {
	M()
}

type TT struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.

func (t TT) M() {
	fmt.Println(t.S)
}

func main7() {
	var i II = TT{"hello"}  // TT实现接口II，通过TT调到M
	i.M()  // hello
}
