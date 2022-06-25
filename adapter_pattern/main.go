package main

import (
	"fmt"
	"net/http"
)

func greetings(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome")
}

// 这里的http.HandlerFunc就是一个适配器函数，将这个greetings函数转换成http.HandlerFunc函数，然后这个http.HandlerFunc函数又实现了http.Handler这个接口，所以就可以用了
func main() {
	http.ListenAndServe(":9999", http.HandlerFunc(greetings))
}
