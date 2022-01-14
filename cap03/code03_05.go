// https://github.com/ellenkorbes/aprendago
// https://www.youtube.com/c/AprendaGo/videos
// https://youtu.be/O0r318FN_Uw
// Cap. 03.05.

package main

import "fmt"

type neytype int
var x neytype
var y int

func main() {
    fmt.Printf("x (%T) = %v\n", x, x)
    x = 42
    fmt.Println(x)
    y = int(x)
    fmt.Printf("y (%T) = %v\n", y, y)
}
