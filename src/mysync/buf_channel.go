package mysync

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
定义常量
*/
const (
	numberGoroutine = 4
	taskLoad        = 10
)

var (
	wg5 sync.WaitGroup
)

func init() {
	//初始化随机数种子
	rand.Seed(time.Now().Unix())
}

func worker4(tasks chan string, worker int) {
	defer wg5.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}

		fmt.Printf("Worker: %d : Started %s\n", worker, task)
		sleep := rand.Int31n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}

func BufTaskTest() {
	//创建一个缓冲通道
	tasks := make(chan string, taskLoad)
	wg5.Add(numberGoroutine)

	for gr := 1; gr <= numberGoroutine; gr++ {
		go worker4(tasks, gr)
	}

	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	//当所有工作都处理完时关闭通道,以便所有 goroutine 退出
	close(tasks)
	wg5.Wait()
}
