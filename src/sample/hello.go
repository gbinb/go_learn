package main

import (
	"fmt"
	"pool"
)

func structSamp() {
	user := User{
		Name:  "张三",
		Email: "san@126.com",
		Age:   10,
		Man:   false,
	}
	PrintUser(user)
	Modify(user, "李四")
	PrintUser(user)
	ModifyRef(&user, "李四")
	PrintUser(user)

	tony := User{"tony", "tony@163.com", 20, true}
	tony.notify()
	tony.changeEmail("hello@163.com")
	tony.notify()

	fmt.Printf("\n")

	poly := &User{"poly", "poly@163.com", 10, false}
	poly.notify()
	poly.changeEmail("poly@126.com")
	poly.notify()
}

func referenceSamp() {
	str := "go language"
	ModifyStr(str)
	fmt.Printf("%s\n", str)
}

func netSample() {
	Curl("https://www.baidu.com")
}

func main() {
	//fmt.Println("Hello World!!")

	//referenceSamp()
	//netSample()

	//GenericTest()
	//Protect()

	//RuntimeTest()
	//RuntimeTest2()

	//mysync.AtomicTest()
	//mysync.WorkTest()
	//mysync.MutexTest()
	//mysync.NoBufChannelTest()
	//mysync.NoBufRunnerTest()
	//mysync.BufTaskTest()
	//runner.RunnerTest()
	pool.PoolTest()
}
