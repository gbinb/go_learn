package pool

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines   = 25
	pooledResources = 2
)

var idCounter int32

type dbConnection struct {
	ID int32
}

func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

func creatConnect() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)
	return &dbConnection{id}, nil
}

func performQuery(query int, p *Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
	}
	defer p.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}

func PoolTest() {
	var wg6 sync.WaitGroup
	//计数加25，表示要等待25个 goroutine
	wg6.Add(maxGoroutines)

	p, err := New(creatConnect, pooledResources)
	if err != nil {
		log.Println(err)
	}

	for query := 0; query < maxGoroutines; query++ {
		//每个 goroutine 需要自己复制一份要查询值的副本，
		//不然所有的查询会共享同一个查询变量
		go func(q int) {
			performQuery(q, p)
			wg6.Done()
		}(query)
	}

	wg6.Wait()
	//关闭池
	log.Println("Shutdown Program.")
	p.Close()
}
