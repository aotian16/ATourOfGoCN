package main

import (
    "fmt"
)

type Message struct {
    M string
}

// 方法接收者为指针时，可以改变对象值
// 避免在方法调用时拷贝值，效率高
func (v *Message) setM0(m string) {
    v.M = m
}

// 方法接收者为对象时，无法改变对象值
func (v Message) setM1(m string) {
    v.M = m
}

func (v *Message) getM() string {
    return v.M
}

func main() {
    msg := &Message{"hello"}
   
    fmt.Println(msg, msg.getM())
    msg.setM1("world")
    fmt.Println(msg, msg.getM())
    msg.setM0("world")
    fmt.Println(msg, msg.getM())
}
