/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/i_3O4ooSVCM
Cap. 08.01. Array
*/

/*
Arrays possuem comprimento fixo e são pouco flexíveis.
Um array possuem tipo [n]type.
*/

package main

import "fmt"

var x [5]int

func main() {
    x[0] = 1
    x[1] = 10
    fmt.Println(x)
    fmt.Println(x[1])
    fmt.Printf("%T\n", x)
}
