/* 一多对
package main

import (
	"fmt"
	"sync"
	"time"
)

type signal struct{}

func worker(i int) {
	fmt.Printf("worker %d is working... \n", i)
	time.Sleep(time.Second)
	fmt.Printf("worker %d is done\n", i)
}

func spawnGroup(f func(i int), num int, groupSignal <-chan signal) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			<-groupSignal
			fmt.Printf("worker %d start to work... \n", i)
			f(i)
			wg.Done()
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- signal{}
	}()
	return c
}

func main() {
	fmt.Println("start a group of workers...")
	groupSignal := make(chan signal)
	c := spawnGroup(worker, 5, groupSignal)
	time.Sleep(5 * time.Second)
	fmt.Println("the group of workers start to work...")
	close(groupSignal)
	<-c
	fmt.Println("the group of workers stop to work...")
}
*/

/* 基于“共享内存”+“互斥锁”实现
package main

import (
	"fmt"
	"sync"
)

type counter struct {
	sync.Mutex
	i int
}

var cter counter

func Increase() int {
	cter.Lock()
	defer cter.Unlock()
	cter.i++
	return cter.i
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			v := Increase()
			fmt.Printf("groutine %d: %v\n", i, v)
			wg.Done()
		}(i)
		wg.Wait()
	}
}
*/

/* 无缓冲channel代替锁的实现
package main

import (
	"fmt"
	"sync"
)

type counter struct {
	c chan int
	i int
}

func NewCounter() *counter {
	cter := &counter{
		c: make(chan int),
	}
	go func() {
		for {
			cter.i++
			cter.c <- cter.i
		}
	}()
	return cter
}

func (c *counter) Increase() int {
	return <-c.c
}

func main() {
	cter := NewCounter()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			v := cter.Increase()
			fmt.Printf("goroutine %d: %d\n", i, v)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
*/

/* 用作计数信号量（counting semaphore)
package main

import (
	"log"
	"sync"
	"time"
)

var active = make(chan struct{}, 3)
var jobs = make(chan int, 10)

func main() {
	go func() {
		for i := 0; i < 8; i++ {
			jobs <- (i + 1)
		}
		close(jobs)
	}()

	var wg sync.WaitGroup
	for j := range jobs {
		wg.Add(1)
		go func(j int) {
			active <- struct{}{}
			log.Printf("handle job:%d\n", j)
			time.Sleep(2 * time.Second)
			<-active
		}(j)
	}
	wg.Wait()
}
*/

/* len相关的测试
package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(c chan<- int) {
	var i int = 1
	for {
		time.Sleep(time.Second * 2)
		ok := trySend(c, i)
		if ok {
			fmt.Printf("[producer]:send [%d] to channel\n", i)
			i++
			continue
		}
		fmt.Printf("[producer]:try send [%d],but channel is empty\n", i)
	}
}

func tryRecv(c <-chan int) (int, bool) {
	select {
	case i := <-c:
		return i, true
	default:
		return 0, false
	}
}

func trySend(c chan<- int, i int) bool {
	select {
	case c <- i:
		return true
	default:
		return false
	}
}

func consumer(c <-chan int) {
	for {
		i, ok := tryRecv(c)
		if !ok {
			fmt.Println("[consumer]: try to recv from channel, but the channel is empty")
			time.Sleep(1 * time.Second)
			continue

		}
		fmt.Printf("[consumer]:recv [%d] from channel\n", i)
		if i >= 3 {
			fmt.Println("[consumer]:exit")
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	c := make(chan int, 3)
	wg.Add(2)
	go func() {
		producer(c)
		wg.Done()
	}()

	go func() {
		consumer(c)
		wg.Done()
	}()

	wg.Wait()
}
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- 5
		close(ch1)
	}()

	go func() {
		time.Sleep(7 * time.Second)
		ch2 <- 7
		close(ch2)
	}()

	// var ok1, ok2 bool
	// for {
	// 	select {
	// 	case x := <-ch1:
	// 		ok1 = true
	// 		fmt.Println(x)
	// 	case x := <-ch2:
	// 		ok2 = true
	// 		fmt.Println(x)
	// 	}
	// 	if ok1 && ok2 {
	// 		break
	// 	}
	// }
	for {
		select {
		case x, ok1 := <-ch1:
			if !ok1 {
				ch1 = nil
			} else {
				fmt.Println(x)
			}
		case x, ok2 := <-ch2:
			if !ok2 {
				ch2 = nil
			} else {
				fmt.Println(x)
			}
		}
		if ch1 == nil && ch2 == nil {
			break
		}
	}

	fmt.Println("program finished successfully")
}
