/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/NxntmGW62Ag
Cap. 05.04.
*/

package main

import "fmt"

func main() {
    var x int
    x = 200

    fmt.Printf("%d, %b, %#x\n", x, x, x)

    y := x << 1
    fmt.Printf("%d, %b, %#x\n", y, y, y)
}
