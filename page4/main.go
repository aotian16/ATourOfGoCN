package main

// 导入包
import (
    "fmt"
    "math/rand"
)

func main() {
    // 打印随机数, 没有设置种子,所以一直都是1
    fmt.Println("My favorite number is", rand.Intn(10))
}