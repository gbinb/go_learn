package mysync

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	wg3 sync.WaitGroup
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func player(name string, court chan int) {
	defer wg3.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			//关闭通道
			close(court)
			return
		}
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		//将球写入通道
		court <- ball
	}
}

func NoBufChannelTest() {
	//创建一个无缓冲通道
	court := make(chan int)
	wg3.Add(2)

	go player("Mike", court)
	go player("Jack", court)

	court <- 1
	wg3.Wait()
}
