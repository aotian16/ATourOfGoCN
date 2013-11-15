package main

import "fmt"

func main() {
    // slice, 切片
    // 一个 slice 会指向一个数组，并且包含了长度信息。
    // []T 是一个元素类型为 T 的 slice。
    
    p := []int{2, 3, 5, 7, 11, 13}
    fmt.Println("p ==", p)

    for i := 0; i < len(p); i++ {
        fmt.Printf("p[%d] == %d\n",
            i, p[i])
    }
}
