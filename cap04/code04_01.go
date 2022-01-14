/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/voisg61hPXA
Cap. 04.01. Tipo booleano
*/

package main

import "fmt"

var x bool

func main() {
    fmt.Println(x) // Zero value

    x = true
    fmt.Println(x)

    x = (10 < 100)
    fmt.Println(x)

    x = (10 == 100)
    fmt.Println(x)

    x = (10 > 100)
    fmt.Println(x)
}
