package practice
// 节14 sync.Mutex 互斥锁
import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter 可以安全同时使用。
type SafeCounter struct {
	mu sync.Mutex
	mapV map[string]int
	testNum int
}

// Inc 递增给定键的计数器
func (c *SafeCounter) Inc(key string)  {
	c.mu.Lock()
	// 锁定，以便一次只有一个 goroutine 可以访问地图 c.mapV.
	c.mapV[key]++
	c.mu.Unlock()
}

// Value 返回给定键的计数器的当前值。
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()  // 锁住避免被篡改，defer是执行方法后再执行的，值已经被返回使用
	defer c.mu.Unlock()
	return c.mapV[key]
}

func main() {
	/**
		Mutex 互斥锁（Go标准库）
		已经看到了channel对于 goroutines 之间的通信非常有用
		但若不需要通信呢？若只想确保一次只有一个 goroutine 可访问一个变量以避免冲突怎么办? 互斥
		提供它的数据结构的常规名称是 mutex ，提供两个方法 Lock  Unlock 来包围需要互斥的片段

	 */
	c := SafeCounter{mapV: make(map[string]int)}
	for i := 0; i < 100; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))  // 100
}
