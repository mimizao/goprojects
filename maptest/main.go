package main

import (
	"fmt"
	"time"
)

func main() {
	// m1:=map[int][]string{
	// 	1:[]string{"vale1_1","vale1_2"},
	// 	3:[]string{"vale3_1","vale3_2","vale3_3"},
	// 	7:[]string{"vale7_1"},
	// }

	// m1:=make(map[int]string)
	// m2:=make(map[int]string,8)

	// m := make(map[int]string)
	// m[1] = "value1"
	// m[2] = "value2"
	// m[3] = "value3"
	// fmt.Println(len(m))
	// m[4] = "value4"
	// fmt.Println(len(m))
	// _, ok := m[5]
	// if !ok {
	// 	fmt.Println("k = 5时, 没有对应的value")
	// }

	// delete(m, 1)
	// fmt.Println(m)
	// for k, v := range m {
	// 	fmt.Printf("[%d, %s] ", k, v)
	// }

	// map与并发
	m := map[int]int{
		1: 11,
		2: 12,
		3: 13,
	}
	go func() {
		for i := 0; i < 1000; i++ {
			doIteration(m)
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			doWrite(m)
		}
	}()

	time.Sleep(5 * time.Second)
}

func doIteration(m map[int]int) {
	for k, v := range m {
		_ = fmt.Sprintf("[%d,%d]", k, v)
	}
}

func doWrite(m map[int]int) {
	for k, v := range m {
		m[k] = v + 1
	}
}
