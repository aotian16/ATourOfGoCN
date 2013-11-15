package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    p := Vertex{1, 2}
    
    // Go 有指针，但是没有指针运算。
    // 结构体字段可以通过结构体指针来访问。通过指针间接的访问是透明的。
    q := &p
    q.X = 3
    fmt.Println(p)
}
