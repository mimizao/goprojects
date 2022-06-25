package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	error
}

var ErrBad = MyError{
	error: errors.New("bad thins happened"),
}

func bad() bool {
	return false
}

func returnError() error {
	var p *MyError = nil
	if bad() {
		p = &ErrBad
	}
	return p
}

func main() {
	err := returnError()
	if err != nil {
		fmt.Printf("error occurred: %v\n", err)
		return
	}
	fmt.Println("ok")
}
