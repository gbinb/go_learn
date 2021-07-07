package search

import (
	"fmt"
	"log"
)

//匹配结果
type Result struct {
	Field   string
	Content string
}

//定义了Matcher接口，默认规范使用`er`结尾
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

//Match 函数，为每个数据源单独启动 goroutine 来执行这个函数
//并发地执行搜索
// chan 通道关键字
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	//对特定的匹配器执行搜索
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	//将结果写入通道
	for _, result := range searchResults {
		results <- result
	}
}

//Display 从每个单独的 goroutine 接收到结果后在终端窗口输出
func Display(results chan *Result) {
	for result := range results {
		fmt.Printf("%s:%s \n\n", result.Field, result.Content)
	}
}
