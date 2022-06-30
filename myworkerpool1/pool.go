package myworkerpool1

import (
	"errors"
	"fmt"
	"sync"
)

type Task func()

const (
	defaultCapacity = 100
	maxCapacity     = 10000
)

var (
	ErrNoIdleWorkerInPool = errors.New("no idle worker in pool")
	ErrWorkerPoolFreed    = errors.New("workerpool freed")
)

type Pool struct {
	capacity int
	wg       sync.WaitGroup
	quit     chan struct{}
	tasks    chan Task
	active   chan struct{}
}

func New(capacity int) *Pool {
	if capacity <= 0 {
		capacity = defaultCapacity
	}
	if capacity > maxCapacity {
		capacity = maxCapacity
	}

	p := &Pool{
		capacity: capacity,
		quit:     make(chan struct{}),
		tasks:    make(chan Task),
		active:   make(chan struct{}, capacity),
	}
	fmt.Println("Starting a new workerpool...")
	p.run()
	return p
}

func (p *Pool) run() {
	idx := 0

	for {
		select {
		case <-p.quit:
			return
		case p.active <- struct{}{}:
			idx++
			p.newWorker(idx)
		}
	}
}

func (p *Pool) newWorker(idx int) {
	p.wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("worker [%03d]:recover panic [%s] and exit", idx, err)
				<-p.active
			}
			p.wg.Done()
		}()

		fmt.Printf("worker [%03d]: start\n", idx)

		for {
			select {
			case <-p.quit:
				fmt.Printf("worker[%03d]: exit\n", idx)
				<-p.active
				return
			case t := <-p.tasks:
				fmt.Printf("worker[%03d]:recive a task\n", idx)
				t()
			}
		}
	}()
}

func (p *Pool) Schedule(t Task) error {
	select {
	case <-p.quit:
		return ErrWorkerPoolFreed
	case p.tasks <- t:
		return nil
	}
}

func (p *Pool) Free() {
	close(p.quit)
	p.wg.Wait()
	fmt.Println("workerpool freed")
}
