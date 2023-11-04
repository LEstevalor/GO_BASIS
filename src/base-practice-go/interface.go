package practice

// 节10- 接口

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func describe(i Abser) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describeTest(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// 节10 - 接口
func main()  {
	/**
		一个接口类型 定义为一组方法签名.
		接口类型的值可保存实现这些方法的任何值.

		接口是隐式实现的，实现中甚至都见不到implements关键字
		甚至可以直接写成 var ab Abser = yFloat(-math.Sqrt2)

		接口值可以被认为是一个值和一个具体类型的元组:(value, type)  默认值(<nil>, <nil>)
		调用接口值上的方法（也就根据type值执行对应方法）会在其底层类型上执行同名方法.
	 */
	var ab Abser
	describe(ab)  // 默认值(<nil>, <nil>)

	f := myFloat(-math.Sqrt2)
	v := VertexT{3, 4}

	ab = f  // a MyFloat implements Abser
	describe(ab)  // (-1.4142135623730951, main.myFloat)
	ab = &v  // a *Vertex implements Abser
	// 切换（Vertex（值类型）不实现 Abser，因 Abs 方法仅在 *Vertex（指针类型）上定义
	// 为什么？对比，因为float64虽然是值类型，但直接返回值，而不是从本身改变，而Vertex的Abs是从本身改变的，需要传递指针

	fmt.Println(ab.Abs())  // 5
	fmt.Println(ab)  // &{3 4}
	fmt.Printf("%T %T %T\n", ab, f, &v)  // *main.Vertex main.myFloat *main.Vertex


	/**
		空接口（指定了零个方法的接口类型）
		interface{} 可以保存任何类型的值。 (每种类型至少实现零个方法。)
		一般用于处理未知类型值
	 */
	var iTest interface{}
	iTest = 66
	describeTest(iTest)  // (66, int)
	iTest = float64(1)
	describeTest(iTest)  // (1, float64)

	/**
		类型断言（提供对接口值的底层具体值的访问）

		t := i.(T)   断言接口值 i 包含具体类型 T，并将底层 T 值赋给变量 t
		value := i.(type) 如果i不持有对应的type会报错

		可以用布尔测试的方法，判断持不持有type（为true持有，否则不持有，不持有将返回对应不持有type的默认空值）
		value, ok := i.(type)

	 */
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)  // hello true

	ff, ok := i.(float64)
	fmt.Println(ff, ok)  // 0 false

	//ff = i.(float64) // 报错panic: interface conversion: interface {} is string, not float64
	//fmt.Println(ff)

	/**
		一般 switch 串联多个类型断言的结构
		switch v := i.(type) {
		case T:
			// 这里 v 的类型是 T
		case S:
			// 这里 v 的类型是 S
		default:
			// 不匹配；这里 v 与 i 具有相同的类型
		}
	 */
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))  // "hello" is 5 bytes long
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type myFloat float64

func (f myFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type VertexT struct {
	X, Y float64
}

func (v *VertexT) Abs() float64 {
	if v == nil {  // 处理空指针
		return 0.0
	}
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
