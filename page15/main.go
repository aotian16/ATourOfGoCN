package main

import "fmt"

const Pi = 3.14

// 常量的定义与变量类似，只不过使用 const 关键字。
// 常量可以是字符、字符串、布尔或数字类型的值。
// 常量不能使用 := 语法定义。
func main() {
    const World = "世界"
    fmt.Println("Hello", World)
    fmt.Println("Happy", Pi, "Day")

    const Truth = true
    fmt.Println("Go rules?", Truth)
}
