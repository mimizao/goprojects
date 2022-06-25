package main

import (
	"fmt"
	"reflect"
)

/* 这里的i = pt不报错，但是i = t报错，很奇怪，明明pt也没有实现所有方法
type Interface interface {
	M1()
	M2()
}

type T struct{}

func (t T) M1()  {}
func (t *T) M2() {}

func main() {
	var t T
	var pt *T
	var i Interface

	i = pt
	i = t
}
*/

type T struct{}
type S T

func (t T) M1() {}
func (t T) M2() {}

func (t *T) M3() {}
func (t *T) M4() {}

func main() {
	var n int
	dumpMethodSet(n)
	dumpMethodSet(&n)

	var t T
	dumpMethodSet(t)
	dumpMethodSet(&t)

	var s S
	dumpMethodSet(s)
	dumpMethodSet(&s)
}

func dumpMethodSet(i interface{}) {
	dynTyp := reflect.TypeOf(i)
	if dynTyp == nil {
		fmt.Printf("there is no dynamic type\n")
		return
	}

	n := dynTyp.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty!\n", dynTyp)
		return
	}

	fmt.Printf("%s's method set:\n", dynTyp)
	for j := 0; j < n; j++ {
		fmt.Println("-", dynTyp.Method(j).Name)
	}
	fmt.Printf("\n")
}
