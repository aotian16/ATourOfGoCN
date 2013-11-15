package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

// 如果顶级的类型只有类型名的话，可以在文法的元素中省略键名。
var m = map[string]Vertex{
    "Bell Labs": {40.68433, -74.39967},
    "Google":    {37.42202, -122.08408},
}

func main() {
    fmt.Println(m)
}
