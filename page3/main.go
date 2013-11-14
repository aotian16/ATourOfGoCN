package main

// 导入
import (
    "fmt"
    "net"
    "os"
    "time"
)

func main() {
    // 打印文字
    fmt.Println("Welcome to the playground!")
    // 打印时间
    fmt.Println("The time is", time.Now())
    // 打印文件
    fmt.Println("And if you try to open a file:")
    fmt.Println(os.Open("main.go"))
    // 打印网络
    fmt.Println("Or access the network:")
    fmt.Println(net.Dial("tcp", "www.google.com:80"))
}