package main

import (
	"fmt"
)

func main() {
	// 	var sl = [][]int{
	// 		{1, 34, 26, 35, 78},
	// 		{3, 45, 13, 24, 99},
	// 		{101, 13, 38, 7, 127},
	// 		{54, 27, 40, 83, 81},
	// 	}

	// continue_outerloop:
	// 	for i := 0; i < len(sl); i++ {
	// 		for j := 0; j < len(sl[i]); j++ {
	// 			if sl[i][j] == 13 {
	// 				fmt.Printf("found 13 at [%d,%d]\n", i, j)
	// 				continue continue_outerloop
	// 			}
	// 		}
	// 	}

	// break_outerloop:
	// 	for i := 0; i < len(sl); i++ {
	// 		for j := 0; j < len(sl[i]); j++ {
	// 			if sl[i][j] == 13 {
	// 				fmt.Printf("found 13 at [%d,%d]\n", i, j)
	// 				break break_outerloop
	// 			}
	// 		}
	// 	}

	// var m = []int{1, 2, 3, 4, 5}
	// 这里实际输出的都是4和5，因为i和v在range中只会被声明一次
	// for i, v := range m {
	// 	go func() {
	// 		time.Sleep(time.Second * 3)
	// 		fmt.Println(i, v)
	// 	}()
	// }
	// time.Sleep(time.Second * 10)

	// for i, v := range m {
	// 	go func(i, v int) {
	// 		time.Sleep(time.Second * 3)
	// 		fmt.Println(i, v)
	// 	}(i, v)
	// }
	// time.Sleep(time.Second * 10)

	// var a = [5]int{1, 2, 3, 4, 5}
	// var r [5]int
	// fmt.Println("original a = ", a)
	// 这里实际不会修改a的值，因为参与range的是a的副本
	// for i, v := range a {
	// 	if i == 0 {
	// 		a[1] = 12
	// 		a[2] = 13
	// 	}
	// 	r[i] = v
	// }

	// 如果用切片的话就可以修改，因为切片虽然也是副本，但是在底层实际指向的是同一个地址
	// for i, v := range a[:] {
	// 	if i == 0 {
	// 		a[1] = 12
	// 		a[2] = 13
	// 	}
	// 	r[i] = v
	// }
	// fmt.Println("after for range loop, r = ", r)
	// fmt.Println("after for range loop, a = ", a)

	var m = map[string]int{
		"tony": 21,
		"tom":  22,
		"jim":  23,
	}
	counter := 0
	// for k, v := range m {
	// 	if counter == 0 {
	// 		delete(m, "tony")
	// 		fmt.Println(m)
	// 	}
	// 	counter++
	// 	fmt.Println(k, v)
	// }
	for k, v := range m {
		if counter == 0 {
			m["lucy"] = 24
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter: ", counter)
}
