package main

// 用括号组合导入包
import (
    "fmt"
)

// 定义函数
// 注意点: 返回类型在方法之后
func add(x int, y int) int {
    return x + y
}

func main() {
	// 调用函数
    fmt.Println(add(42, 13))
}