package main

import "fmt"

// 一个结构体（`struct`）就是一个字段的集合。

// 点结构体 
type Point struct {
    X int
    Y int
}

// 线结构体
type Line struct {
    A Point
    B Point
}

func main() {
    var p0=Point{1, 2}
    var p1=Point{3, 4}
    var line=Line{p0, p1}
    fmt.Println(p0, p1)
    fmt.Println(line)
}
