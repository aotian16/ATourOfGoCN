package main

import "fmt"

func sum(a []int, c chan int) {
    sum := 0
    for _, v := range a {
        sum += v
    }
    c <- sum // 将和送入 c
}

// channel 是有类型的管道，可以用 channel 操作符 <- 对其发送或者接收值。
// ch <- v    // 将 v 送入 channel ch。
// v := <-ch  // 从 ch 接收，并且赋值给 v。
// （“箭头”就是数据流的方向。）
// 和 map 与 slice 一样，channel 使用前必须创建：
// ch := make(chan int)
// 默认情况下，在另一端准备好之前，发送和接收都会阻塞。这使得 goroutine 可以在没有明确的锁或竞态变量的情况下进行同步。

func main() {
    a := []int{7, 2, 8, -9, 4, 0}

    // 创建管道
    c := make(chan int)
    
    // 开启新线程计算前一半数字的和
    go sum(a[:len(a)/2], c)
    
    // 开启新线程计算后一半数字的和
    go sum(a[len(a)/2:], c)
    
    x, y := <-c, <-c // 从 c 中获取

    fmt.Println(x, y, x+y)
}
