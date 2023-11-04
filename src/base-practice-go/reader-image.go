package practice

import (
	"fmt"
	"image"
	"io"
	"strings"
)

// 节12, Reader及Image接口

func main()  {
	/**
		Reader接口（io 包指定 io.Reader 接口，表示数据流的读取端）
			提供Reader方法 func (T) Read(b []byte) (n int, err error)
			Read 用数据填充给定的字节片，并返回填充的字节数和错误值。返回一个 io.EOF 错误当流结束时（正常情况读到结束）

		标准库包含此接口的许多实现，包括文件、网络连接、压缩、密码等
	 */
	r := strings.NewReader("为什么我连分开都迁就着你")
	b := make([]byte, 18)

	fmt.Println(r, b)  // &{为什么我连分开都迁就着你 0 -1} [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			fmt.Println("结束啦")
			break
		}
	}  /**
		n = 18 err = <nil> b = [228 184 186 228 187 128 228 185 136 230 136 145 232 191 158 229 136 134]
		b[:n] = "为什么我连分"
		n = 18 err = <nil> b = [229 188 128 233 131 189 232 191 129 229 176 177 231 157 128 228 189 160]
		b[:n] = "开都迁就着你"
		n = 0 err = EOF b = [229 188 128 233 131 189 232 191 129 229 176 177 231 157 128 228 189 160]
		b[:n] = ""
		结束啦
	 */

	/**
	// 实现一个 Reader 类型，该类型发出无限的 ASCII 字符 'A' 流
	reader := AReader{}
	buffer := make([]byte, 10)

	n, err := reader.Read(buffer)
	if err != nil {
		panic(err)
	}

	println(string(buffer[:n])) // 输出：AAAAAAAAAA
	 */

	/**
		Image
		image包定义了 Image 接口
		package image

		type Image interface {
			ColorModel() color.Model
			Bounds() Rectangle
			At(x, y int) color.Color
		}

	 */
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds(), m.ColorModel())  // (0,0)-(100,100) &{0x40b340}
	fmt.Println(m.At(0, 0).RGBA())  // 0 0 0 0
}

type AReader struct{}

func (r AReader) Read(p []byte) (int, error) {
	for i := 0; i < len(p); i++ {
		p[i] = 'A'
	}
	return len(p), nil
}
