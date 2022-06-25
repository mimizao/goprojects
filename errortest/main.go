package main

import (
	"errors"
	"fmt"
)

var ErrSentinel = errors.New("the underlying sendinel error")

type MyError struct {
	e string
}

func (e *MyError) Error() string {
	return e.e
}

func main() {
	err1 := fmt.Errorf("wrap sentinel: %w", ErrSentinel)
	err2 := fmt.Errorf("wrap err1: %w", err1)
	println(err2 == ErrSentinel) // false
	if errors.Is(err2, ErrSentinel) {
		println("err2 is ErrSentinel") // 这里会输出，说明进来了
		//return
	}
	println("err2 is not ErrSentinel")

	println("--------------------------------")
	var err = &MyError{"MyError error demo"}
	err3 := fmt.Errorf("wrap err: %w", err)
	err4 := fmt.Errorf("wrap err3: %w", err3)
	var e *MyError
	if errors.As(err4, &e) {
		println("MyError is on the chain of err4")
		println(e == err)
		return
	}
	println("MyError is not on the chain of err4")
}
