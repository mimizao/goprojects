package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}

func (p *field) sleepPrint() {
	time.Sleep(time.Second)
	fmt.Println(p.name)
}

// 上面的方法等价于下面的函数
// func print(p *field) { fmt.Println(p.name) }

func main() {
	data1 := []*field{{"one"}, {"two"}, {"three"}}
	// for _, v := range data1 {
	// 	go v.print()
	// }

	// 等价于，这里也是输出ont,two,three
	// for _, v := range data1 {
	// 	go (*field).print(v)
	// }

	// 但是神奇的事情是如果用下面这种话，就都是three
	// for _, v := range data1 {
	// 	go func() {
	// 		time.Sleep(time.Second)
	// 		(*field).print(v)
	// 	}()
	// }

	// 当然我可以通过在闭包时直接传入参数的形式使得输出one,two,three
	// for _, v := range data1 {
	// 	go func(v *field) { // 注意这里传入的是v，后面闭包里面传入的也是v
	// 		time.Sleep(time.Second)
	// 		(*field).print(v)
	// 	}(v)
	// }

	//再换一种方式试试
	for _, v := range data1 {
		go (*field).sleepPrint(v)
	}
	// 经过这么多轮测试我好像有点明白了，就是在上面的for _, v := range那里，
	// 如果是v的类型是*field的话，那么我实际上就是把中三个*field都传到go后面的函数里面去了，
	// 当这个函数真正去运行的时候，输出的是我传入的那个*field
	// 下面为啥不行呢，因为我最开始虽然用&v实现传入函数的参数是*field，也就是这个v的地址
	// 但是当我的函数真正去运行的时候，再去这个地址寻找值的时候，存储的都是最后一个six的了
	data2 := []field{{"four"}, {"five"}, {"six"}}
	// for _, v := range data2 {
	// 	go v.print()
	// }
	// 等价于
	for _, v := range data2 {
		go (*field).print(&v)
	}
	time.Sleep(time.Second * 3)
}
