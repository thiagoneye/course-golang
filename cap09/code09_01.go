/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/cjUFrS3hqgU
Cap. 09.01.
*/

package main

import "fmt"

func main() {
    x := [5]int{10, 20, 30, 40, 50} // Literal Composta

    for i, v := range x {
        fmt.Printf("No índice %v, temos o valor %v.\n", i, v)
    }

    fmt.Printf("%v: %T", x, x)
}
