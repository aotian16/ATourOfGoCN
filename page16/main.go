package main

import "fmt"

// 数值常量是高精度的 _值_。
//一个未指定类型的常量由上下文来决定其类型。

const (
    Big   = 1 << 100
    Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
    return x * 0.1
}

func main() {
    fmt.Println(needInt(Small))
    // fmt.Println(needInt(Big)) 溢出
    fmt.Println(needFloat(Small))
    fmt.Println(needFloat(Big))
}
