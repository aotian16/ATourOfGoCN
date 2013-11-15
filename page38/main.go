package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

// map 的文法跟结构体文法相似，不过必须有键名。
var m = map[string]Vertex{
    "Bell Labs": Vertex{
        40.68433, -74.39967,
    },
    "Google": Vertex{
        37.42202, -122.08408,
    },
}

func main() {
    fmt.Println(m)
}
