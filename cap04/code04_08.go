/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/USntMXkOihY
Cap. 04.09. Bit Shift
*/

/*
Bitwise operation (Lógica binária)
110000
011000
001100
000110
000011
*/

/*
001 == 1
010 == 2
011 == 3
100 == 4
101 == 5
110 == 6
111 == 7
*/

package main

import "fmt"

func main() {
    x := 1
    y := x << 1

    fmt.Printf("%v\n", x)
    fmt.Printf("%b\n", x)
    fmt.Printf("%v\n", y)
    fmt.Printf("%b\n", y)

    fmt.Println("")

    a := 6
    b := a >> 1

    fmt.Printf("%v\n", a)
    fmt.Printf("%b\n", a)
    fmt.Printf("%v\n", b)
    fmt.Printf("%b\n", b)
}
