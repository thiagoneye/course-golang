/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/dRNNC7VpztE
Cap. 08.08. Slice: The underlying array surprise
*/

/*
É necessário tomar cuidado quando se cria um slice a partir de outro slice, pois
pode ser alterado de forma indesejada o primeiro slice (assim como quando se
trabalha com objetos em Python).
*/

package main

import "fmt"

func main() {
    slice1 := []int{1, 2, 3, 4, 5}
    fmt.Println(slice1)

    slice2 := append(slice1[:2], slice1[4:]...)
    fmt.Println(slice2)

    fmt.Println(slice1)
}
