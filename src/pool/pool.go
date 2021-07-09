package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

/**
Pool 管理一组可以安全地在多个 goroutine 间共享的资源。
被管理的资源必须实现 io.Closer 接口
*/
type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

var ErrPoolClosed = errors.New("Pool has been closed.")

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small.")
	}

	return &Pool{
		resources: make(chan io.Closer, size),
		factory:   fn,
	}, nil
}

//Acquire 从池中获取一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

//Release 将一个使用后的资源放回池里
func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	//如果池已经被关闭，销毁这个资源
	if p.closed {
		r.Close()
		return
	}

	select {
	//试图将这个资源放入队列
	case p.resources <- r:
		log.Println("Release:", "In Queue")
	//如果队列已满，则关闭这个资源
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}

//Close 会让资源池停止工作，并关闭所有现有的资源
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true
	//在清空通道里的资源之前，将通道关闭
	//如果不这样做，会发生死锁
	close(p.resources)

	//关闭资源
	for r := range p.resources {
		r.Close()
	}
}
