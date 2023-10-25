package main

import "fmt"

type IIII interface {
	M()
}

type TTTT struct {
	S string
}

func (t *TTTT) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i IIII

	var t *TTTT
	i = t
	describe(i)   // (<nil>, *main.TTTT)
	i.M()         // <nil> （这里就没有M方法去实现I）

	i = &TTTT{"hello"}
	describe(i)   // (&{hello}, *main.TTTT)
	i.M()

	var ii IIIII
	describe(ii)  // (<nil>, <nil>)
	//ii.M()  // 错误

	var iii interface{}
	describei(iii)  // (<nil>, <nil>)

	iii = 42
	describei(iii)  // (42, int)

	iii = "hello"
	describei(iii)  // (hello, string)


	var ini interface{} = "hello"

	s := ini.(string)
	fmt.Println(s)

	s, ok := ini.(string)
	fmt.Println(s, ok)

	f, ok := ini.(float64)
	fmt.Println(f, ok)

	f = ini.(float64) // panic 报错
	fmt.Println(f)
}

type IIIII interface {
	M()
}

func describei(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
