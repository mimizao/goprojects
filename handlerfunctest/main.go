package main

import (
	"fmt"
	"net/http"
)

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, Gopher!\n")
}

func main() {
	http.ListenAndServe(":8888", http.HandlerFunc(greeting))
}

// 1、说一下对于http.ListenAndServe的理解，代码如下需要一个string类型，一个Handler的参数
// func ListenAndServe(addr string, handler Handler) error {
//     server := &Server{Addr: addr, Handler: handler}
//     return server.ListenAndServe()
// }

// 2、可以看到这里的Handler参数，这是一个自定义接口类型，其中有一个ServeHttp方法，但是我们在上面传入的参数是http.HandlerFunc(greeting)，并不是Handler
// type Handler interface {
// 		ServeHTTP(ResponseWriter, *Request)
// }

// 3、之后我们可以看下HandlerFunc的定义，可以看到这是一个自定义的函数，并且这个函数实现了ServeHTTP方法，所以我们可以说这个函数也就是Handler接口类型
// type HandlerFunc func(ResponseWriter, *Request)

// func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
// 		f(w, r)
// }

// 4、总结一下，http.HandlerFunc(greeting)这句就是将函数greeting显示转换为HandlerFunc类型，后者实现了Handler接口，满足ListenAndServe函数第二个参数的要求
