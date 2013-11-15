package main

import "fmt"

func main() {
    // slice 由函数 make 创建。
    // 这会分配一个零长度的数组并且返回一个 slice 指向这个数组：
    // a := make([]int, 5)  // len(a)=5
    // 为了指定容量，可传递第三个参数到 `make`：
    // b := make([]int, 0, 5) // len(b)=0, cap(b)=5
    // b = b[:cap(b)] // len(b)=5, cap(b)=5
    // b = b[1:]      // len(b)=4, cap(b)=4
    
    a := make([]int, 5)
    printSlice("a", a)
    b := make([]int, 0, 5)
    printSlice("b", b)
    c := b[:2]
    c[1]=5
    printSlice("c", c)
    d := c[2:5]
    printSlice("d", d)
    e := c[0:5]
    printSlice("e", e)
}

func printSlice(s string, x []int) {
    fmt.Printf("%s len=%d cap=%d %v\n",
        s, len(x), cap(x), x)
}
