package main

import "fmt"

func main() {
    pow := make([]int, 10)
    
    // 可以将值赋值给 _ 来忽略序号和值。
    // 如果只需要索引值，去掉“, value”的部分即可忽略值。
    for i := range pow {
        pow[i] = 1 << uint(i)
    }
    
    for _, value := range pow {
        fmt.Printf("%d\n", value)
    }
}
