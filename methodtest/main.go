package main

import "fmt"

type T struct {
	a int
}

func (t T) M(n int) {}

func (t T) Get() int {
	return t.a
}

// 类型T的方法Get的等价函数
func Get(t T) int {
	return t.a
}

func (t *T) Set(a int) int {
	t.a = a
	return t.a
}

// 类型*T的方法Set的等价函数
func Set(t *T, a int) int {
	t.a = a
	return t.a
}

func main() {
	var t T
	t.M(1) // 通过类型T的变量实例调用方法M

	p := &T{}
	p.M(2) // 通过类型*T的变量实例调用方法M

	// 这种直接以类型名 T 调用方法的表达方式，被称为 Method Expression。通过 Method Expression 这种形式，类型 T 只能调用 T 的方法集合（Method Set）中的方法，同理类型 *T 也只能调用 *T 的方法集合中的方法。
	// Go 语言中的方法的本质就是，一个以方法的 receiver 参数作为第一个参数的普通函数。
	t.Get()
	T.Get(t)
	(&t).Set(1)
	(*T).Set(&t, 1)

	f1 := (*T).Set                           // f1的类型，也是*T类型Set方法的类型: func(t *T,int)int
	f2 := T.Get                              // f2的类型，也是T类型Get方法的类型：func(t T)int
	fmt.Printf("the type of f1 is %T\n", f1) // the type of f1 is func(*main.T, int) int 这里有一点没有明白，为什么是*main.T，难道是因为main包？
	fmt.Printf("the type of f2 is %T\n", f2) // the type of f2 is func(main.T) int
	f1(&t, 3)
	fmt.Println(f2(t))
}
