package pool

import (
	"log"
	"sync"
)

// Func is a function that can be executed by Pool.
type Func = func() error

// Pool is a worker which has pools to executes tasks.
type Pool struct {
	cap int
	buf chan Func
	rec chan interface{}
	wg  sync.WaitGroup
}

// New creates a new pool with given capacity.
func New(cap int) *Pool {
	w := &Pool{
		cap: cap,
		buf: make(chan Func, cap),
		rec: make(chan interface{}, cap),
	}
	w.init()
	return w
}

func (w *Pool) init() {
	f := func() {
		defer w.wg.Done()
		defer func() {
			if r := recover(); r != nil {
				log.Println("recovered", r)
				w.rec <- r
			}
		}()
		for f := range w.buf {
			if f != nil {
				if err := f(); err != nil {
					log.Println("task error", err)
				} else {
					log.Println("task done")
				}
			}
		}
	}
	for i := 0; i < w.cap; i++ {
		w.wg.Add(1)
		go f()
	}

	go func() {
		for range w.rec {
			w.wg.Add(1)
			go f()
		}
	}()
}

// Execute adds a task to the pool.
func (w *Pool) Execute(f Func) {
	w.buf <- f
}

// ShutDown stands for shutdown the pool, no new tasks will be accepted.
func (w *Pool) ShutDown() {
	close(w.buf)
}

// Await waits for all tasks to be done.
func (w *Pool) Await() {
	w.wg.Wait()
	close(w.rec)
}
