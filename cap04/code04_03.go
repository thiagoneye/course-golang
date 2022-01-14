/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/g0j0XaVk2EI
Cap. 04.04. Overflow
*/

package main

import "fmt"

func main() {
    var i uint16 = 65535
    fmt.Println(i)

    // i = 65536
    // fmt.Println(i)

    i = 65535
    i++
    fmt.Println(i)

    i++
    fmt.Println(i)
}
