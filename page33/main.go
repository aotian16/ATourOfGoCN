package main

import "fmt"

func main() {
    // slice 的零值是 `nil`。
    // 一个 nil 的 slice 的长度和容量是 0。
    var z []int
    fmt.Println(z, len(z), cap(z))
    if z == nil {
        fmt.Println("nil!")
    }
}
