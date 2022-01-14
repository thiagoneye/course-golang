/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/3yIKCLSWAHA
Cap. 04.03. Tipos numéricos
*/

package main

import "fmt"

func main() {
    a := "e"
    b := "é"
    c := "愛"
    fmt.Printf("%v, %v, %v\n", a, b, c)

    d := []byte(a)
    e := []byte(b)
    f := []byte(c)
    fmt.Printf("%v, %v, %v\n", d, e, f)

    x := 10
    y := 10.0
    fmt.Printf("%v: %T \n%v: %T", x, x, y, y)
}
