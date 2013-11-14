package main

// 用括号组合导入包
import (
    "fmt"
)

// 命名返回值
// 函数接受参数。在 Go 中，函数可以返回多个“结果参数”，而不仅仅是一个值。它们可以像变量那样命名和使用。
// 如果命名了返回值参数，一个没有参数的 return 语句，会将当前的值作为返回值返回。
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func main() {
    fmt.Println(split(17))
}