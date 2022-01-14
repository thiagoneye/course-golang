/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/Yaw80pKukMc
Cap. 04.07. Constants
*/

/*
Comments:
- Constantes são imutáveis.
- Constantes podem ser tipadas ou não.
- As constantes não tipadas só terão um tipo atribuido a elas quando forem usadas.
*/

package main

import "fmt"

const (
    s1 = "Oi"         // Constante não tipada
    s2 string = "Oi"  // Constante tipada
)

const x = 10            // Tipo indefinido
var y = 10              // Tipo definido

var (
    a int
    b float64
)

func main() {
    fmt.Println(s1, s2)
    fmt.Println(x, y)

    fmt.Println("")

    fmt.Printf("%v: %T\n", x, x)
    fmt.Printf("%v: %T\n", y, y)

    fmt.Println("")

    a = x
    b = x

    fmt.Printf("%v: %T\n", a, a)
    fmt.Printf("%v: %T\n", b, b)

    /*
    fmt.Println("")

    a = y
    b = y

    fmt.Printf("%v: %T\n", a, a)
    fmt.Printf("%v: %T\n", b, b)
    */
}
