package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Print("Go runs on ")
    
    // switch 可以用于string
    // 除非使用 fallthrough 语句作为结尾，否则 case 部分会自动终止。
    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("OS X.")
    case "linux":
        fmt.Println("Linux.")
    default:
        // freebsd, openbsd,
        // plan9, windows...
        fmt.Printf("%s.", os)
    }
}
