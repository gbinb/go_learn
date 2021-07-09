package sample

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func RuntimeTest2() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("Final Counter: ", counter)
}

func RuntimeTest() {
	//分配一个逻辑处理器给调度器使用
	runtime.NumCPU()
	runtime.GOMAXPROCS(1)

	//wg 用来等待程序完成
	//计数加 2，表示要等待两个 goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Printf("Start Goroutines")

	//声明一个匿名函数，并创建一个 goroutine
	go func() {
		//在函数退出时调用 Done 来通知 main 函数工作已经完成
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Printf("\n")
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Printf("\n")
		}
	}()

	fmt.Println("Waiting To Finish")
	//等待 goroutine 结束
	wg.Wait()

	fmt.Println("\nTerminating Program")
}

func incCounter(id int) {
	defer wg.Done()

	for i := 0; i < 2; i++ {
		value := counter

		//当前 goroutine 从线程退出，并放回到队列
		//用于将 goroutine 从当前线程退出， 给其他 goroutine 运行的机会
		runtime.Gosched()

		value++
		counter = value
	}

	fmt.Println(id, "Final Counter: ", counter)
}
