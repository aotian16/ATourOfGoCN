package main

import (
    "fmt"
    "time"
)

// 定义异常结构
type MyError struct {
    When time.Time
    What string
}

// 实现异常方法
func (e *MyError) Error() string {
    return fmt.Sprintf("at %v, %s",
        e.When, e.What)
}

// 要抛出异常的方法
func run() error {
    return &MyError{
        time.Now(),
        "it didn't work",
    }
}

func main() {
    if err := run(); err != nil {
        fmt.Println(err)
    }
}
