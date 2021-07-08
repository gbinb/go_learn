package main

/*
隐藏类型
*/
import (
	"fmt"
)

func Protect() {
	counter := alertCounter(10)

	fmt.Printf("Counter: %d\n", counter)
}
