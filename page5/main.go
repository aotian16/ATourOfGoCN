package main

// 用括号组合导入包
import (
    "fmt"
    "math"
)

// 也可以用多行导入,效果一样
// import "fmt"
// import "math"

func main() {
    fmt.Printf("Now you have %g problems.",
        math.Nextafter(2, 3))
}