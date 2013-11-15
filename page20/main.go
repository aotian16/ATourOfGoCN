package main

import "fmt"

func main() {
    var n =1
    // 如果省略了循环条件，循环就不会结束，因此可以用更简洁地形式表达死循环。
    for {
        fmt.Println(n)
        n++;
        if(n>10){
        	break
        }
    }
}
