package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

// Go 没有类。然而，仍然可以在结构体类型上定义方法。
// 方法接收者 出现在 func 关键字和方法名之间的参数中。
func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
    v := &Vertex{3, 4}
    fmt.Println(v.Abs())
}
