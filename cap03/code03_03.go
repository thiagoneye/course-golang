// https://github.com/ellenkorbes/aprendago
// https://www.youtube.com/c/AprendaGo/videos
// https://youtu.be/VkuZ4QMnoZM
// Cap. 03.03.

package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
    s := fmt.Sprintf("x = %v \ny = %v \nz = %v", x, y, z)
    fmt.Println(s)
}
