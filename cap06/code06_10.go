/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/Y5HamAGQzUg
Cap. 06.11. Logical Operators
*/

package main

import "fmt"

func main() {
    x := 3

    if ((x == 2) || (x == 3)) {
        fmt.Println("x é 2 ou 3")
    }

    if ((x % 2 != 0) && (x > 0)) {
        fmt.Println("x é um positivo não múltiplo de 2")
    }
}
