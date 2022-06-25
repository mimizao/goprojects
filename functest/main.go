package main

import "fmt"

func myAppend(sl []int, elems ...int) []int {
	fmt.Printf("%T\n", elems) // 输出elems的类型
	if len(elems) == 0 {
		println("no elems to append")
		return sl
	}
	sl = append(sl, elems...)
	return sl
}

func setup(task string) func() {
	println("do some setup stuff for", task)
	// 这个返回的函数就是上下文拆除函数
	return func() {
		println("do some teardown stuff for", task)
	}
}

func partialTimes(x int) func(int) int {
	return func(y int) int {
		return times(x, y)
	}
}

func times(x int, y int) int {
	return x * y
}

func main() {
	sl := []int{1, 2, 3}
	sl = myAppend(sl) // 通过这里的输出能看出，变长参数的elems的类型是[]int，也就是切片
	fmt.Println(sl)
	sl = myAppend(sl, 4, 5, 6)
	fmt.Println(sl)

	// 这里的teardown就是上下文拆除函数
	teardown := setup("demo")
	defer teardown()
	println("do some bussiness stuff")
}
