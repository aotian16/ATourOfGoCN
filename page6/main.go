package main

// 用括号组合导入包
import (
    "fmt"
    "math"
)

func main() {
	// 首字母大写的成员才可以导出(即可以访问)
	// 改为pi将报错
    fmt.Println(math.Pi)
}