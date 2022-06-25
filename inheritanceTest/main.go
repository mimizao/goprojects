package main

import (
	"fmt"
	"io"
	"strings"
)

/* 第一种方式，通过接口类型的类型嵌入
type E interface {
	M1()
	M2()
}

type I interface {
	E
	M3()
}
*/

/* 第二种，类型嵌入
type T1 int
type t2 struct {
	n int
	m int
}

type I interface {
	M1()
}

type S1 struct {
	T1
	*t2
	I
	a int
	b string
}
*/

type MyInt int

func (n *MyInt) Add(m int) {
	*n = *n + MyInt(m)
}

type t struct {
	a int
	b int
}

type S struct {
	*MyInt
	t
	io.Reader
	s string
	n int
}

func main() {
	m := MyInt(17)
	r := strings.NewReader("hello world")
	s := S{
		MyInt: &m,
		t: t{
			a: 1,
			b: 2,
		},
		Reader: r,
		s:      "demo",
	}
	var sl = make([]byte, len("hello world"))
	s.Reader.Read(sl)
	s.Read(sl) // 这两个方法都已经被继承过来了，可以直接用
	fmt.Println(string(sl))
	s.MyInt.Add(5)
	s.Add(5) // 这个同上
	fmt.Println(*(s.MyInt))
}
