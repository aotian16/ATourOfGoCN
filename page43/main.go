package main

import "fmt"

// Go 函数可以是闭包的。
// 闭包是一个函数值，它来自函数体的外部的变量引用。
// 函数可以对这个引用值进行访问和赋值；换句话说这个函数被“绑定”在这个变量上。
// (百度百科)在编程领域我们可以通俗的说：子函数可以使用父函数中的局部变量，这种行为就叫做闭包！

// 例如，函数 adder 返回一个闭包。每个闭包都被绑定到其各自的 sum 变量上。
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    pos, neg := adder(), adder()
    for i := 0; i < 10; i++ {
        fmt.Println(
            pos(i),
            neg(-2*i),
        )
    }
}
