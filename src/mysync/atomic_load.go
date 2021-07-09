package mysync

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	showdown int64
	wg1      sync.WaitGroup
)

func doWork(name string) {
	defer wg1.Done()

	for {
		fmt.Printf("Doing %s Work\n", name)
		if atomic.LoadInt64(&showdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}

func WorkTest() {
	wg1.Add(2)

	go doWork("A")
	go doWork("B")

	time.Sleep(1 * time.Second)
	fmt.Printf("Shutdown Now")
	atomic.StoreInt64(&showdown, 1)
	wg1.Wait()
}
