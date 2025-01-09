package threadpool

import (
	"fmt"
	"sync"
)

// Task represents a unit of work to be executed by the thread pool.
type Task struct {
	Func     func(uniqueID string)
	UniqueID string
}

type ThreadPool struct {
	poolSize int

	queue chan Task

	wg sync.WaitGroup

	mu sync.Mutex // Protects the closed flag

	closed bool
}

// NewThreadPool creates a new thread pool with the given number of workers and queue size.

func NewThreadPool(poolSize int, queueSize int) *ThreadPool {

	pool := &ThreadPool{

		poolSize: poolSize,

		queue: make(chan Task, queueSize),
	}

	for i := 0; i < pool.poolSize; i++ {
		go startWorkers(pool.queue)
	}

	return pool

}

// startWorkers starts the worker goroutines.

func startWorkers(queue chan Task) {

	for task := range queue {
		task.Func(task.UniqueID)
	}

}

// Execute submits a task to the thread pool.

// It returns an error if the pool is closed or the queue is full.

func (p *ThreadPool) Execute(task Task) error {

	p.mu.Lock()

	defer p.mu.Unlock()

	if p.closed {

		return fmt.Errorf("thread pool is closed")

	}

	select {

	case p.queue <- task:

		return nil

	default:

		return fmt.Errorf("task queue is full")

	}

}

// Close shuts down the thread pool and waits for all workers to finish.

func (p *ThreadPool) Close() {

	p.mu.Lock()

	defer p.mu.Unlock()

	if p.closed {

		return

	}

	p.closed = true

	close(p.queue)

	p.wg.Wait()

}
