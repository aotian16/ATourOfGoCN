package main

import "fmt"

func main() {
    // 在函数中，`:=` 简洁赋值语句在明确类型的地方，可以用于替代 var 定义。
    // 函数外的每个语法块都必须以关键字开始（`var`、`func`、等等），`:=` 结构不能使用在函数外。
    var x, y, z int = 1, 2, 3
    c, python, java := true, false, "no!"

    fmt.Println(x, y, z, c, python, java)
}
