// https://github.com/ellenkorbes/aprendago
// https://www.youtube.com/c/AprendaGo/videos
// https://youtu.be/5-0S-gefNe0
// Cap. 03.04.

package main

import "fmt"

type neytype int
var x neytype

func main() {
    fmt.Printf("x (%T) = %v\n", x, x)
    x = 42
    fmt.Println(x)
}
