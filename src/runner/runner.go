package runner

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

/**
Runner 在给定的超时时间内执行一组任务，
并且在操作系统发送中断信号时结束这些任务
*/
type Runner struct {
	//interrupt 通道报告从操作系统发送的信号
	interrupt chan os.Signal
	//complete 通道报告处理任务已经完成
	complete chan error
	//timeout 报告处理任务已经超时
	timeout <-chan time.Time
	//tasks 持有一组以索引顺序依次执行的函数
	tasks []func(int)
}

//设置超时时间
const timeout = 3 * time.Second

var (
	//ErrTimeout 会在任务执行超时时返回
	ErrorTimeout = errors.New("received timeout")
	//ErrInterrupt 会在接收到操作系统的事件时返回
	ErrorInterrupt = errors.New("received interrupt")
)

//New 返回一个新的准备使用的 Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

//Add 将一个任务附加到 Runner 上
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

//验证是否接收到了中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true

	default:
		return false
	}
}

//run 执行每一个已注册的任务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrorInterrupt
		}
		//执行任务
		task(id)
	}
	return nil
}

func (r *Runner) Start() error {
	//接收系统的所有中断信号
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrorTimeout
	}
}

func RunnerTest() {
	log.Println("Starting work.")

	r := New(timeout)
	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {

		case ErrorTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case ErrorInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}

	log.Println("Process ended.")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
