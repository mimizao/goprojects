package main

import "fmt"

type MyInterface interface {
	M1()
}

type T int

func (T) M1() {
	println("T's M1")
}

func main() {
	// var a int64 = 13
	// var i any = a
	// v1, ok := i.(int64)
	// fmt.Printf("v1 = %d,the type of v1 is %T, ok = %t\n", v1, v1, ok)
	// v2, ok := i.(string)
	// fmt.Printf("v2 = %d,the type of v2 is %T, ok = %t\n", v2, v2, ok)
	// v3 := i.(int64)
	// fmt.Printf("v3 = %d,the type of v3 is %T\n", v3, v3)
	// v4 := i.([]int)
	// fmt.Printf("v4 = %d,the type of v4 is %T\n", v4, v4)

	var t T
	var i any = t
	v1, ok := i.(MyInterface)
	if !ok {
		panic("the value is not a MyInterface")
	}
	v1.M1()
	fmt.Printf("the type of v1 is %T\n", v1)

	i = int64(13)
	v2, ok := i.(MyInterface)
	fmt.Printf("the type of v2 is %T\n", v2)
}
