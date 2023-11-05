package question_test
// WEB爬虫

import (
	"fmt"
)

type Fetcher interface {
	// Fetch 返回 URL 的正文和在该网页上找到的网址切片
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl 使用提取器递归爬网，以 URL 开头的页面，达到最大深度
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO是一个特殊的注释关键字，用于标记代码中需要完成或待办的部分
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.  // 可用 Go 语言的并发特性，如 goroutines 和 channels。此外还需一个映射（map）存储已访问的 URL，避免重复访问
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

/**
func Crawl(url string, depth int, fetcher Fetcher, fetched map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 0 || fetched[url] {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fetched[url] = true
	fmt.Printf("found: %s %q\n", url, body)

	wg.Add(len(urls))
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, fetched, wg)
	}
 */

func main() {
	/**
	var wg sync.WaitGroup  // WaitGroup 用于等待一组 goroutines 的完成。它主要用于同步多个并发任务的执行。WaitGroup 包含一个计数器，Add() 方法增加计数，通过 Done() 方法减少计数。当计数器变为零时，Wait() 方法会解除阻塞
	fetched := make(map[string]bool)

	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, fetched, &wg)
	wg.Wait()
	 */
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher 是返回预制结果的 Fetcher
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是一个填充的 fakeFetcher。
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
