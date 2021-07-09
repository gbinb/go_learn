package mysync

import (
	"fmt"
	"sync"
)

import (
	"time"
)

var (
	wg4 sync.WaitGroup
)

func Runner(baton chan int) {
	var newRunner int

	runner := <-baton
	fmt.Printf("Runner %d Running With Baton\n", runner)

	//create new runner
	if runner < 4 {
		newRunner = runner + 1

		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton)

		time.Sleep(100 * time.Millisecond)
	} else {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg4.Done()
		return
	}

	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)

	baton <- newRunner
}

func NoBufRunnerTest() {
	baton := make(chan int)
	wg4.Add(1)

	go Runner(baton)

	baton <- 1
	wg4.Wait()
}
