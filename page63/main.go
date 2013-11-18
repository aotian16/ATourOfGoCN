package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

// goroutine 是由 Go 运行时环境管理的轻量级线程。
// go f(x, y, z)
// 开启一个新的 goroutine 执行
// f(x, y, z)

func main() {
    go say("world")
    say("hello")
}
