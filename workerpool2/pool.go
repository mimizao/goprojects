package workerpool2

import (
	"errors"
	"fmt"
	"sync"
)

var (
	ErrNoIdleWorkerInPool = errors.New("no idle worker in pool")
	ErrWorkPoolFreed      = errors.New("workerpool freed")
)

type Pool struct {
	capacity int // workerpool大小

	active chan struct{}
	tasks  chan Task

	wg   sync.WaitGroup
	quit chan struct{}

	preAlloc bool
	block    bool
}

type Task func()

const (
	defaultCapacity = 100
	maxCapacity     = 10000
)

func New(capacity int, opts ...Option) *Pool {
	if capacity <= 0 {
		capacity = defaultCapacity
	}
	if capacity > maxCapacity {
		capacity = maxCapacity
	}

	p := &Pool{
		capacity: capacity,
		tasks:    make(chan Task),
		quit:     make(chan struct{}),
		active:   make(chan struct{}, capacity),
	}

	for _, opt := range opts {
		opt(p)
	}

	fmt.Printf("workerpool start(preAlloc=%t)\n", p.preAlloc)

	if p.preAlloc {
		for i := 0; i < p.capacity; i++ {
			p.newWorker(i + 1)
			p.active <- struct{}{}
		}
	}

	go p.run()

	return p
}

func (p *Pool) newWorker(i int) {
	p.wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("worker[%03d]:recover panic[%s] and exit\n", i, err)
				<-p.active
			}
			p.wg.Done()
		}()

		fmt.Printf("worker[%03d]: start\n", i)
		for {
			select {
			case <-p.quit:
				fmt.Printf("worker[%03d]:exit\n", i)
				<-p.active
				return
			case t := <-p.tasks:
				fmt.Printf("worder[%03d]: receive a task\n", i)
				t()
			}
		}
	}()
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

func (p *Pool) Schedule(t Task) error {
	select {
	case <-p.quit:
		return ErrWorkPoolFreed
	case p.tasks <- t:
		return nil
	}
}

func (p *Pool) Free() {
	close(p.quit)
	p.wg.Wait()
	fmt.Println("workerpool freed")
}
