package main

import "fmt"

func main() {
    m := make(map[string]int)

    // 在 map m 中插入或修改一个元素：
    m["Answer"] = 42
    // 获得元素：
    fmt.Println("The value:", m["Answer"])

    // 在 map m 中插入或修改一个元素：
    m["Answer"] = 48
    // 获得元素：
    fmt.Println("The value:", m["Answer"])

    // 删除元素：
    delete(m, "Answer")
    fmt.Println("The value:", m["Answer"])

    // 通过双赋值检测某个键存在：
    v, ok := m["Answer"]
    fmt.Println("The value:", v, "Present?", ok)
}
