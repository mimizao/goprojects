// package main

// import (
// 	"sync"
// 	"testing"
// )

// var cs = 0
// var mu sync.Mutex
// var c = make(chan struct{}, 1)

// func critiacalSectionSyncByMutex() {
// 	mu.Lock()
// 	cs++
// 	mu.Unlock()
// }

// func criticalSectionSyncByChan() {
// 	c <- struct{}{}
// 	cs++
// 	<-c
// }

// func BenchmarkCriticalSectionSyncByMutex(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		critiacalSectionSyncByMutex()
// 	}
// }

// func BenchmarkCriticalSectionSyncByMutexInParallel(b *testing.B) {
// 	b.RunParallel(func(pb *testing.PB) {
// 		for pb.Next() {
// 			criticalSectionSyncByChan()
// 		}
// 	})
// }

// func BenchmarkCriticalSectionSyncByChan(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		criticalSectionSyncByChan()
// 	}
// }

// func BenchmarkCriticalSectionSyncByChanInParallel(b *testing.B) {
// 	b.RunParallel(func(pb *testing.PB) {
// 		for pb.Next() {
// 			criticalSectionSyncByChan()
// 		}
// 	})
// }

/* 不要复制sync的锁
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	i := 0
	var mu sync.Mutex

	wg.Add(1)
	go func(mu1 sync.Mutex) {
		mu1.Lock()
		i = 10
		time.Sleep(10 * time.Second)
		fmt.Printf("g1: i = %d\n", i)
		mu1.Unlock()
		wg.Done()
	}(mu)

	time.Sleep(time.Second)

	mu.Lock()
	i = 1
	fmt.Printf("g0: i = %d\n", i)
	mu.Unlock()

	wg.Wait()
}
*/

// 条件变量

/* 这里是通过条件轮询来实现的
package main

import (
	"fmt"
	"sync"
	"time"
)

type signale struct{}

var ready bool

func worker(i int) {
	fmt.Printf("worker %d is working...\n", i)
	time.Sleep(time.Second * 1)
	fmt.Printf("worker %d works done\n", i)
}

func spawnGroup(f func(i int), num int, mu *sync.Mutex) <-chan signale {
	c := make(chan signale)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			for {
				mu.Lock()
				// 其实就是在这里实现的轮训
				if !ready {
					mu.Unlock()
					time.Sleep(100 * time.Millisecond)
					continue
				}
				mu.Unlock()
				fmt.Printf("worker %d: start to work...\n", i)
				f(i)
				wg.Done()
				return
			}
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- signale(struct{}{})
	}()
	return c
}

func main() {
	fmt.Println("start a group of workers...")
	mu := &sync.Mutex{}
	c := spawnGroup(worker, 5, mu)

	time.Sleep(5 * time.Second)
	fmt.Println("the group of workers start to work...")

	mu.Lock()
	ready = true
	mu.Unlock()

	<-c
	fmt.Println("the group of workers work done")
}
*/

// 通过sync.Cond来实现
package main

import (
	"fmt"
	"sync"
	"time"
)

type signale struct{}

var ready bool

func worker(i int) {
	fmt.Printf("worker %d is working...\n", i)
	time.Sleep(1 * time.Second)
}

func spawnGroup(f func(i int), n int, groupSignal *sync.Cond) <-chan signale {
	c := make(chan signale)
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			groupSignal.L.Lock()
			for !ready {
				groupSignal.Wait()
			}
			groupSignal.L.Unlock()
			fmt.Printf("worker %d: start to work...\n", i)
			f(i)
			wg.Done()
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- signale(struct{}{})
	}()

	return c
}

func main() {
	fmt.Println("start a group of workers...")
	// sync.Cond实例的初始化，需要一个满足实现了sync.Locker接口的类型实例，通常我们使用sync.Mutex
	groupSignal := sync.NewCond(&sync.Mutex{})
	c := spawnGroup(worker, 5, groupSignal)

	time.Sleep(time.Second * 5)
	fmt.Println("the group of workers start to work...")

	groupSignal.L.Lock()
	ready = true
	groupSignal.Broadcast()
	groupSignal.L.Unlock()

	<-c
	fmt.Println("the group of workers work done")
}
