package main

// 用括号组合导入包
import (
    "fmt"
)

// 定义函数
// 注意点: 返回类型在方法之后
// 当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略。
func add(x , y int) int {
    return x + y
}

func main() {
	// 调用函数
    fmt.Println(add(42, 13))
}