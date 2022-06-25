package main

import "fmt"

type T int

func (t T) Error() string {
	return "bad error"
}

func printNilInterface() {
	var i any
	var err error
	println(i)
	println(err)
	println("i = nil:", i == nil)
	println("err = nil:", err == nil)
	println("i = err:", i == err)
	println("--------------------------------")
}

func printEmptyInterface() {
	var eif1 any
	var eif2 any
	var n, m int = 17, 18

	eif1 = n
	eif2 = m

	println("eif1:", eif1)
	println("eif2:", eif2)
	println("eif1 = eif2:", eif1 == eif2)
	println("--------------------------------")
	eif2 = 17
	println("eif1:", eif1)
	println("eif2:", eif2)
	println("eif1 = eif2:", eif1 == eif2)
	println("--------------------------------")
	eif2 = int64(17)
	println("eif1:", eif1)
	println("eif2:", eif2)
	println("eif1 = eif2:", eif1 == eif2)
	println("--------------------------------")
}

func printNonEmptyInterface() {
	var err1 error
	var err2 error
	err1 = (*T)(nil)
	println("err1", err1)
	println("err1 = nil", err1 == nil)
	println("--------------------------------")

	err1 = T(5)
	err2 = T(6)
	println("err1:", err1)
	println("err2:", err2)
	println("err1 = err2:", err1 == err2)
	println("--------------------------------")

	err2 = fmt.Errorf("%d\n", 5)
	println("err1:", err1)
	println("err2:", err2)
	println("err1 = err2:", err1 == err2)
	println("--------------------------------")
}

func printEmptyInterfaceAndNonEmptyInterface() {
	var eif any = T(5)
	var err error = T(5)
	println("eif:", eif)
	println("err:", err)
	println("eif = err:", eif == err)
	println("--------------------------------")

	err = T(6)
	println("eif:", eif)
	println("err:", err)
	println("eif = err:", eif == err)
	println("--------------------------------")
}

func main() {
	printNilInterface()
	printEmptyInterface()
	printNonEmptyInterface()
	printEmptyInterfaceAndNonEmptyInterface()
}
