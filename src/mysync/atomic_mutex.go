package mysync

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter2 int
	wg2      sync.WaitGroup
	mutex    sync.Mutex
)

func incCounter2(id int) {
	defer wg2.Done()
	for count := 0; count < 2; count++ {
		//同一时刻只允许一个 goroutine 进入这个临界区
		mutex.Lock()
		{
			value := counter2
			runtime.Gosched()
			value++
			counter2 = value
		}
		mutex.Unlock()
	}
	fmt.Printf("doWork %d\n", id)
}

func MutexTest() {
	wg2.Add(2)

	go incCounter2(1)
	go incCounter2(2)

	wg2.Wait()
	fmt.Printf("Final Counter: %d \n", counter2)
}
