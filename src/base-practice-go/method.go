package practice

// 节9 - 方法

import (
	"fmt"
	"math"
)
// 节9 - 方法

type VertexMethod struct {
	x, y string
}

func (v VertexMethod) StringAnd() string {
	res := v.x + v.y
	fmt.Println(res)
	return res
}

func StringAmazing(method VertexMethod) bool {
	fmt.Println(method.x + method.y + "!")
	return true
}

type myInt int  // 只能声明具有接收器的方法，不能使用其类型在另一个包中定义的接收器（包括内置类型，例如 int ）来声明方法

func (mi myInt) Abs() float64 {
	return math.Abs(float64(mi))
}

// NextWord 指针接收器，相当于直接修改原始副本
func (v *VertexMethod) NextWord() {
	fmt.Printf("NextWord地址：%p 类型：%T\n", v, v)  // NextWord地址：0xc0000983a0 类型：*main.VertexMethod
	v.x = v.x + "，"
	v.y = v.y + "！"
}

func (v VertexMethod) NextWordTest() {
	fmt.Printf("NextWord地址：%p 类型：%T\n", v, v)  // NextWord地址：%!p(main.VertexMethod={风筝在， 阴天搁浅！}) 类型：main.VertexMethod
	v.x = v.x + "，"
	v.y = v.y + "！"
}

func ScaleFunc(v *VertexMethod, string2 string)  {
	v.x = v.x + string2
}

func ScaleFuncTest(v VertexMethod, string2 string)  {
	v.x = v.x + string2
}

func main()  {
	/**
		方法：（方法只是一个带有接收器参数的函数）
		Go 没有类。但可在类型上定义方法。
		具有特殊 接收器 参数的函数：接收器出现在 func 关键字和方法名称之间的自己的参数列表中.
		比如：func xxx
	 */
	// 两种不同的调用方式
	// 函数类型上定义方法
	v1 := VertexMethod{"风筝在", "阴天搁浅"}
	v1.StringAnd()
	// 函数参数放接收器
	StringAmazing(v1)

	mi := myInt(-1)
	fmt.Println(mi.Abs())

	/**
		可使用指针接收器声明方法
		优处：
			方法能够修改其接收者指向的值.
			第二种是避免在每次方法调用时复制值。例如，如果接收器是大型结构，则效率更高.
	*/
	v1.NextWord()  // 如果是参数 NextWord(v *VertexMethod)，那么此时应该传参 &v1
	fmt.Println(v1)
	v1.NextWordTest()
	fmt.Println(v1)

	/**
		方法与指针间接寻址
			带有指针参数的函数必须接受一个指针:
				ScaleFunc(v1, "1")  // 编译错误！
				ScaleFunc(&v1, "1") // OK
			以指针为接收器的方法被调用时，接收器既能为值又能为指针
				v1.NextWord()  // OK
				v2 := &v1
				v2.NextWord()  // OK
			Go 将语句 v1.NextWord() 解释为 (&v1).NextWord() ，因为 NextWord 方法有一个指针接收器.

		相反
			接受一个值作为参数的函数必须接受一个指定类型的值
				ScaleFuncTest(v1)  // OK
				ScaleFuncTest(&v1)  // 编译错误！
			以值为接收器的方法被调用时，接收器既能为值又能为指针
				v1.NextWordTest()  // OK
				(&v1).NextWordTest()  // OK 此时因为方法里参数为值 &v1 会被解释成 *(&v1)
	 */

}
