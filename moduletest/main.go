package main

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Println("hello, gopath mode")
	logrus.Println(uuid.NewString())
}
