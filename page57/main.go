package main

import (
    "fmt"
    "net/http"
)

type Hello struct{}

// 类型 Hello 实现了 `http.Handler`。
func (h Hello) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
    fmt.Fprint(w, "Hello!")
}

func main() {
    var h Hello
    http.ListenAndServe("localhost:4000", h)
}
