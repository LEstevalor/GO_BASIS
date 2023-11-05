package question_test

/** 等效二叉树
import (
	"fmt"
)

type TreeTest struct {
	Left  *Tree
	Value int
	Right *Tree
}

type Tree = TreeTest

func Walk(t *Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Same(t1, t2 *Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()

	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if ok1 != ok2 || v1 != v2 {
			return false
		}

		if !ok1 {
			break
		}
	}

	return true
}

func main() {
	t1 := TreeTest.New(1)
	t2 := TreeTest.New(1)
	t3 := TreeTest.New(2)

	fmt.Println("t1:", t1)
	fmt.Println("t2:", t2)
	fmt.Println("t3:", t3)

	fmt.Println("Same(t1, t2):", Same(t1, t2))
	fmt.Println("Same(t1, t3):", Same(t1, t3))
}

 */
